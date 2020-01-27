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
	"github.com/cloustone/pandas/models/factory"
	"github.com/cloustone/pandas/models/notifications"
	"github.com/cloustone/pandas/pkg/broadcast"
	broadcast_util "github.com/cloustone/pandas/pkg/broadcast/util"
	"github.com/cloustone/pandas/pkg/readers"

	logr "github.com/sirupsen/logrus"
)

// runtimeController manage all rulechain's runtime
type runtimeController struct {
	mutex      sync.RWMutex
	rulechains map[string]*ruleChain
}

// newRuntimeController create controller instance used in rule chain service
func newRuntimeController() *runtimeController {
	controller := &runtimeController{
		mutex:      sync.RWMutex{},
		rulechains: make(map[string]*ruleChain),
	}
	broadcast_util.RegisterObserver(controller, nameOfRuleChain)
	return controller
}

// OnBroadcase will be notified when rulechain model object is changed
func (r *runtimeController) Onbroadcast(b broadcast.Broadcast, notify broadcast.Notification) {
	rulechainNotify := notify.Param.(notifications.RuleChainNotification)
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner(rulechainNotify.UserID)

	rulechainModel, err := pf.Get(owner, rulechainNotify.RuleChainID)
	if err != nil {
		logr.WithError(err)
		return
	}
	rulechain := rulechainModel.(*models.RuleChain)

	switch notify.Action {
	case broadcast.ActionCreated:
	case broadcast.ActionUpdated:
		switch rulechain.Status {
		case models.RuleStatusStarted:
			err = r.startRuleChain(rulechain)
		case models.RuleStatusStopped:
			err = r.stopRuleChain(rulechain)
		default:
			err = fmt.Errorf("invalid runtime status '%s'", rulechain.Status)
		}

	case broadcast.ActionDeleted:
		err = r.deleteRuleChain(rulechain)
	default:
		err = fmt.Errorf("invalid model action '%s'", notify.Action)
	}
	logr.WithError(err)
}

// loadAllRuleChains load runtimes in models and deploy them according to rulechain's status
func (r *runtimeController) loadAllRuleChains() error {
	pf := factory.NewFactory(models.RuleChain{})
	owner := factory.NewOwner("") // TODO
	query := models.NewQuery().WithQuery("status", models.RuleStatusStarted)
	rulechainModels, err := pf.List(owner, query)
	if err != nil {
		logr.WithError(err)
		return err
	}
	for _, rulechainModel := range rulechainModels {
		rulechain := rulechainModel.(*models.RuleChain)
		if err := r.startRuleChain(rulechain); err != nil {
			logr.WithError(err)
		}
	}
	return nil
}

// startRuleChain start the rule chain and receiving incoming data
func (r *runtimeController) startRuleChain(rulechainModel *models.RuleChain) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, found := r.rulechains[rulechainModel.ID]; found {
		logr.Debugf("rule chain '%s' is already started", rulechainModel.ID)
		return nil
	}
	// create the internal runtime rulechain
	rulechain, errs := newRuleChain(rulechainModel.Payload)
	if len(errs) > 0 {
		return errs[0]
	}
	// create reader for the rule chain
	reader, err := readers.NewReader("", nil) // todo
	if err != nil {
		logr.WithError(err)
		return err
	}
	reader.RegisterObserver(rulechain)
	r.rulechains[rulechainModel.ID] = rulechain
	go reader.Start()
	return nil
}

// stopRuleChain stop the rule chain
func (r *runtimeController) stopRuleChain(rulechainModel *models.RuleChain) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rulechain, found := r.rulechains[rulechainModel.ID]
	if !found {
		logr.Debugf("rule chain '%s' is not found", rulechainModel.ID)
		return fmt.Errorf("rule chain '%s' no exist", rulechainModel.ID)
	}
	delete(r.rulechains, rulechainModel.ID)
	rulechain.reader.GracefulShutdown()
	return nil
}

// deleteRuleChain remove rule chain
func (c *runtimeController) deleteRuleChain(rulechain *models.RuleChain) error {
	return nil
}
