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
	"errors"
	"time"

	"github.com/cloustone/pandas/models"
	modelsoptions "github.com/cloustone/pandas/models/options"
)

// noneCache is just a simple helper to making source code unified when no
// cache backend is specified
type noneCache struct{}

// newNoneCache return cache backend instance
func newNoneCache(options *modelsoptions.ServingOptions) Cache { return &noneCache{} }

// Set just return nil to tell client cache operation is successful
func (r *noneCache) Set(key string, val models.Model) error { return nil }

// Get just return errors to tell client to use database backend
func (r *noneCache) Get(key string, val models.Model) error {
	return errors.New("no item exist")
}

// Delete just return nil to to tell client tha cache deleting is successful
func (r *noneCache) Delete(key string) error { return nil }

// Expire just return ok
func (r *noneCache) Expire(key string, duration time.Duration) error { return nil }

// ListPush just return ok
func (r *noneCache) ListPush(key string, args ...models.Model) error {
	return nil
}

// ListRange just return errors that no items exist
func (r *noneCache) ListRange(key string, start string, size string, model models.Model) ([]models.Model, error) {
	return nil, errors.New("no items exist")
}

func (r *noneCache) Flush() {}
