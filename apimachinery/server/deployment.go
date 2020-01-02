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
package server

import (
	"github.com/cloustone/pandas/apimachinery/restapi/operations/deployment"
	"github.com/cloustone/pandas/models"

	"github.com/go-openapi/runtime/middleware"
)

func GetDeployments(param deployment.GetDeploymentsParams, principal *models.Principal) middleware.Responder {
	return &deployment.GetDeploymentsOK{}
}

func CreateDeployment(params deployment.CreateDeploymentParams, principal *models.Principal) middleware.Responder {
	return &deployment.CreateDeploymentOK{}
}

func DeleteDeployment(params deployment.DeleteDeploymentParams, principal *models.Principal) middleware.Responder {
	return &deployment.DeleteDeploymentOK{}
}

func GetDeployment(params deployment.GetDeploymentParams, principal *models.Principal) middleware.Responder {
	return &deployment.GetDeploymentOK{}
}
func SetDeploymentStatus(params deployment.SetDeploymentStatusParams, principal *models.Principal) middleware.Responder {
	return &deployment.SetDeploymentStatusOK{}
}
func UpdateDeployment(params deployment.UpdateDeploymentParams, principal *models.Principal) middleware.Responder {
	return &deployment.UpdateDeploymentOK{}
}
