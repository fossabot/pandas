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

type Direction int

const (
	DirectionUp   Direction = 0
	DirectionDown Direction = 1
	DirectionAll  Direction = 2
)

// swagger:model LogMetaInfo
type LogMetaInfo struct {
	ProjectId  string     `json:"projectId"`
	DeviceId   string     `json:"deviceId"`
	DeviceName string     `json:"deviceName"`
	MessageId  string     `json:"messageId"`
	Type       string     `json:"type"`
	Direction  int        `json:"direction"`
	Time       *time.Time `json:"time"`
	RIndex     int64      `json:"rindex"`
}

// Validate validates this deployment
func (m *LogMetaInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogMetaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogMetaInfo) UnmarshalBinary(b []byte) error {
	var res LogMetaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// swagger:model LogData
type LogData struct {
	Meta    LogMetaInfo `json:"meta"`
	Line    string      `json:"line"`
	Payload []byte      `json:"payload"`
	Result  string      `json:"result"`
}

// Validate validates this deployment
func (m *LogData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogData) UnmarshalBinary(b []byte) error {
	var res LogData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
