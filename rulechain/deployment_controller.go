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
	"github.com/cloustone/pandas/pkg/synchron"

	logr "github.com/sirupsen/logrus"
)

const PATH = "deployments"

type deploymentController struct {
	mutex       sync.RWMutex
	deployments map[string]*deploymentRuntime
}

func newDeploymentController() *deploymentController {
	controller := &deploymentController{
		mutex:       sync.RWMutex{},
		deployments: make(map[string]*deploymentRuntime),
	}

	//models.RegisterModelObserver(PATH, controller)
	return controller
}

func (c *deploymentController) Shutdown() {}

func (c *deploymentController) OnModelNotified(path string, action string, obj interface{}) {
	var err error = nil
	deployment := obj.(models.Deployment)

	switch action {
	case synchron.ActionCreated:
	case synchron.ActionUpdated:
		switch deployment.Status {
		case models.DeploymentStatusRunning:
			err = c.startDeployment(deployment)
		case models.DeploymentStatusStopped:
			err = c.stopDeployment(deployment)
		default:
			err = fmt.Errorf("invalid deployment status '%s'", deployment.Status)
		}

	case synchron.ActionDeleted:
		err = c.deleteDeployment(deployment)
	default:
		err = fmt.Errorf("invalid model action '%s'", action)
	}
	logr.WithError(err)
}

// loadAllDeployments load deployments in models and deploy them according
// to deployment's status
func (c *deploymentController) loadAllDeployments() error {
	/*
		deployments, err := models.ListDeployments(models.Query{})
		if err != nil {
			return err
		}
		for _, deployment := range deployments {
			// creating and add the new deployments in this controller
			r, err := newDeploymentRuntime(*deployment)
			if err != nil {
				logr.WithError(err).Errorf("create user '%s' deployment '%s' failed", deployment.UserId, deployment.Id)
				continue
			}
			c.mutex.Lock()
			c.deployments[r.Id] = r
			c.mutex.Unlock()

			switch deployment.Status {
			case models.DEPLOYMENT_STATUS_RUNNING:
				if err := c.startDeployment(*deployment); err != nil {
					logr.WithError(err).Errorf("start user '%s' deployment '%s' failed", deployment.UserId, deployment.Id)
				}
				break
			case models.DEPLOYMENT_STATUS_CREATED:
			case models.DEPLOYMENT_STATUS_STOPPED:
				break
			default:
				logr.Errorf("unknown deployment '%s' status found", deployment.Id)
				break
			}
		}
	*/
	return nil
}

func (c *deploymentController) startDeployment(deployment models.Deployment) error {
	// Add deployment in case the runtime deployment is not sychronized till now
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.deployments[deployment.ID] == nil {
		r, err := newDeploymentRuntime(deployment)
		if err != nil {
			return err
		}
		c.deployments[deployment.ID] = r
	}

	deploymentRuntime := c.deployments[deployment.ID]

	switch deployment.Status {
	case models.DeploymentStatusCreated:
		if err := deploymentRuntime.start(); err == nil {
			deploymentRuntime.Status = models.DeploymentStatusRunning
			return nil
		}
	}

	return fmt.Errorf("invalid delopment '%s' status", deployment.ID)
}

func (c *deploymentController) stopDeployment(deployment models.Deployment) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, found := c.deployments[deployment.ID]; !found {
		return fmt.Errorf("deployment '%s' not found", deployment.ID)
	}

	switch deployment.Status {
	case models.DeploymentStatusRunning:
		c.deployments[deployment.ID].stop()
		c.deployments[deployment.ID].Status = models.DeploymentStatusStopped
	default:
		return fmt.Errorf("invalid delopment '%s' status", deployment.ID)
	}
	return nil
}

func (c *deploymentController) deleteDeployment(deployment models.Deployment) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, found := c.deployments[deployment.ID]; !found {
		return fmt.Errorf("deployment '%s' not found", deployment.ID)
	}

	switch deployment.Status {
	case models.DeploymentStatusRunning:
		c.deployments[deployment.ID].stop()
		delete(c.deployments, deployment.ID)
	default:
		return fmt.Errorf("invalid delopment '%s' status", deployment.ID)
	}
	return nil
}
