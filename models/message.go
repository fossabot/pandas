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
	"github.com/sirupsen/logrus"
)

// Message ...
type Message interface {
	GetID() string
	GetOriginator() string
	GetType() string
	GetPayload() []byte
	GetMetadata() Metadata
	SetType(string)
	SetPayload([]byte)
	SetMetadata(Metadata)
	SetOriginator(string)
	Validate(formats strfmt.Registry) error
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
}

// Metadata ...
type Metadata interface {
	Keys() []string
	GetKeyValue(key string) interface{}
	SetKeyValue(key string, val interface{})
}

const (
	OBJECT_PATH_MESSAGES = "/pandas/messages"
)

// Predefined metadata key
const (
	MetadataDeviceType = "deviceType"
	MetadataDeviceName = "deviceName"
	MetadataRequestID  = "requestId"
	MetadataUserName   = "userName"
	MetadataUserID     = "userId"
	MetadataTimestamp  = "timestamp"
)

// Predefined message types
const (
	MessageTypePostAttributesRequest = "Post attributes"
	MessageTypePostTelemetryRequest  = "Post telemetry"
	MessageTypeActivityEvent         = "Activity event"
	MessageTypeInactivityEvent       = "Inactivity event"
	MessageTypeConnectEvent          = "Connect event"
	MessageTypeDisconnectEvent       = "Disconnect event"
)

// NewMessage ...
func NewMessage() Message {
	return &defaultMessage{
		payload: []byte{},
	}
}

type defaultMessage struct {
	id          string
	originator  string
	messageType string
	payload     []byte
	metadata    Metadata
}

// NewMessageWithDetail ...
func NewMessageWithDetail(id string, originator string, messageType string, payload []byte, metadata Metadata) Message {
	return &defaultMessage{
		id:          id,
		originator:  originator,
		messageType: messageType,
		payload:     payload,
		metadata:    metadata,
	}
}

func (t *defaultMessage) GetID() string                   { return t.id }
func (t *defaultMessage) GetOriginator() string           { return t.originator }
func (t *defaultMessage) GetType() string                 { return t.messageType }
func (t *defaultMessage) GetPayload() []byte              { return t.payload }
func (t *defaultMessage) GetMetadata() Metadata           { return t.metadata }
func (t *defaultMessage) SetType(messageType string)      { t.messageType = messageType }
func (t *defaultMessage) SetPayload(payload []byte)       { t.payload = payload }
func (t *defaultMessage) SetMetadata(metadata Metadata)   { t.metadata = metadata }
func (t *defaultMessage) SetOriginator(originator string) { t.originator = originator }

func (t *defaultMessage) Validate(formats strfmt.Registry) error { return nil }
func (t *defaultMessage) MarshalBinary() ([]byte, error)         { return nil, nil }
func (t *defaultMessage) UnmarshalBinary(b []byte) error         { return nil }

// NewMetadata ...
func NewMetadata() Metadata {
	return &defaultMetadata{
		values: make(map[string]interface{}),
	}
}

type defaultMetadata struct {
	values map[string]interface{}
}

func newdefaultMetadata(vals map[string]interface{}) Metadata {
	return &defaultMetadata{
		values: vals,
	}
}

func (t *defaultMetadata) Keys() []string {
	keys := []string{}
	for key := range t.values {
		keys = append(keys, key)
	}
	return keys
}

func (t *defaultMetadata) GetKeyValue(key string) interface{} {
	if _, found := t.values[key]; !found {
		logrus.Fatalf("no key '%s' in metadata", key)
	}
	return t.values[key]
}

func (t *defaultMetadata) SetKeyValue(key string, val interface{}) {
	t.values[key] = val
}
