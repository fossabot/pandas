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

	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
)

type testingMessage struct {
	id          string
	originator  string
	messageType string
	payload     []byte
	metadata    Metadata
}

func NewTestingMessage(id string, originator string, messageType string, payload []byte, metadata Metadata) Message {
	return &testingMessage{
		id:          id,
		originator:  originator,
		messageType: messageType,
		payload:     payload,
		metadata:    metadata,
	}
}

func (t *testingMessage) GetID() string                   { return t.id }
func (t *testingMessage) GetOriginator() string           { return t.originator }
func (t *testingMessage) GetType() string                 { return t.messageType }
func (t *testingMessage) GetPayload() []byte              { return t.payload }
func (t *testingMessage) GetMetadata() Metadata           { return t.metadata }
func (t *testingMessage) SetType(messageType string)      { t.messageType = messageType }
func (t *testingMessage) SetPayload(payload []byte)       { t.payload = payload }
func (t *testingMessage) SetMetadata(metadata Metadata)   { t.metadata = metadata }
func (t *testingMessage) SetOriginator(originator string) { t.originator = originator }

func (t *testingMessage) Validate(formats strfmt.Registry) error { return nil }
func (t *testingMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
func (t *testingMessage) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, t)
}

type testingMetadata struct {
	values map[string]interface{}
}

func newTestingMetadata(vals map[string]interface{}) Metadata {
	return &testingMetadata{
		values: vals,
	}
}

func (t *testingMetadata) Keys() []string {
	keys := []string{}
	for key := range t.values {
		keys = append(keys, key)
	}
	return keys
}

func (t *testingMetadata) GetKeyValue(key string) interface{} {
	if _, found := t.values[key]; !found {
		logrus.Fatalf("no key '%s' in metadata", key)
	}
	return t.values[key]
}

func (t *testingMetadata) SetKeyValue(key string, val interface{}) {
	t.values[key] = val
}
