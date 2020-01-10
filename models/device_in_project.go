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

// DeviceInProject represent a device in project
type DeviceInProject struct {
	ModelTypeInfo
	DeviceID      string    `json:"deviceID"`
	DeviceName    string    `json:"deviceName"`
	ProjectID     string    `json:"projectID"`
	UserID        string    `json:"userDd"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	Status        string    `json:"status"`
}

// NewDeviceInProject ...
func NewDeviceInProject() *DeviceInProject {
	return &DeviceInProject{}
}

// Validate validates this deployment
func (m *DeviceInProject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInProject) UnmarshalBinary(b []byte) error {
	var res DeviceInProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
