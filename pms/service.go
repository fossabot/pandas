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
package pms

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/factory"
	modeloptions "github.com/cloustone/pandas/models/options"
	"github.com/cloustone/pandas/pms/converter"
	pb "github.com/cloustone/pandas/pms/grpc_pms_v1"
)

// ProjectManagementService implement grpc service for pms
type ProjectManagementService struct{}

// NewProjectManagementService return service instance used in main server
func NewProjectManagementService(servingOptions *modeloptions.ServingOptions) *ProjectManagementService {
	factory.RegisterFactory(models.Project{}, newProjectFactory(servingOptions))
	factory.RegisterFactory(models.Workshop{}, newWorkshopFactory(servingOptions))
	factory.RegisterFactory(models.View{}, newViewFactory(servingOptions))

	return &ProjectManagementService{}
}

// grpcError return grpc error according to models errors
func grpcError(err error) error {
	if errors.Is(err, factory.ErrObjectNotFound) {
		return status.Errorf(codes.NotFound, "%w", err)

	} else if errors.Is(err, factory.ErrObjectAlreadyExist) {
		return status.Errorf(codes.AlreadyExists, "%w", err)

	} else if errors.Is(err, factory.ErrObjectInvalidArg) {
		return status.Errorf(codes.InvalidArgument, "%w", err)

	} else {
		return status.Errorf(codes.Internal, "%s", err)
	}
}

// CreateProject create a new project
func (s *ProjectManagementService) CreateProject(ctx context.Context, in *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	pf := factory.NewFactory(models.Project{})
	owner := factory.NewOwner(in.UserID)
	query := models.NewQuery().WithQuery("projectName", in.Project.Name)

	if _, err := pf.List(owner, query); err == nil {
		return nil, grpcError(err)
	}
	project, err := pf.Save(owner, converter.NewProjectModel(in.Project))
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.CreateProjectResponse{Project: converter.NewProject(project)}, nil
}

// GetProject return specified project detail
func (s *ProjectManagementService) GetProject(ctx context.Context, in *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	pf := factory.NewFactory(models.Project{})
	owner := factory.NewOwner(in.UserID)

	project, err := pf.Get(owner, in.ProjectID)
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetProjectResponse{Project: converter.NewProject(project)}, nil
}

// GetProjects return user's all projects
func (s *ProjectManagementService) GetProjects(ctx context.Context, in *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	pf := factory.NewFactory(models.Project{})
	owner := factory.NewOwner(in.UserID)

	projects, err := pf.List(owner, models.NewQuery())
	if err != nil {
		return nil, grpcError(err)
	}

	return &pb.GetProjectsResponse{Projects: converter.NewProjects(projects)}, nil
}

// DeleteProject delete specified project
func (s *ProjectManagementService) DeleteProject(ctx context.Context, in *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	pf := factory.NewFactory(models.Project{})
	owner := factory.NewOwner(in.UserID)

	if err := pf.Delete(owner, in.ProjectID); err != nil {
		return nil, grpcError(err)
	}
	return &pb.DeleteProjectResponse{}, nil
}

