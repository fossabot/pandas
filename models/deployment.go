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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
)

// Deployment status
const (
	DeploymentStatusCreated = "created"
	DeploymentStatusRunning = "running"
	DeploymentStatusStopped = "stopped"
	DeploymentStatusUnknown = "unknown"
)

// Deployment Deployment
// swagger:model Deployment
type Deployment struct {
	ModelTypeInfo
	gorm.Model
	UserID        string            `json:"userId" yaml:"userId" gorm:"size:255"`
	Name          string            `json:"name" yaml:"name" gorm:"size:255"`
	RuleChainID   string            `json:"ruleChainId" yaml:"ruleChainId" gorm:"type:vchar(100)"`
	Reader        string            `json:"reader" yaml:"reader"`
	ReaderConfigs map[string]string `json:"readerConfigs" yaml:"readerConfigs"`
	Status        string            `json:"status" yaml:"status"`
	CreatedAt     *time.Time        `json:"createdAt,omitempty" yaml:"createdAt"`
	ID            string            `json:"id,omitempty" yaml:"id" gorm:"type:vchar(100),unique_index"`
}

// Validate validates this deployment
func (m *Deployment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Deployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deployment) UnmarshalBinary(b []byte) error {
	var res Deployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
