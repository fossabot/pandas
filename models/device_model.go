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
	"github.com/rs/xid"
)

// DeviceModel DeviceModel
// swagger:model DeviceModel
type DeviceModel struct {
	ModelTypeInfo
	gorm.Model
	Name          string       `json:"name"`
	ID            string       `json:"id"`
	Description   string       `json:"description"`
	Domain        string       `json:"domain"`
	Version       string       `json:"version"`
	Endpoints     []*Endpoint  `json:"endpoints"`
	DataModels    []*DataModel `json:"dataModel"`
	IsLogical     bool         `json:"isLogical"`
	IsCompound    bool         `json:"isCompound"`
	ChildModels   []string     `json:"childModels" gorm:"type:string[]"`
	IsPreset      bool         `json:"isPreset"`
	UserID        string       `json:"userId"`
	CreatedAt     time.Time    `json:"createdAt"`
	LastUpdatedAt time.Time    `json:"updatedAt"`
	Icon          string       `json:"icon"`
}

// NewDeviceModel ...
func NewDeviceModel() *DeviceModel {
	return &DeviceModel{
		ID:         xid.New().String(),
		Endpoints:  []*Endpoint{},
		DataModels: []*DataModel{},
		CreatedAt:  time.Now(),
	}
}

// WithName ...
func (m *DeviceModel) WithName(name string) *DeviceModel {
	m.Name = name
	return m
}

// WithVersion ...
func (m *DeviceModel) WithVersion(version string) *DeviceModel {
	m.Version = version
	return m
}

// WithDomain ...
func (m *DeviceModel) WithDomain(domain string) *DeviceModel {
	m.Domain = domain
	return m
}

// WithLogical ...
func (m *DeviceModel) WithLogical(logical bool) *DeviceModel {
	m.IsLogical = logical
	return m
}

// WithCompound ...
func (m *DeviceModel) WithCompound(compound bool) *DeviceModel {
	m.IsCompound = compound
	return m
}

// WithDescription ...
func (m *DeviceModel) WithDescription(description string) *DeviceModel {
	m.Description = description
	return m
}

// WithIcon ...
func (m *DeviceModel) WithIcon(icon string) *DeviceModel {
	m.Icon = icon
	return m
}

// WithChildModels ...
func (m *DeviceModel) WithChildModels(models []string) *DeviceModel {
	m.ChildModels = models
	return m
}

// AddEndpoint ...
func (m *DeviceModel) AddEndpoint(ep *Endpoint) { m.Endpoints = append(m.Endpoints, ep) }

// AddDataModel ...
func (m *DeviceModel) AddDataModel(dataModel *DataModel) {
	m.DataModels = append(m.DataModels, dataModel)
}

// Validate validates this deployment
func (m *DeviceModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceModel) UnmarshalBinary(b []byte) error {
	var res DeviceModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
