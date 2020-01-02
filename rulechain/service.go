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

	"github.com/cloustone/pandas/pkg/synchron"
	syncutil "github.com/cloustone/pandas/pkg/synchron/util"
	pb "github.com/cloustone/pandas/rulechain/grpc_rulechain_v1"
	logr "github.com/sirupsen/logrus"
)

var (
	synchronizer synchron.Synchronizer
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

// Initialize will add prestart behaivor such as synchronization initialization
func (s *RuleChainService) Initialize(syncOptions *synchron.SyncServingOptions) {
	s.controllers = loadControllers()
	synchronizer = syncutil.NewSynchronizer(syncOptions).WithRootPath("pandas/rulechain")
	synchronizer.AsMember()
	for path, _ := range s.controllers {
		synchronizer.RegisterObserver(path, s)
	}
}

// OnSynchonronizedNotified will be notified when rulechain model object is changed
func (s *RuleChainService) OnSynchronizationNotified(sync synchron.Synchronizer, notify synchron.Notification) {
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

// notify is internal helper to simplify synchronization notificaiton
func notify(action string, path string) {
	synchronizer.Notify(synchron.Notification{
		Path:   path,
		Action: action,
		Param:  path,
	})
}

func (s *RuleChainService) CheckRuleChain(context.Context, *pb.CheckRuleChainRequest) (*pb.CheckRuleChainResponse, error) {
	return nil, nil
}

func (s *RuleChainService) CreateRuleChain(context.Context, *pb.CreateRuleChainRequest) (*pb.CreateRuleChainResponse, error) {
	return nil, nil
}

func (s *RuleChainService) DeleteRuleChain(context.Context, *pb.DeleteRuleChainRequest) (*pb.DeleteRuleChainResponse, error) {
	return nil, nil
}

func (s *RuleChainService) UpdateRuleChain(context.Context, *pb.UpdateRuleChainRequest) (*pb.UpdateRuleChainResponse, error) {
	return nil, nil
}

func (s *RuleChainService) GetRuleChain(context.Context, *pb.GetRuleChainRequest) (*pb.GetRuleChainResponse, error) {
	return nil, nil
}

func (s *RuleChainService) GetUserRuleChains(context.Context, *pb.GetUserRuleChainsRequest) (*pb.GetUserRuleChainsResponse, error) {
	return nil, nil
}

func (s *RuleChainService) StartRuleChain(context.Context, *pb.StartRuleChainRequest) (*pb.StartRuleChainResponse, error) {
	return nil, nil
}

func (s *RuleChainService) StopRuleChain(context.Context, *pb.StopRuleChainRequest) (*pb.StopRuleChainResponse, error) {
	return nil, nil
}
