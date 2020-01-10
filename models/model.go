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
	"encoding/json"
	"reflect"

	"github.com/go-openapi/strfmt"
)

type Model interface {
	Validate(formats strfmt.Registry) error
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
}

// ModelTypeInfo is added to model which should be saved into factory
type ModelTypeInfo struct {
	ModelTypeName string `json:"modelTypeName"`
}

// Unmarshal return real model in val buffer
func UnmarshalBinary(val []byte) (Model, error) {
	modelTypeInfo := ModelTypeInfo{}
	if err := json.Unmarshal(val, &modelTypeInfo); err != nil {
		return nil, err
	}
	switch modelTypeInfo.ModelTypeName {
	case reflect.TypeOf(DeviceModel{}).Name():
		deviceModel := &DeviceModel{}
		err := json.Unmarshal(val, deviceModel)
		return deviceModel, err
	default:
	}
	return nil, nil
}
