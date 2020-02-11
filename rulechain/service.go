//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package rulechain

import (
	"context"
	"errors"
	"reflect"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/factory"
	"github.com/cloustone/pandas/models/notifications"
	"github.com/cloustone/pandas/pkg/broadcast"
	broadcast_util "github.com/cloustone/pandas/pkg/broadcast/util"
	"github.com/cloustone/pandas/rulechain/converter"
	pb "github.com/cloustone/pandas/rulechain/grpc_rulechain_v1"
	"github.com/cloustone/pandas/rulechain/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	nameOfRuleChain = reflect.TypeOf(models.RuleChain{}).Name()
)

// RuleChainService implement all rulechain interface
type RuleChainService struct {
	servingOptions *options.ServingOptions
	controller     *runtimeController
}

// NewRuleChainService return rulechain service object
func NewRuleChainService(servingOptions *options.ServingOptions) *RuleChainService {
	//factory.RegisterFactory(models.RuleChain{}, newRuleChainFactory(servingOptions))
	return &RuleChainService{
		servingOptions: servingOptions,
		controller:     newRuntimeController(),
	}
}

// notify is internal helper to simplify broadcastization notificaiton
func notify(path string, action string, param models.Model) {
	broadcast_util.Notify(path, action, param)
}

// CheckRuleChain check wether the rule chain is valid
func (s *RuleChainService) CheckRuleChain(ctx context.Context, in *pb.CheckRuleChainRequest) (*pb.CheckRuleChainResponse, error) {
	resp := pb.CheckRuleChainResponse{
		Reasons: []string{},
	}

	_, errs := newRuleChain(in.RuleChain.Payload)
	if len(errs) > 0 {
		for _, err := range errs {
			resp.Reasons = append(resp.Reasons, err.Error())
		}
		return &resp, status.Error(codes.InvalidArgument, "")
	}

	return &resp, nil
}

// CreateRuleChain add a new rulechain into  repository
func (s *RuleChainService) CreateRuleChain(ctx context.Context, in *pb.CreateRuleChainRequest) (*pb.CreateRuleChainResponse, error) {
	resp := pb.CreateRuleChainResponse{
		Reasons: []string{},
	}
	_, errs := newRuleChain(in.RuleChain.Payload)
	if len(errs) > 0 {
		for _, err := range errs {
			resp.Reasons = append(resp.Reasons, err.Error())
		}
		return &resp, status.Error(codes.InvalidArgument, "")
	}

	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.RuleChain.UserID)
	rulechain := converter.NewRuleChainModel(in.RuleChain)
	_, err := pf.Save(owner, rulechain)

	return &resp, xerror(err)
}

// DeleteRuleChain remove a rulechain from rulechain service
// In the cluster environmnent, the peer nodes should be notified
func (s *RuleChainService) DeleteRuleChain(ctx context.Context, in *pb.DeleteRuleChainRequest) (*pb.DeleteRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechain, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.DeleteRuleChainResponse{}, xerror(err)
	}
	// if rule chain's status is not allowed to be deleted, also return errors
	if rulechain.(*models.RuleChain).Status == models.RULE_STATUS_STARTED {
		return nil, status.Error(codes.FailedPrecondition, "")
	}

	if err := pf.Delete(owner, in.RuleChainID); err != nil {
		return nil, xerror(err)
	}
	notify(broadcast.OBJECT_DELETED, nameOfRuleChain, rulechain)
	return &pb.DeleteRuleChainResponse{}, nil
}

// UpdateRuleChain update an existed rule chain
func (s *RuleChainService) UpdateRuleChain(ctx context.Context, in *pb.UpdateRuleChainRequest) (*pb.UpdateRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.RuleChain.UserID)

	// If the rule chain no exist, just return error
	rulechain, err := pf.Get(owner, in.RuleChain.ID)
	if err != nil {
		return &pb.UpdateRuleChainResponse{}, xerror(err)
	}
	// if rule chain's status is not allowed to be deleted, also return errors
	if rulechain.(*models.RuleChain).Status == models.RULE_STATUS_STARTED {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	rulechainModel := converter.NewRuleChainModel(in.RuleChain)
	if err := pf.Update(owner, rulechainModel); err != nil {
		return nil, xerror(err)
	}
	notify(broadcast.OBJECT_UPDATED, nameOfRuleChain,
		&notifications.RuleChainNotification{
			UserID:      in.RuleChain.UserID,
			RuleChainID: rulechainModel.ID,
		})
	return &pb.UpdateRuleChainResponse{}, nil
}

// GetRuleChian return specified rulechain
func (s *RuleChainService) GetRuleChain(ctx context.Context, in *pb.GetRuleChainRequest) (*pb.GetRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechainModel, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.GetRuleChainResponse{}, xerror(err)
	}
	return &pb.GetRuleChainResponse{
		RuleChain: converter.NewRuleChain(rulechainModel),
	}, nil
}

// GetRuleChains returns user's all rulechain informations
func (s *RuleChainService) GetRuleChains(ctx context.Context, in *pb.GetRuleChainsRequest) (*pb.GetRuleChainsResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechainModels, err := pf.List(owner, models.NewQuery())
	if err != nil {
		return &pb.GetRuleChainsResponse{}, xerror(err)
	}
	return &pb.GetRuleChainsResponse{
		RuleChains: converter.NewRuleChains(rulechainModels),
	}, nil
}

// StartRuleChain start a rule chain to receive incoming data
func (s *RuleChainService) StartRuleChain(ctx context.Context, in *pb.StartRuleChainRequest) (*pb.StartRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechainModel, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.StartRuleChainResponse{}, xerror(err)
	}
	rulechain := rulechainModel.(*models.RuleChain)
	if rulechain.Status != models.RULE_STATUS_CREATED &&
		rulechain.Status != models.RULE_STATUS_STOPPED {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	notify(broadcast.OBJECT_UPDATED, nameOfRuleChain,
		&notifications.RuleChainNotification{
			UserID:      in.UserID,
			RuleChainID: in.RuleChainID,
		})
	return &pb.StartRuleChainResponse{}, nil
}

// StopRuleChain stop a rule chain to receive incoming data
func (s *RuleChainService) StopRuleChain(ctx context.Context, in *pb.StopRuleChainRequest) (*pb.StopRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechainModel, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.StopRuleChainResponse{}, xerror(err)
	}
	rulechain := rulechainModel.(*models.RuleChain)
	if rulechain.Status != models.RULE_STATUS_STARTED {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	notify(broadcast.OBJECT_UPDATED, nameOfRuleChain,
		&notifications.RuleChainNotification{
			UserID:      in.UserID,
			RuleChainID: in.RuleChainID,
		})
	return &pb.StopRuleChainResponse{}, nil
}

// xerror return grpc error according to models errors
func xerror(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, factory.ErrObjectNotFound) {
		return status.Errorf(codes.NotFound, "%w", err)
	} else if errors.Is(err, factory.ErrObjectAlreadyExist) {
		return status.Errorf(codes.AlreadyExists, "%w", err)
	} else if errors.Is(err, factory.ErrObjectInvalidArg) {
		return status.Errorf(codes.InvalidArgument, "%w", err)
	} else {
		return status.Errorf(codes.Internal, "%s", err)
	}
}
