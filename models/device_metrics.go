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

package models

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceMetrics DeviceMetrics
type DeviceMetrics struct {
	DeviceID              string    `json:"deviceID"`
	CreatedAt             time.Time `json:"createdAt"`
	LastUpdatedAt         time.Time `json:"lastUpdatedAt"`
	ConnectCount          int32     `json:"connectCount"`
	DisconnectCount       int32     `json:"disconnectCount"`
	LastConnectedAt       time.Time `json:"lastConnectedAt"`
	LastDisconnectedAt    time.Time `json:"lastDisconnectedAt"`
	MessageCount          int32     `json:"messageCount"`
	LastMessageReceivedAt time.Time `json:"lastMessageReceivedAt"`
}

// NewDeviceMetrics ...
func NewDeviceMetrics() *DeviceMetrics {
	return &DeviceMetrics{
		CreatedAt: time.Now(),
	}
}

// Validate validates this deployment
func (m *DeviceMetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceMetrics) UnmarshalBinary(b []byte) error {
	var res DeviceMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
