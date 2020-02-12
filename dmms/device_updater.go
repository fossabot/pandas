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
	"encoding/json"
	"errors"
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/factory"
	"github.com/sirupsen/logrus"
)

type DeviceUpdater struct{}

func NewDeviceUpdater() *DeviceUpdater {
	return &DeviceUpdater{}
}

// UpdateDeviceValues will update device real values using message received and device model
func (*DeviceUpdater) UpdateDeviceValues(n *DeviceNotification) {
	// Unmarshal device message and match with device model
	msg := models.DeviceMessage{}
	if err := msg.UnmarshalBinary(n.Payload); err != nil {
		logrus.WithError(err)
		return
	}
	// Get device model and device object to updated values
	pf := factory.NewFactory(models.Device{})
	owner := factory.NewOwner(n.UserID)
	obj, err := pf.Get(owner, n.DeviceID)
	if err != nil {
		logrus.Errorf("failed to get device '%s'", n.DeviceID)
		return
	}
	device := obj.(*models.Device)

	// Get device model
	pf = factory.NewFactory(models.DeviceModel{})
	obj, err = pf.Get(owner, device.ModelID)
	if err != nil {
		logrus.Errorf("failed to get device '%s' with model '%s'", n.DeviceID, device.ModelID)
		return
	}
	deviceModel := obj.(*models.DeviceModel)
	dataModelName := ""

	// Use endpoint to get in device data model name
	for _, endpoint := range deviceModel.Endpoints {
		if endpoint.Path == n.Endpoint {
			dataModelName = endpoint.Models[models.KEndpointDirectionIn]
			break // found
		}
	}
	// Not found
	if dataModelName == "" {
		logrus.Errorf("device '%s' data model '%s' not found", n.DeviceID, device.ModelID)
		return

	}
	// Update device
	for _, deviceDataModel := range device.Values {
		if deviceDataModel.Name == dataModelName { // found
			updateDeviceValueWithMessage(&deviceDataModel, &msg)
			break
		}
	}
	pf.Update(owner, device)
}

// updateDeviceValueWithMessage update data model with incomming device message
func updateDeviceValueWithMessage(dataModel *models.DataModel, msg *models.DeviceMessage) {
	values := make(map[string]interface{})
	if err := json.Unmarshal(msg.Payload, values); err != nil {
		logrus.Errorf("device message '%s' payload error", msg.ID)
		return
	}
	for index, field := range dataModel.Fields {
		if value, found := values[field.Key]; found {
			dataModel.Fields[index].Value = value.(string)
		}
	}
}

// UpdateDeviceStatus update device status
func (u *DeviceUpdater) UpdateDeviceStatus(n *DeviceNotification) {
	// The device should be authenticated to be as valid device
	pf := factory.NewFactory(models.Device{})
	owner := factory.NewOwner(n.UserID)

	deviceModel, err := pf.Get(owner, n.DeviceID)
	if err != nil {
		logrus.Errorf("ilegal device '%s' notification received", n.DeviceID)
		return
	}
	device := deviceModel.(*models.Device)
	switch n.Type {
	case KDeviceConnected:
		device.Status = models.KDeviceStatusConnected
		break
	case KDeviceDisconnected:
		device.Status = models.KDeviceStatusDisconnected
		break
	default:
		device.Status = models.KDeviceStatusUnknown
	}
	device.LastUpdatedAt = time.Now()
	pf.Update(owner, device)
}

// UpdateDeviceMetrics update device metrics
func (u *DeviceUpdater) UpdateDeviceMetrics(n *DeviceNotification) {
	var deviceMetrics *models.DeviceMetrics

	pf := factory.NewFactory(models.DeviceMetrics{})
	owner := factory.NewOwner(n.UserID)
	deviceMetricsModel, err := pf.Get(owner, n.DeviceID)
	if err != nil {
		if errors.Is(err, factory.ErrObjectNotFound) { // the device metrics not exist
			deviceMetrics = &models.DeviceMetrics{
				DeviceID:      n.DeviceID,
				CreatedAt:     time.Now(),
				LastUpdatedAt: time.Now(),
			}
		}
	} else {
		deviceMetrics = deviceMetricsModel.(*models.DeviceMetrics)
	}
	switch n.Type {
	case KDeviceConnected:
		deviceMetrics.LastUpdatedAt = time.Now()
		deviceMetrics.ConnectCount += 1
		deviceMetrics.LastConnectedAt = time.Now()
		break

	case KDeviceDisconnected:
		deviceMetrics.LastUpdatedAt = time.Now()
		deviceMetrics.DisconnectCount += 1
		deviceMetrics.LastDisconnectedAt = time.Now()
		break

	case KDeviceMessageReceived:
		deviceMetrics.LastUpdatedAt = time.Now()
		deviceMetrics.MessageCount += 1
		deviceMetrics.LastMessageReceivedAt = time.Now()
	}
	pf.Save(owner, deviceMetrics)
}
