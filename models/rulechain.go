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

// RuleStatus ...
const (
	RULE_STATUS_CREATED = "created"
	RULE_STATUS_STARTED = "started"
	RULE_STATUS_STOPPED = "stopped"
	RULE_STATUS_UNKNOWN = "unknown"
)

const (
	OBJECT_PATH_RULECHAIN = "/pandas/rulechains"
)

type DataSource struct {
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	IsProvider   bool   `json:"isProvider"`
	ServicePort  string `json:"servicePort"`
	IsTLSEnabled bool   `json:"isTlsEnabled"`
	ConnectURL   string `json:"connectURL"`
	CertFile     []byte `json:"certFile"`
	KeyFile      []byte `json:"keyFile"`
}

// RuleChain RuleChain
// swagger:model RuleChain
type RuleChain struct {
	ModelTypeInfo
	gorm.Model
	Name          string     `json:"name" gorm:"size:255"`
	ID            string     `json:"chainId" gorm:"size:100,unique_index"`
	Description   string     `json:"description"gorm:"size:255"`
	DebugMode     bool       `json:"debugMode" `
	UserID        string     `json:"userId" gorm:"size:255"`
	Type          string     `json:"type" gorm:"type:char(100)"`
	Domain        string     `json:"domain"`
	Status        string     `json:"status""`
	Payload       []byte     `json:"payload"`
	Root          bool       `json:"bool"`
	CreatedAt     time.Time  `json:"createdAt"`
	LastUpdatedAt time.Time  `json:"lastUpdatedAt"`
	DataSource    DataSource `json:"dataSource"`
}

// Validate validates this deployment
func (m *RuleChain) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleChain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleChain) UnmarshalBinary(b []byte) error {
	var res RuleChain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
