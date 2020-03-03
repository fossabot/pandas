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
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

type headmastObserver interface {
	onRulechainAdded(r *models.RuleChain) error
	onRulechainDeleted(r *models.RuleChain) error
}

type headmastConnector struct {
	client   *headmast.Client
	clientID string
	observer headmastObserver
}

func newHeadmastConnector(servingOptions *options.ServingOptions) *headmastConnector {
	opts := &headmast.ClientOptions{ServerAddr: servingOptions.HeadmastEndpoint}
	c := &headmastConnector{
		client:   headmast.NewClient(opts),
		clientID: xid.New().String(),
	}
	err1 := c.client.WatchJobPath("/headmast/workers/"+c.clientID+"/jobs", c.headmastJobAdded)
	err2 := c.client.WatchJobPath("/headmast/workers/"+c.clientID+"/killer", c.headmastJobDeleted)
	if err1 != nil || err2 != nil {
		logrus.Fatalf("watching headmast failed")
		return nil
	}
	return c
}

func (c *headmastConnector) registerObserver(obs headmastObserver) {
	c.observer = obs
}

func (c *headmastConnector) headmastJobAdded(path string, job *headmast.Job) {
	rulechain := &models.RuleChain{}
	if err := rulechain.UnmarshalBinary(job.Payload); err != nil {
		logrus.WithError(err)
		return
	}
	c.observer.onRulechainAdded(rulechain) // TODO: how to deal with failed job
}

func (c *headmastConnector) headmastJobDeleted(path string, job *headmast.Job) {
	rulechain := &models.RuleChain{}
	if err := rulechain.UnmarshalBinary(job.Payload); err != nil {
		logrus.WithError(err)
		return
	}
	c.observer.onRulechainDeleted(rulechain) // TODO: how to deal with failed job
}