// UpdateProject update specified project
func (s *ProjectManagementService) UpdateProject(ctx context.Context, in *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	pf := factory.NewFactory(models.Project{})
	owner := factory.NewOwner(in.UserID)

	if _, err := pf.Get(owner, in.Project.ID); err != nil {
		return nil, grpcError(err)
	}
	if err := pf.Update(owner, converter.NewProjectModel(in.Project)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.UpdateProjectResponse{}, nil
}

// AddDevice add a device into the project
func (s *ProjectManagementService) AddDevice(ctx context.Context, in *pb.AddDeviceRequest) (*pb.AddDeviceResponse, error) {
	pf := factory.NewFactory(models.DeviceInProject{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)
	device := models.DeviceInProject{
		UserID:    in.UserID,
		ProjectID: in.ProjectID,
		DeviceID:  in.DeviceID,
	}
	if _, err := pf.Save(owner, &device); err != nil {
		return nil, grpcError(err)
	}
	return &pb.AddDeviceResponse{}, nil
}

// AddDevices add a batch of devices into the project
func (s *ProjectManagementService) AddDevices(ctx context.Context, in *pb.AddDevicesRequest) (*pb.AddDevicesResponse, error) {
	pf := factory.NewFactory(models.DeviceInProject{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	for _, deviceID := range in.DeviceIDs {
		device := models.DeviceInProject{
			UserID:    in.UserID,
			ProjectID: in.ProjectID,
			DeviceID:  deviceID,
		}
		if _, err := pf.Save(owner, &device); err != nil {
			return nil, grpcError(err)
		}
	}

	return &pb.AddDevicesResponse{}, nil
}

// DeleteDevice remove a device from project
func (s *ProjectManagementService) DeleteDevice(ctx context.Context, in *pb.DeleteDeviceRequest) (*pb.DeleteDeviceResponse, error) {
	pf := factory.NewFactory(models.DeviceInProject{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)
	if err := pf.Delete(owner, in.DeviceID); err != nil {
		return nil, grpcError(err)
	}

	return &pb.DeleteDeviceResponse{}, nil
}

// DeleteDevices remove a batch of devices from project
func (s *ProjectManagementService) DeleteDevices(ctx context.Context, in *pb.DeleteDevicesRequest) (*pb.DeleteDevicesResponse, error) {
	pf := factory.NewFactory(models.DeviceInProject{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	for _, deviceID := range in.DeviceIDs {
		if err := pf.Delete(owner, deviceID); err != nil {
			return nil, grpcError(err)
		}
	}
	return &pb.DeleteDevicesResponse{}, nil
}

// GetDevices return a project's all devices
func (s *ProjectManagementService) GetDevices(ctx context.Context, in *pb.GetDevicesRequest) (*pb.GetDevicesResponse, error) {
	pf := factory.NewFactory(models.DeviceInProject{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	devices, err := pf.List(owner, models.NewQuery())
	if err != nil {
		return nil, grpcError(err)
	}
	resp := &pb.GetDevicesResponse{
		DeviceIDs: []string{},
	}
	for _, device := range devices {
		resp.DeviceIDs = append(resp.DeviceIDs, device.(*models.DeviceInProject).DeviceID)
	}
	return resp, nil
}

// Workshop
// AddWorkshop add a workshop into the project
func (s *ProjectManagementService) AddWorkshop(ctx context.Context, in *pb.AddWorkshopRequest) (*pb.AddWorkshopResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)
	if _, err := pf.Save(owner, converter.NewWorkshopModel(in.Workshop)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.AddWorkshopResponse{}, nil

}

// DeleteWorkshop remove a workshop from project
func (s *ProjectManagementService) DeleteWorkshop(ctx context.Context, in *pb.DeleteWorkshopRequest) (*pb.DeleteWorkshopResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	if err := pf.Delete(owner, in.WorkshopID); err != nil {
		return nil, grpcError(err)
	}

	return &pb.DeleteWorkshopResponse{}, nil
}

// GetWorkshops return a project's all workshops
func (s *ProjectManagementService) GetWorkshops(ctx context.Context, in *pb.GetWorkshopsRequest) (*pb.GetWorkshopsResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	workshopModels, err := pf.List(owner, models.NewQuery())
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetWorkshopsResponse{Workshops: converter.NewWorkshops(workshopModels)}, nil
}

// GetWorkshop return specified workshop
func (s *ProjectManagementService) GetWorkshop(ctx context.Context, in *pb.GetWorkshopRequest) (*pb.GetWorkshopResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	workshopModel, err := pf.Get(owner, in.WorkshopID)
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetWorkshopResponse{Workshop: converter.NewWorkshop(workshopModel)}, nil
}

// UpdateWorkshop update specified workshop
func (s *ProjectManagementService) UpdateWorkshop(ctx context.Context, in *pb.UpdateWorkshopRequest) (*pb.UpdateWorkshopResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	if err := pf.Update(owner, converter.NewWorkshopModel(in.Workshop)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.UpdateWorkshopResponse{}, nil

}

// CreateView create a new project's view
func (s *ProjectManagementService) CreateView(ctx context.Context, in *pb.CreateViewRequest) (*pb.CreateViewResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	if _, err := pf.Save(owner, converter.NewViewModel(in.View)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.CreateViewResponse{}, nil

}

// DeleteView delete a project's view
func (s *ProjectManagementService) DeleteView(ctx context.Context, in *pb.DeleteViewRequest) (*pb.DeleteViewResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)

	if err := pf.Delete(owner, in.ViewID); err != nil {
		return nil, grpcError(err)
	}

	return &pb.DeleteViewResponse{}, nil
}

// GetViews return a project's all views
func (s *ProjectManagementService) GetViews(ctx context.Context, in *pb.GetViewsRequest) (*pb.GetViewsResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)

	viewModels, err := pf.List(owner, models.NewQuery())
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetViewsResponse{Views: converter.NewViews(viewModels)}, nil
}

// GetView return a view's detail informaiton
func (s *ProjectManagementService) GetView(ctx context.Context, in *pb.GetViewRequest) (*pb.GetViewResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)

	viewModel, err := pf.Get(owner, in.ViewID)
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetViewResponse{View: converter.NewView(viewModel)}, nil
}

// UpdateView update a specified view
func (s *ProjectManagementService) UpdateView(ctx context.Context, in *pb.UpdateViewRequest) (*pb.UpdateViewResponse, error) {
	pf := factory.NewFactory(models.Workshop{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	if err := pf.Update(owner, converter.NewViewModel(in.View)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.UpdateViewResponse{}, nil
}

// Variables
// CreateVariable create a new variable in view or project
func (s *ProjectManagementService) CreateVariable(ctx context.Context, in *pb.CreateVariableRequest) (*pb.CreateVariableResponse, error) {
	pf := factory.NewFactory(models.Variable{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	if _, err := pf.Save(owner, converter.NewVariableModel(in.Variable)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.CreateVariableResponse{}, nil
}

// GetVariable return a variable's detail information
func (s *ProjectManagementService) GetVariable(ctx context.Context, in *pb.GetVariableRequest) (*pb.GetVariableResponse, error) {
	pf := factory.NewFactory(models.Variable{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)

	variableModel, err := pf.Get(owner, in.VariableID)
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetVariableResponse{Variable: converter.NewVariable(variableModel)}, nil
}

// GetVariables return all variables in a view or project
func (s *ProjectManagementService) GetVariables(ctx context.Context, in *pb.GetVariablesRequest) (*pb.GetVariablesResponse, error) {
	pf := factory.NewFactory(models.Variable{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)
	query := models.NewQuery()

	variableModels, err := pf.List(owner, query)
	if err != nil {
		return nil, grpcError(err)
	}
	return &pb.GetVariablesResponse{Variables: converter.NewVariables(variableModels)}, nil
}

// DeleteVariable delete a variable in view or project
func (s *ProjectManagementService) DeleteVariable(ctx context.Context, in *pb.DeleteVariableRequest) (*pb.DeleteVariableResponse, error) {
	pf := factory.NewFactory(models.View{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)

	if err := pf.Delete(owner, in.VariableID); err != nil {
		return nil, grpcError(err)
	}

	return &pb.DeleteVariableResponse{}, nil
}

// DeleteVariables delete a batch of variables
func (s *ProjectManagementService) DeleteVariables(ctx context.Context, in *pb.DeleteVariablesRequest) (*pb.DeleteVariablesResponse, error) {
	pf := factory.NewFactory(models.View{})
	owner := factory.NewOwner(in.UserID).
		WithProject(in.ProjectID).
		WithWorkshop(in.WorkshopID)

	for _, variableID := range in.VariableIDs {
		if err := pf.Delete(owner, variableID); err != nil {
			return nil, grpcError(err)
		}
	}
	return &pb.DeleteVariablesResponse{}, nil
}

// UpdateVariable update a specified variable in view or project
func (s *ProjectManagementService) UpdateVariable(ctx context.Context, in *pb.UpdateVariableRequest) (*pb.UpdateVariableResponse, error) {
	pf := factory.NewFactory(models.View{})
	owner := factory.NewOwner(in.UserID).WithProject(in.ProjectID)

	if err := pf.Update(owner, converter.NewVariableModel(in.Variable)); err != nil {
		return nil, grpcError(err)
	}
	return &pb.UpdateVariableResponse{}, nil
}
