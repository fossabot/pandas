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
package nodes

import (
	"fmt"

	"github.com/goinggo/mapstructure"
)

const (
	NODE_CONFIG_MESSAGE_TYPE_KEY    = "messageTypeKey"
	NODE_CONFIG_ORIGINATOR_TYPE_KEY = "originatorTypeKey"
)

type Metadata interface {
	Keys() []string
	With(key string, val interface{}) Metadata
	Value(key string) (interface{}, error)
	DecodePath(rawVal interface{}) error
}

type nodeMetadata struct {
	keypairs map[string]interface{}
}

func NewMetadata() Metadata {
	return &nodeMetadata{
		keypairs: make(map[string]interface{}),
	}
}

func NewMetadataWithString(vals string) Metadata {
	return &nodeMetadata{}
}

func NewMetadataWithValues(vals map[string]interface{}) Metadata {
	return &nodeMetadata{
		keypairs: vals,
	}
}

func (c *nodeMetadata) Keys() []string {
	keys := []string{}
	for key, _ := range c.keypairs {
		keys = append(keys, key)
	}
	return keys
}

func (c *nodeMetadata) Value(key string) (interface{}, error) {
	if val, found := c.keypairs[key]; found {
		return val, nil
	}
	return nil, fmt.Errorf("key '%s' not found", key)
}

func (c *nodeMetadata) With(key string, val interface{}) Metadata {
	c.keypairs[key] = val
	return c
}

func (c *nodeMetadata) DecodePath(rawVal interface{}) error {
	return mapstructure.DecodePath(c.keypairs, rawVal)
}
