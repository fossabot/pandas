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
	"github.com/cloustone/pandas/rulechain/converter"
	pb "github.com/cloustone/pandas/rulechain/grpc_rulechain_v1"
	"github.com/cloustone/pandas/rulechain/nodes"
	"github.com/cloustone/pandas/rulechain/options"
	logr "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	nameOfRuleChain = reflect.TypeOf(models.RuleChain{}).Name()
)

// standaloneService implement all rulechain interface
type standaloneService struct {
	servingOptions  *options.ServingOptions
	instanceManager *instanceManager
}

// NewstandaloneService return rulechain service object
func newStandaloneService(servingOptions *options.ServingOptions, instanceManager *instanceManager) *standaloneService {
	s := &standaloneService{
		servingOptions:  servingOptions,
		instanceManager: instanceManager,
	}
	return s
}

// loadAllRuleChains load runtimes in models and deploy them according to rulechain's status
func (s *standaloneService) loadAllRuleChains() error {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner("") // TODO
	query := models.NewQuery().WithQuery("status", models.RULE_STATUS_STARTED)
	rulechainModels, err := pf.List(owner, query)
	if err != nil {
		logr.WithError(err)
		return err
	}
	for _, rulechainModel := range rulechainModels {
		rulechain := rulechainModel.(*models.RuleChain)
		if err := s.instanceManager.startRuleChain(rulechain); err != nil {
			logr.WithError(err)
		}
	}
	return nil
}

// The following is standalone service methods

// CheckRuleChain check wether the rule chain is valid
func (s *standaloneService) CheckRuleChain(ctx context.Context, in *pb.CheckRuleChainRequest) (*pb.CheckRuleChainResponse, error) {
	resp := pb.CheckRuleChainResponse{
		Reasons: []string{},
	}

	_, errs := newRuleChainInstance(in.RuleChain.Payload)
	if len(errs) > 0 {
		for _, err := range errs {
			resp.Reasons = append(resp.Reasons, err.Error())
		}
		return &resp, status.Error(codes.InvalidArgument, "")
	}
	return &resp, nil
}

// CreateRuleChain add a new rulechain into  repository
func (s *standaloneService) CreateRuleChain(ctx context.Context, in *pb.CreateRuleChainRequest) (*pb.CreateRuleChainResponse, error) {
	resp := pb.CreateRuleChainResponse{
		Reasons: []string{},
	}
	_, errs := newRuleChainInstance(in.RuleChain.Payload)
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
func (s *standaloneService) DeleteRuleChain(ctx context.Context, in *pb.DeleteRuleChainRequest) (*pb.DeleteRuleChainResponse, error) {
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
	return &pb.DeleteRuleChainResponse{}, nil
}

// UpdateRuleChain update an existed rule chain
func (s *standaloneService) UpdateRuleChain(ctx context.Context, in *pb.UpdateRuleChainRequest) (*pb.UpdateRuleChainResponse, error) {
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
	return &pb.UpdateRuleChainResponse{}, nil
}

// GetRuleChian return specified rulechain
func (s *standaloneService) GetRuleChain(ctx context.Context, in *pb.GetRuleChainRequest) (*pb.GetRuleChainResponse, error) {
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
func (s *standaloneService) GetRuleChains(ctx context.Context, in *pb.GetRuleChainsRequest) (*pb.GetRuleChainsResponse, error) {
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
func (s *standaloneService) StartRuleChain(ctx context.Context, in *pb.StartRuleChainRequest) (*pb.StartRuleChainResponse, error) {
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
	return &pb.StartRuleChainResponse{}, nil
}

// StopRuleChain stop a rule chain to receive incoming data
func (s *standaloneService) StopRuleChain(ctx context.Context, in *pb.StopRuleChainRequest) (*pb.StopRuleChainResponse, error) {
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
	return &pb.StopRuleChainResponse{}, nil
}

// GetNodeConfigs return all nodes' configs
func (s *standaloneService) GetNodeConfigs(ctx context.Context, in *pb.GetNodeConfigsRequest) (*pb.GetNodeConfigsResponse, error) {
	nodeConfigs := []*pb.NodeConfig{}
	categoryNodes := nodes.GetCategoryNodes()
	allNodeConfigs := nodes.GetAllNodeConfigs()

	for category, nodes := range categoryNodes {
		for _, nodeType := range nodes {
			if nodeConfig, found := allNodeConfigs[nodeType]; found {
				nodeConfigs = append(nodeConfigs, &pb.NodeConfig{
					Type:     nodeType,
					Category: category,
					Payload:  []byte(nodeConfig),
				})
			}
		}
	}
	return &pb.GetNodeConfigsResponse{NodeConfigs: nodeConfigs}, nil
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
