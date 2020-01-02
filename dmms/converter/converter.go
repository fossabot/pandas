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
)

// grpc_dmms_v1.Device and models.Device

func NewDevice(deviceModel *models.Device) *grpc_dmms_v1.Device {
	return nil
}

func NewDevices(deviceModels []*models.Device) []*grpc_dmms_v1.Device {
	return nil
}

func NewDeviceModel(device *grpc_dmms_v1.Device) *models.Device {
	return nil
}

func NewDeviceModels(devices []*grpc_dmms_v1.Device) []models.Device {
	return nil
}

// grpc_dmms_v1.DeviceModel and models.DeviceModel
func NewDeviceWithModel(deviceModel *models.DeviceModel) *grpc_dmms_v1.DeviceModel {
	return nil
}

func NewDevicesWithModel(deviceModels []*models.DeviceModel) []*grpc_dmms_v1.DeviceModel {
	return nil
}

func NewDeviceModelWithModel(device *grpc_dmms_v1.DeviceModel) *models.DeviceModel {
	return nil
}

func NewDeviceModels(devices []*grpc_dmms_v1.DeviceModel) []models.DeviceModel {
	return nil
}
