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
	"fmt"
	"sync"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/rulechain/adaptors"
	"github.com/cloustone/pandas/rulechain/options"

	logr "github.com/sirupsen/logrus"
)

// instanceManager manage all rulechain's runtime
type instanceManager struct {
	servingOptions *options.ServingOptions
	mutex          sync.RWMutex
	rulechains     map[string]*ruleChainInstance
	adaptors       map[string][]string
}

// newInstanceManager create controller instance used in rule chain service
func newInstanceManager(servingOptions *options.ServingOptions) *instanceManager {
	controller := &instanceManager{
		servingOptions: servingOptions,
		mutex:          sync.RWMutex{},
		rulechains:     make(map[string]*ruleChainInstance),
		adaptors:       make(map[string][]string),
	}
	return controller
}

// getAdaptorRuleChains return all rule chains that handle incomming data from
// specified adaptors
func (r *instanceManager) getAdaptorRuleChains(adaptorID string) []*ruleChainInstance {
	rulechains := []*ruleChainInstance{}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, rulechainID := range r.adaptors[adaptorID] {
		rulechains = append(rulechains, r.rulechains[rulechainID])
	}
	return rulechains
}

// buildAdaptorOptions
func buildAdaptorOptions(c *models.DataSource) *adaptors.AdaptorOptions {
	return &adaptors.AdaptorOptions{
		Name:         c.Name,
		Protocol:     c.Protocol,
		IsProvider:   c.IsProvider,
		ServicePort:  c.ServicePort,
		ConnectURL:   c.ConnectURL,
		IsTLSEnabled: c.IsTLSEnabled,
		KeyFile:      c.KeyFile,
		CertFile:     c.CertFile,
	}
}

// startRuleChain start the rule chain and receiving incoming data
func (r *instanceManager) startRuleChain(rulechainModel *models.RuleChain) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, found := r.rulechains[rulechainModel.ID]; found {
		logr.Debugf("rule chain '%s' is already started", rulechainModel.ID)
		return nil
	}
	// create the internal runtime rulechain
	rulechain, errs := newRuleChainInstance(rulechainModel.Payload)
	if len(errs) > 0 {
		return errs[0]
	}

	adaptorOptions := buildAdaptorOptions(&rulechainModel.DataSource)
	r.addInstanceInternal(rulechainModel.ID, rulechain, adaptorOptions.Name)
	return nil
}

// addInstanceInternal add a new rulechain instance internally with
// specified adaptor id
func (r *instanceManager) addInstanceInternal(rulechainID string, instance *ruleChainInstance, adaptorID string) {
	if _, found := r.adaptors[adaptorID]; !found {
		r.adaptors[adaptorID] = []string{}
	}
	r.adaptors[adaptorID] = append(r.adaptors[adaptorID], rulechainID)
	r.rulechains[rulechainID] = instance
}

// stopRuleChain stop the rule chain
func (r *instanceManager) stopRuleChain(rulechainModel *models.RuleChain) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, found := r.rulechains[rulechainModel.ID]; !found {
		logr.Debugf("rule chain '%s' is not found", rulechainModel.ID)
		return fmt.Errorf("rule chain '%s' no exist", rulechainModel.ID)
	}
	delete(r.rulechains, rulechainModel.ID)
	return nil
}

// deleteRuleChain remove rule chain
func (c *instanceManager) deleteRuleChain(rulechain *models.RuleChain) error {
	return nil
}
