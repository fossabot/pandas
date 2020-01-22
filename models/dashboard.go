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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Dashboard Dashboard
// swagger:model Dashboard
type Dashboard struct {
	ID                      string `json:"id"`
	ProjectNumber           string `json:"projectNumber"`
	ProjectActiveNumber     string `json:"projectActivieNumber"`
	DeviceModelNumber       string `json:"modelNumber"`
	DeviceModelActiveNumber string `json:"deviceModelActiveNumber"`
	DeviceNumber            string `json:"deviceNumber"`
	DeviceActiveNumber      string `json:"deviceActiveNumber"`
	WorkshopNumber          string `json:"workshopNumber"`
	WorkshopActiveNumber    string `json:"workshopActiveNumber"`
}

// Validate validates this deployment
func (m *Dashboard) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Dashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dashboard) UnmarshalBinary(b []byte) error {
	var res Dashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
