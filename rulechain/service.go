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
	"github.com/cloustone/pandas/pkg/broadcast"
	broadcast_util "github.com/cloustone/pandas/pkg/broadcast/util"
	"github.com/cloustone/pandas/rulechain/converter"
	pb "github.com/cloustone/pandas/rulechain/grpc_rulechain_v1"
	logr "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	nameOfRuleChain = reflect.TypeOf(models.RuleChain{}).Name()
)

// Controller monitor rulechain's change and adjust the deployment dynamically
type Controller interface {
	OnModelNotified(path string, reason string, obj interface{})
	Shutdown()
}

// RuleChainService implement all rulechain interface
type RuleChainService struct {
	controllers map[string]Controller
}

// NewRuleChainService return rulechain service object
func NewRuleChainService() *RuleChainService {
	return &RuleChainService{
		controllers: make(map[string]Controller),
	}
}

// Initialize will add prestart behaivor such as broadcastization initialization
func (s *RuleChainService) Initialize(options *broadcast.ServingOptions) {
	s.controllers = loadControllers()
	for path, _ := range s.controllers {
		broadcast_util.RegisterObserver(s, path)
	}
}

// OnSynchonronizedNotified will be notified when rulechain model object is changed
func (s *RuleChainService) Onbroadcast(b broadcast.Broadcast, notify broadcast.Notification) {
	if controller, found := s.controllers[notify.Path]; found {
		controller.OnModelNotified(notify.Path, notify.Action, notify.Param)
		return
	}
	logr.Errorf("no observer existed for model '%s'", notify.Path)
}

// loadControllers will create controllers according to evnrionment's setting
// in future, the loaded controllers can be configured with config file
func loadControllers() map[string]Controller {
	controllers := map[string]Controller{
		"deployments": newDeploymentController(),
	}
	return controllers
}

// notify is internal helper to simplify broadcastization notificaiton
func notify(action string, path string, param interface{}) {
	broadcast_util.Notify(action, path, param)
}

// CheckRuleChain check wether the rule chain is valid
func (s *RuleChainService) CheckRuleChain(ctx context.Context, in *pb.CheckRuleChainRequest) (*pb.CheckRuleChainResponse, error) {
	resp := pb.CheckRuleChainResponse{
		Reasons: []string{},
	}

	_, errs := NewRuleChain(in.RuleChain.Payload)
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
	_, errs := NewRuleChain(in.RuleChain.Payload)
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

	return &resp, grpcError(err)
}

// DeleteRuleChain remove a rulechain from rulechain service
// In the cluster environmnent, the peer nodes should be notified
func (s *RuleChainService) DeleteRuleChain(ctx context.Context, in *pb.DeleteRuleChainRequest) (*pb.DeleteRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechain, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.DeleteRuleChainResponse{}, grpcError(err)
	}
	// if rule chain's status is not allowed to be deleted, also return errors
	if rulechain.(*models.RuleChain).Status == models.RuleStatusStarted {
		return nil, status.Error(codes.FailedPrecondition, "")
	}

	if err := pf.Delete(owner, in.RuleChainID); err != nil {
		return nil, grpcError(err)
	}
	notify(broadcast.ActionDeleted, nameOfRuleChain, rulechain)
	return &pb.DeleteRuleChainResponse{}, nil
}

// UpdateRuleChain update an existed rule chain
func (s *RuleChainService) UpdateRuleChain(ctx context.Context, in *pb.UpdateRuleChainRequest) (*pb.UpdateRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.RuleChain.UserID)

	// If the rule chain no exist, just return error
	rulechain, err := pf.Get(owner, in.RuleChain.ID)
	if err != nil {
		return &pb.UpdateRuleChainResponse{}, grpcError(err)
	}
	// if rule chain's status is not allowed to be deleted, also return errors
	if rulechain.(*models.RuleChain).Status == models.RuleStatusStarted {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	rulechainModel := converter.NewRuleChainModel(in.RuleChain)
	if err := pf.Update(owner, rulechainModel); err != nil {
		return nil, grpcError(err)
	}
	notify(broadcast.ActionUpdated, nameOfRuleChain, rulechainModel)
	return &pb.UpdateRuleChainResponse{}, nil
}

// GetRuleChian return specified rulechain
func (s *RuleChainService) GetRuleChain(ctx context.Context, in *pb.GetRuleChainRequest) (*pb.GetRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechainModel, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.GetRuleChainResponse{}, grpcError(err)
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
		return &pb.GetRuleChainsResponse{}, grpcError(err)
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
		return &pb.StartRuleChainResponse{}, grpcError(err)
	}
	rulechain := rulechainModel.(*models.RuleChain)
	if rulechain.Status != models.RuleStatusCreated ||
		rulechain.Status != models.RuleStatusStopped {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	notify(broadcast.ActionUpdated, nameOfRuleChain, rulechain)
	return &pb.StartRuleChainResponse{}, nil
}

// StopRuleChain stop a rule chain to receive incoming data
func (s *RuleChainService) StopRuleChain(ctx context.Context, in *pb.StopRuleChainRequest) (*pb.StopRuleChainResponse, error) {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(in.UserID)

	// If the rule chain no exist, just return error
	rulechainModel, err := pf.Get(owner, in.RuleChainID)
	if err != nil {
		return &pb.StopRuleChainResponse{}, grpcError(err)
	}
	rulechain := rulechainModel.(*models.RuleChain)
	if rulechain.Status != models.RuleStatusStarted {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	notify(broadcast.ActionUpdated, nameOfRuleChain, rulechain)
	return &pb.StopRuleChainResponse{}, nil
}

// grpcError return grpc error according to models errors
func grpcError(err error) error {
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
