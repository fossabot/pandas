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
package cache

import (
	"time"

	modelsoptions "github.com/cloustone/pandas/models/options"
)

// Cache is simple wrapper of redis or memcahed client
type Cache interface {
	Set(key string, val interface{}) (interface{}, error)
	Get(key string) (interface{}, error)
	Delete(key string) error
	Setnx(key string, val interface{}) (interface{}, error)
	Expire(key string, duration time.Duration) error
	ListPush(key string, args ...interface{}) (interface{}, error)
	ListRange(key string, start string, size string) ([]interface{}, error)
	Flush()
}

// NewCache return cache client endpoint
func NewCache(options *modelsoptions.ServingOptions) Cache {
	return newRedisCache(options)
}
