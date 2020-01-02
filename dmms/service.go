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
package dmms

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	pb "github.com/cloustone/pandas/dmms/grpc_dmms_v1"
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/factory"
	"github.com/cloustone/pandas/models/manifest"
	"github.com/sirupsen/logrus"
)

// DeviceManager manage all device and device models which include model definition and
// presentation. Model definition and presentation are wrapped into bundle to
// store into backend storage.
type DeviceManagementService struct{}

func NewDeviceManagementService() *DeviceManagementService {
	return &DeviceManagementService{}
}

// LoadDefaultDeviceModels walk through the specified path and load model
// deinitiontion into manager
func (s *DeviceManagementService) loadPresetDeviceModels(path string) error {
	deviceModels := []*models.DeviceModel{}

	handler := func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		contentType := models.BundleSchemeWithNameSuffix(fi.Name())
		switch contentType {
		case models.BundleSchemeYaml, models.BundleSchemeJSON:
			logrus.Debugf("model definition file '%s' found", filename)

			data, err := ioutil.ReadFile(filename)
			if err != nil {
				logrus.WithError(err).Errorf("read file '%s' failed", filename)
				return err
			}
			deviceModel, err := createDeviceModelInternal(
				models.NewBundle(filename, models.BundleKindDefinition, data, contentType),
			)
			if err != nil {
				logrus.WithError(err)
				return err
			}
			deviceModels = append(deviceModels, deviceModel)
		}
		return nil
	}

	if err := filepath.Walk(path, handler); err != nil {
		logrus.WithError(err).Errorf("failed to load preset device models with path '%s'", path)
		return err
	}

	// These models should be upload to backend database after getting models
	// mgr.repo.SaveDeviceModels(deviceModels)
	return nil
}

func createDeviceModelInternal(b models.Bundle) (*models.DeviceModel, error) {
	model, err := manifest.ParseModelDefinition(b.Content(), string(b.Scheme()))
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}

	deviceModel := models.NewDeviceModel().WithName(model.Name).
		WithDescription(model.Description).
		WithDomain(model.Domain).
		WithVersion(model.Version).
		WithLogical(model.IsLogical).
		WithCompound(model.IsCompound).
		WithIcon(model.Icon)

	// Endpoints
	for _, ep := range model.Endpoints {
		nep := models.NewEndpoint().WithPath(ep.Path).
			WithDataModel(ep.DataModel, ep.Permission)
		deviceModel.AddEndpoint(nep)
	}
	// DataModels
	for _, dataModel := range model.DataModels {
		ndm := models.NewDataModel().WithName(dataModel.Name)
		for _, field := range dataModel.Fields {
			ndm.AddField(models.NewDataModelField(field.Key, field.Type, field.DefaultValue))
		}
	}
	// Child Models
	return deviceModel, nil
}

// CreateDeviceModel create device model with device model bundle,
// After user create device model using web-console, as for user, the
// device model should be created, the creation includ model definition
// creation and model presentation saving
// User can also using the method to create device model with inmemory
// bundle, for this case, the device should also be save to repo
func (s *DeviceManagementService) CreateDeviceModel(ctx context.Context, in *pb.CreateDeviceModelRequest) (*pb.CreateDeviceModelResponse, error) {

	pf := factory.NewFactory(reflect.TypeOf(models.DeviceModel{}).Name())
	owner := factory.NewOwner(in.UserId)
	bundle := models.NewBundle(in.Name, models.BundleKindDefinition, in.Payload, models.BundleSchemeYaml)

	deviceModel, err := createDeviceModelInternal(bundle)
	if err != nil {
		logrus.WithError(err).Errorf("create device model failed")
		return nil, err
	}
	pf.Save(owner, deviceModel)
	return &pb.CreateDeviceModelResponse{}, nil
}

func (s *DeviceManagementService) GetDeviceModel(ctx context.Context, in *pb.GetDeviceModelRequest) (*pb.GetDeviceModelResponse, error) {
	return nil, nil
}
func (s *DeviceManagementService) GetDeviceModelWithName(ctx context.Context, in *pb.GetDeviceModelWithNameRequest) (*pb.GetDeviceModelWithNameResponse, error) {
	return nil, nil
}
func (s *DeviceManagementService) DeleteDeviceModel(ctx context.Context, in *pb.DeleteDeviceModelRequest) (*pb.DeleteDeviceModelResponse, error) {
	return nil, nil
}

// UpdateDeviceModel is called when model presentation is changed using web
// console, the model definition can not be changed without using
// presentation in web console
func (s *DeviceManagementService) UpdateDeviceModel(ctx context.Context, in *pb.UpdateDeviceModelRequest) (*pb.UpdateDeviceModelResponse, error) {
	pf := factory.NewFactory(reflect.TypeOf(models.DeviceModel{}).Name())
	owner := factory.NewOwner(in.UserId)
	bundle := models.NewBundle("", models.BundleKindDefinition, in.Payload, models.BundleSchemeYaml)

	if _, err := pf.Get(owner, in.ModelId); err != nil {
		return nil, models.InvalidModelError.With(in.ModelId)
	}

	deviceModel, err := createDeviceModelInternal(bundle)
	if err != nil {
		logrus.WithError(err).Errorf("model factory update device model failed")
		return nil, err
	}

	pf.Update(owner, deviceModel)
	return &pb.UpdateDeviceModelResponse{}, nil
}

func (s *DeviceManagementService) GetDeviceModels(ctx context.Context, in *pb.GetDeviceModelsRequest) (*pb.GetDeviceModelsResponse, error) {
	return nil, nil
}

func (s *DeviceManagementService) GetDeviceModelBundle(ctx context.Context, in *pb.GetDeviceModelBundleRequest) (*pb.GetDeviceModelBundleResponse, error) {
	return nil, nil
}
