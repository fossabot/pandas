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
	"github.com/cloustone/pandas/apimachinery/restapi/operations/project"
	"github.com/cloustone/pandas/models"
	serverconverter "github.com/cloustone/pandas/pms/converter"
	pb "github.com/cloustone/pandas/pms/grpc_pms_v1"

	"github.com/go-openapi/runtime/middleware"
)

func GetProjects(params project.GetProjectsParams, principal *models.Principal) middleware.Responder {
	client, err := pb.NewClient()
	if err != nil {
		return serverError(err)
	}
	defer client.Close()

	req := &pb.GetProjectsRequest{UserID: principal.ID}
	resp, err := client.ProjectManager().GetProjects(params.HTTPRequest.Context(), req)
	if err != nil {
		return serverError(err)
	}
	return project.NewGetProjectsOK().WithPayload(
		serverconverter.NewProjectModels(resp.Projects),
	)
}

func CreateProject(params project.CreateProjectParams, principal *models.Principal) middleware.Responder {
	client, err := pb.NewClient()
	if err != nil {
		return serverError(err)
	}
	defer client.Close()

	req := &pb.CreateProjectRequest{
		UserID:  principal.ID,
		Project: serverconverter.NewProject(params.Project),
	}
	resp, err := client.ProjectManager().CreateProject(params.HTTPRequest.Context(), req)
	if err != nil {
		return serverError(err)
	}
	return project.NewCreateProjectOK().WithPayload(*serverconverter.NewProjectModel(resp.Project))
}
