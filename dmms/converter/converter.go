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
	"github.com/cloustone/pandas/dmms/grpc_dmms_v1"
	"github.com/cloustone/pandas/models"
	"github.com/golang/protobuf/ptypes"
)

// grpc_dmms_v1.Device and models.Device

func NewDevice(deviceModel *models.Device) *grpc_dmms_v1.Device {
	createdAt, _ := ptypes.TimestampProto(deviceModel.CreatedAt)
	lastUpdatedAt, _ := ptypes.TimestampProto(deviceModel.LastUpdatedAt)
	dataModel := grpc_dmms_v1.DataModel{
		DataModelFields: []*grpc_dmms_v1.DataModelField{},
	}

	for _, field := range deviceModel.DataModel.Fields {
		dataModel.DataModelFields = append(dataModel.DataModelFields,
			&grpc_dmms_v1.DataModelField{
				Key:          field.Key,
				Value:        field.Value,
				Type:         field.Type,
				DefaultValue: field.DefaultValue,
			})
	}

	return &grpc_dmms_v1.Device{
		ID:            deviceModel.ID,
		Name:          deviceModel.Name,
		Description:   deviceModel.Description,
		Status:        deviceModel.Status,
		UserID:        deviceModel.UserID,
		ProjectID:     deviceModel.ProjectID,
		ModelID:       deviceModel.ModelID,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		DataModel:     &dataModel,
	}
}

func NewDevices(deviceModels []*models.Device) []*grpc_dmms_v1.Device {
	devices := []*grpc_dmms_v1.Device{}
	for _, deviceModel := range deviceModels {
		devices = append(devices, NewDevice(deviceModel))
	}
	return devices
}

func NewDeviceModel(device *grpc_dmms_v1.Device) *models.Device {
	createdAt, _ := ptypes.Timestamp(device.CreatedAt)
	lastUpdatedAt, _ := ptypes.Timestamp(device.LastUpdatedAt)

	// DataModels
	dataModel := models.DataModel{
		Fields: []*models.DataModelField{},
	}
	for _, field := range device.DataModel.DataModelFields {
		dataModel.Fields = append(dataModel.Fields,
			&models.DataModelField{
				Key:          field.Key,
				Value:        field.Value,
				Type:         field.Type,
				DefaultValue: field.DefaultValue,
			})
	}

	return &models.Device{
		ID:            device.ID,
		Name:          device.Name,
		Description:   device.Description,
		Status:        device.Status,
		UserID:        device.UserID,
		ProjectID:     device.ProjectID,
		ModelID:       device.ModelID,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		DataModel:     dataModel,
	}
}

func NewDeviceModels(devices []*grpc_dmms_v1.Device) []models.Device {
	deviceModels := []models.Device{}
	for _, device := range devices {
		deviceModels = append(deviceModels, *NewDeviceModel(device))
	}
	return deviceModels
}

// grpc_dmms_v1.DeviceModel and models.DeviceModel
func NewDeviceModel2(model2 *models.DeviceModel) *grpc_dmms_v1.DeviceModel {
	createdAt, _ := ptypes.TimestampProto(model2.CreatedAt)
	lastUpdatedAt, _ := ptypes.TimestampProto(model2.LastUpdatedAt)

	// DataModels
	dataModels := []*grpc_dmms_v1.DataModel{}
	for _, dataModel := range model2.DataModels {
		fields := []*grpc_dmms_v1.DataModelField{}
		for _, field := range dataModel.Fields {
			fields = append(fields,
				&grpc_dmms_v1.DataModelField{
					Key:          field.Key,
					Value:        field.Value,
					Type:         field.Type,
					DefaultValue: field.DefaultValue,
				})
		}
	}
	// Endpoints
	endpoints := []*grpc_dmms_v1.Endpoint{}
	for _, endpoint := range model2.Endpoints {
		endpoints = append(endpoints, &grpc_dmms_v1.Endpoint{
			Path:   endpoint.Path,
			Format: endpoint.Format,
			Models: endpoint.Models,
		})
	}

	return &grpc_dmms_v1.DeviceModel{
		ID:            model2.ID,
		Name:          model2.Name,
		Description:   model2.Description,
		Domain:        model2.Domain,
		Version:       model2.Version,
		Endpoints:     endpoints,
		DataModels:    dataModels,
		IsLogical:     model2.IsLogical,
		IsCompound:    model2.IsCompound,
		ChildModels:   model2.ChildModels,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		UserID:        model2.UserID,
		Icon:          model2.Icon,
	}
}

func NewDeviceModels2(model2s []*models.DeviceModel) []*grpc_dmms_v1.DeviceModel {
	model2models := []*grpc_dmms_v1.DeviceModel{}
	for _, model := range model2s {
		model2models = append(model2models, NewDeviceModel2(model))
	}
	return model2models
}

func NewDeviceModel2Model(device *grpc_dmms_v1.DeviceModel) *models.DeviceModel {
	createdAt, _ := ptypes.Timestamp(device.CreatedAt)
	lastUpdatedAt, _ := ptypes.Timestamp(device.LastUpdatedAt)

	dataModels := []*models.DataModel{}
	for _, dataModel := range device.DataModels {
		fields := []*models.DataModelField{}
		for _, field := range dataModel.DataModelFields {
			fields = append(fields,
				&models.DataModelField{
					Key:          field.Key,
					Value:        field.Value,
					Type:         field.Type,
					DefaultValue: field.DefaultValue,
				})
		}
	}

	// Endpoints
	endpoints := []*models.Endpoint{}
	for _, endpoint := range device.Endpoints {
		endpoints = append(endpoints, &models.Endpoint{
			Path:   endpoint.Path,
			Format: endpoint.Format,
			Models: endpoint.Models,
		})
	}

	return &models.DeviceModel{
		ID:            device.ID,
		Name:          device.Name,
		Description:   device.Description,
		Domain:        device.Domain,
		Version:       device.Version,
		Endpoints:     endpoints,
		DataModels:    dataModels,
		IsLogical:     device.IsLogical,
		IsCompound:    device.IsCompound,
		ChildModels:   device.ChildModels,
		CreatedAt:     createdAt,
		LastUpdatedAt: lastUpdatedAt,
		UserID:        device.UserID,
		Icon:          device.Icon,
	}

}

func NewDeviceModel2Models(devices []*grpc_dmms_v1.DeviceModel) []models.DeviceModel {
	model2models := []models.DeviceModel{}
	for _, model := range devices {
		model2models = append(model2models, *NewDeviceModel2Model(model))
	}
	return model2models
}
