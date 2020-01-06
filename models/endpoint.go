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

const (
	KEndpointDirectionIn  = "in"
	KEndpointDirectionOut = "out"
)

// Endpoint Endpoint
// swagger:model Endpoint
type Endpoint struct {
	Path   string            `json:"path"`
	Format string            `json:"format"`
	Models map[string]string `json:"models"`
}

// NewEndpoint ...
func NewEndpoint() *Endpoint {
	return &Endpoint{
		Models: make(map[string]string),
		Format: "application/json",
	}
}

// WithPath ...
func (m *Endpoint) WithPath(path string) *Endpoint {
	m.Path = path
	return m
}

// WithDataModel ...
func (m *Endpoint) WithDataModel(modelName string, direction string) *Endpoint {
	m.Models[modelName] = direction
	return m
}

// WithFormat ...
func (m *Endpoint) WithFormat(format string) *Endpoint {
	m.Format = format
	return m
}

// Validate validates this deployment
func (m *Endpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Endpoint) UnmarshalBinary(b []byte) error {
	var res Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
