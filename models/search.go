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

type SearchType int

const (
	SearchTypeTerm   SearchType = 0
	SearchTypeNested SearchType = 1
	SearchTypeAll    SearchType = 2
)

// swagger:model Search
type Search struct {
	Type     SearchType  `json:"type"`
	Meta     LogMetaInfo `json:"meta"`
	Since    string      `json:"sinice"`
	PageSize int32       `json:"pageSize"`
	Pos      int32       `json:"pos"`
}

// Validate validates this deployment
func (m *Search) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Search) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Search) UnmarshalBinary(b []byte) error {
	var res Search
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
