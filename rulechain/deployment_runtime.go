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
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/plugins"
	"github.com/sirupsen/logrus"
)

type deploymentRuntime struct {
	models.Deployment
	reader    models.Reader
	configs   map[string]interface{}
	rulechain RuleChain
	plugin    plugins.Plugin
}

func newDeploymentRuntime(model models.Deployment) (*deploymentRuntime, error) {
	/*
		reader, err := readers.NewReader(model.Reader, model.ReaderConfigs)
		if err != nil {
			return nil, err
		}
		ruleChainModel, err := models.GetRuleChain(model.UserId, model.RuleChainId)
		if err != nil {
			return nil, err
		}
		rulechain, err := NewRuleChain(ruleChainModel.Metadata)
		if err != nil {
			return nil, err
		}
		plugin, err := plugins.LoadPlugin(model.ReaderConfigs)
		if err != nil {
			return nil, err
		}

		r := &deploymentRuntime{
			reader:    reader,
			rulechain: rulechain,
			plugin:    plugin,
		}
		reader.RegisterObserver(r)
		return r, nil
	*/
	return nil, nil
}

func (r *deploymentRuntime) start() error {
	go func(r *deploymentRuntime) {
		r.reader.Start()
	}(r)
	return nil
}

func (r *deploymentRuntime) stop() error {
	err := r.reader.GracefulShutdown()
	return err
}

func (r *deploymentRuntime) OnDataAvailable(reader models.Reader, payload []byte, param interface{}) {
	msg, err := r.plugin.ConstructMessage(payload)
	if err != nil {
		logrus.WithError(err)
		return
	}

	go func(r *deploymentRuntime, msg models.Message) {
		r.rulechain.ApplyMessage(msg)
	}(r, msg)
}
