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
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

type redisCache struct {
	conn redis.Conn
}

// newRedisCache return cache instance using redis
func newRedisCache(options *modelsoptions.ServingOptions) Cache {
	conn, err := redis.Dial("tcp", options.CacheConnectedUrl)
	if err != nil {
		logrus.Fatal(err)
	}
	return &redisCache{conn: conn}
}

// Set add a key val pair into cache
func (r *redisCache) Set(key string, val models.Model) error {
	buf, err := val.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = r.conn.Do("SETNX", key, buf)
	return err
}

// Get retrieve val with specified key
func (r *redisCache) Get(key string, model models.Model) error {
	val, err := redis.Bytes(r.conn.Do("GET", key))
	if err != nil {
		return err
	}
	return model.UnmarshalBinary(val)
}

// Delete remove a cache item from cache
func (r *redisCache) Delete(key string) error {
	_, err := r.conn.Do("DEL", key)
	return err
}

// Expire set cache timeout
func (r *redisCache) Expire(key string, duration time.Duration) error {
	_, err := r.conn.Do("EXPIRE", key, duration)
	return err
}

// ListPush push cache item as list
func (r *redisCache) ListPush(key string, args ...models.Model) error {
	values := []interface{}{}
	for _, val := range args {
		v, err := val.MarshalBinary()
		if err != nil {
			return err
		}
		values = append(values, v)
	}
	_, err := r.conn.Do("lpush", values)
	return err
}

// ListRange return specified cache item as list
func (r *redisCache) ListRange(key string, start string, size string, model models.Model) ([]models.Model, error) {
	values, err := redis.Values(r.conn.Do("lrange", key, start, size))
	if err != nil {
		return nil, err
	}
	modelSlice := []models.Model{}
	for _, val := range values {
		model, err := models.UnmarshalBinary(val.([]byte))
		if err != nil {
			return nil, err
		}
		modelSlice = append(modelSlice, model)
	}
	return modelSlice, nil
}
