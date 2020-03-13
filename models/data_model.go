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

import "github.com/rs/xid"

// DataModelField ...
type DataModelField struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	Type         string `json:"type"`
	DefaultValue string `json:"defaultValue"`
}

// NewDataModelField ...
func NewDataModelField(key string, tp string, defaultValue string) *DataModelField {
	return &DataModelField{
		Key:          key,
		Type:         tp,
		DefaultValue: defaultValue,
	}
}

// DataModel ...
type DataModel struct {
	Name        string            `json:"name"`
	DataModelID string            `json:"datamodelid"`
	Domain      string            `json:"domain"`
	Fields      []*DataModelField `json:"fields"`
}

// NewDataModel ...
func NewDataModel() *DataModel {
	return &DataModel{
		DataModelID: xid.New().String(),
		Fields:      []*DataModelField{},
	}
}

// WithName ...
func (m *DataModel) WithName(name string) *DataModel {
	m.Name = name
	return m
}

// WithDomain ...
func (m *DataModel) WithDomain(domain string) *DataModel {
	m.Domain = domain
	return m
}

// AddField ...
func (m *DataModel) AddField(field *DataModelField) {
	m.Fields = append(m.Fields, field)
}
