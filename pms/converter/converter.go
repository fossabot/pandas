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
package converter

import (
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/pms/grpc_pms_v1"
	"github.com/golang/protobuf/ptypes"
)

// NewProject return a grpc pms object converted from projec model
func NewProject(m interface{}) *grpc_pms_v1.Project {
	objectModel := m.(models.Project)
	createdAt, _ := ptypes.TimestampProto(objectModel.CreatedAt)
	lastUpdatedAt, _ := ptypes.TimestampProto(objectModel.LastUpdatedAt)

	object := &grpc_pms_v1.Project{
		ID:            objectModel.ID,
		Name:          objectModel.Name,
		UserID:        objectModel.UserID,
		Description:   objectModel.Description,
		Status:        objectModel.Status,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
	}
	// TODO: reader configuration should be added
	return object
}

// NewProjects return pms grpc objects that converted from project object models
func NewProjects(objectModels []models.Model) []*grpc_pms_v1.Project {
	projects := []*grpc_pms_v1.Project{}
	for _, objectModel := range objectModels {
		projects = append(projects, NewProject(objectModel))
	}
	return projects
}

// NewProjectModel return project model that converted from grpc object
func NewProjectModel(object *grpc_pms_v1.Project) *models.Project {
	createdAt, _ := ptypes.Timestamp(object.CreatedAt)
	lastUpdatedAt, _ := ptypes.Timestamp(object.LastUpdatedAt)
	model := &models.Project{
		ID:            object.ID,
		Name:          object.Name,
		UserID:        object.UserID,
		Description:   object.Description,
		Status:        object.Status,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
	}
	return model
}

// NewProjectModels return project models from grpc project objects
func NewProjectModels(objects []*grpc_pms_v1.Project) []models.Project {
	projectModels := []models.Project{}
	for _, object := range objects {
		projectModels = append(projectModels, *NewProjectModel(object))
	}
	return projectModels

}

// NewWorkshop return grpc workshop object
func NewWorkshop(m models.Model) *grpc_pms_v1.Workshop {
	ws := m.(*models.Workshop)
	createdAt, _ := ptypes.TimestampProto(ws.CreatedAt)
	lastUpdatedAt, _ := ptypes.TimestampProto(ws.LastUpdatedAt)

	return &grpc_pms_v1.Workshop{
		ID:            ws.ID,
		Name:          ws.Name,
		UserID:        ws.UserID,
		Description:   ws.Description,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		Status:        ws.Status,
		ViewIDs:       ws.ViewIDs,
	}
}

// NewWorkshops return grpc workshop objects
func NewWorkshops(workshopModels []models.Model) []*grpc_pms_v1.Workshop {
	workshops := []*grpc_pms_v1.Workshop{}
	for _, ws := range workshopModels {
		workshops = append(workshops, NewWorkshop(ws))
	}
	return workshops
}

// NewWorkshopModel return workshop model
func NewWorkshopModel(ws *grpc_pms_v1.Workshop) *models.Workshop {
	createdAt, _ := ptypes.Timestamp(ws.CreatedAt)
	lastUpdatedAt, _ := ptypes.Timestamp(ws.LastUpdatedAt)

	return &models.Workshop{
		ID:            ws.ID,
		Name:          ws.Name,
		UserID:        ws.UserID,
		Description:   ws.Description,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		Status:        ws.Status,
		ViewIDs:       ws.ViewIDs,
	}
}

// NewWorkshopModels return workshop models
func NeworskhopModels(workshops []*grpc_pms_v1.Workshop) []*models.Workshop {
	workshopModels := []*models.Workshop{}
	for _, ws := range workshops {
		workshopModels = append(workshopModels, NewWorkshopModel(ws))
	}
	return workshopModels

}

// NewView return grpc view object
func NewView(m models.Model) *grpc_pms_v1.View {
	v := m.(*models.View)
	createdAt, _ := ptypes.TimestampProto(v.CreatedAt)
	lastUpdatedAt, _ := ptypes.TimestampProto(v.LastUpdatedAt)

	return &grpc_pms_v1.View{
		ID:            v.ID,
		Name:          v.Name,
		ProjectID:     v.ProjectID,
		WorkshopID:    v.WorkshopID,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		Status:        v.Status,
		Variables:     v.Variables,
	}
}

// NewView return grpc view objects
func NewViews(viewModels []models.Model) []*grpc_pms_v1.View {
	views := []*grpc_pms_v1.View{}

	for _, viewModel := range viewModels {
		views = append(views, NewView(viewModel))
	}
	return views
}

// NewViewModel return view model
func NewViewModel(v *grpc_pms_v1.View) *models.View {
	createdAt, _ := ptypes.Timestamp(v.CreatedAt)
	lastUpdatedAt, _ := ptypes.Timestamp(v.LastUpdatedAt)

	return &models.View{
		ID:            v.ID,
		Name:          v.Name,
		ProjectID:     v.ProjectID,
		WorkshopID:    v.WorkshopID,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		Status:        v.Status,
		Variables:     v.Variables,
	}
}

// NewViewModels return view models
func NewViewModels(views []*grpc_pms_v1.View) []*models.View {
	viewModels := []*models.View{}
	for _, view := range views {
		viewModels = append(viewModels, NewViewModel(view))
	}
	return viewModels
}

// NewVariable return grpc variable object
func NewVariable(m interface{}) *grpc_pms_v1.Variable {
	v := m.(models.Variable)
	createdAt, _ := ptypes.TimestampProto(v.CreatedAt)
	lastUpdatedAt, _ := ptypes.TimestampProto(v.LastUpdatedAt)

	return &grpc_pms_v1.Variable{
		ID:             v.ID,
		Name:           v.Name,
		ProjectID:      v.ProjectID,
		Description:    v.Description,
		CreatedAt:      createdAt,
		LastUpdatedAt:  lastUpdatedAt,
		BindedDeviceID: v.BindedDeviceID,
		BindedEndpoint: v.BindedEndpoint,
		// TODO: variable value
	}
}

// NewVariables return grpc view objects
func NewVariables(variableModels []models.Model) []*grpc_pms_v1.Variable {
	variables := []*grpc_pms_v1.Variable{}
	for _, variableModel := range variableModels {
		variables = append(variables, NewVariable(variableModel))
	}
	return variables
}

// NewVariableModel return view model
func NewVariableModel(v *grpc_pms_v1.Variable) *models.Variable {
	createdAt, _ := ptypes.Timestamp(v.CreatedAt)
	lastUpdatedAt, _ := ptypes.Timestamp(v.LastUpdatedAt)

	return &models.Variable{
		ID:             v.ID,
		Name:           v.Name,
		ProjectID:      v.ProjectID,
		Description:    v.Description,
		CreatedAt:      createdAt,
		LastUpdatedAt:  lastUpdatedAt,
		BindedDeviceID: v.BindedDeviceID,
		BindedEndpoint: v.BindedEndpoint,
	}
}

// NewVariableModels return variable models
func NewVariableModels(variables []*grpc_pms_v1.Variable) []*models.Variable {
	variableModels := []*models.Variable{}
	for _, variable := range variables {
		variableModels = append(variableModels, NewVariableModel(variable))
	}
	return variableModels
}
