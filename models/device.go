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
	"github.com/jinzhu/gorm"
)

const (
	KDeviceStatusConnected    = "CONNECTED"
	KDeviceStatusDisconnected = "DISCONNECTED"
	KDeviceStatusUnknown      = "UNKNOWN"
)

// Device Device
// swagger:model Device
type Device struct {
	ModelTypeInfo
	gorm.Model
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Status        string      `json:"status"`
	UserID        string      `json:"userID"`
	ProjectID     string      `json:"projectID"`
	ModelID       string      `json:"modelID"`
	CreatedAt     time.Time   `json:"createdAt"`
	LastUpdatedAt time.Time   `json:"lastUpdatedAt"`
	Values        []DataModel `json:"dataModels"`
}

// Validate validates this deployment
func (m *Device) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
