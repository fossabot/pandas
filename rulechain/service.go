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
	"github.com/cloustone/pandas/headmast"
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/rulechain/options"
	"github.com/sirupsen/logrus"
)

// RuleChainService implement all rulechain interface
type RuleChainService struct {
	standaloneService
	instanceManager *instanceManager
	client          *headmast.Client
}

// NewRuleChainService return rulechain service object
func NewRuleChainService(servingOptions *options.ServingOptions) *RuleChainService {
	instanceManager := newInstanceManager(servingOptions)
	s := &RuleChainService{
		standaloneService: *newStandaloneService(servingOptions, instanceManager),
		instanceManager:   instanceManager,
	}

	if !servingOptions.IsStandalone() {
		opts := &headmast.ClientOptions{ServerAddr: servingOptions.HeadmastEndpoint}
		client := headmast.NewClient(opts)
		err1 := client.WatchJobPath("/headmast/workers/"+servingOptions.ServiceID+"/jobs", s.headmastJobAdded)
		err2 := client.WatchJobPath("/headmast/workers/"+servingOptions.ServiceID+"/killer", s.headmastJobDeleted)
		if err1 != nil || err2 != nil {
			logrus.Fatalf("watching headmast failed")
			return nil
		}
		s.client = client
	}
	return s
}

func (s *RuleChainService) headmastJobAdded(path string, job *headmast.Job) {
	rulechain := &models.RuleChain{}
	if err := rulechain.UnmarshalBinary(job.Payload); err != nil {
		logrus.WithError(err)
		return
	}
	s.instanceManager.startRuleChain(rulechain) // TODO: how to deal with failed job
}

func (s *RuleChainService) headmastJobDeleted(path string, job *headmast.Job) {
	rulechain := &models.RuleChain{}
	if err := rulechain.UnmarshalBinary(job.Payload); err != nil {
		logrus.WithError(err)
		return
	}
	s.instanceManager.deleteRuleChain(rulechain) // TODO: how to deal with failed job
}
