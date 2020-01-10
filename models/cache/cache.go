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

	"github.com/cloustone/pandas/models"
	modelsoptions "github.com/cloustone/pandas/models/options"
)

// Cache is simple wrapper of redis or memcahed client
type Cache interface {
	Set(key string, val models.Model) error
	Get(key string, val models.Model) error
	Delete(key string) error
	Expire(key string, duration time.Duration) error
	ListPush(key string, args ...models.Model) error
	ListRange(key string, start string, size string, model models.Model) ([]models.Model, error)
}

// NewCache return cache client endpoint
func NewCache(options *modelsoptions.ServingOptions) Cache {
	switch options.Cache {
	case modelsoptions.KCacheRedis:
		return newRedisCache(options)
	default:
		return newNoneCache(options)
	}
}
