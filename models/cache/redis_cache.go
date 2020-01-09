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
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

type redisCache struct {
	conn redis.Conn
}

func newRedisCache(options *modelsoptions.ServingOptions) Cache {
	conn, err := redis.Dial("tcp", options.CacheConnectedUrl)
	if err != nil {
		logrus.Fatal(err)
	}
	return &redisCache{conn: conn}
}

func (r *redisCache) Set(key string, val interface{}) (interface{}, error) {
	return nil, nil
}

func (r *redisCache) Get(key string) (interface{}, error) {
	return nil, nil
}

func (r *redisCache) Delete(key string) error {
	return nil
}

func (r *redisCache) Setnx(key string, val interface{}) (interface{}, error) {
	return nil, nil
}

func (r *redisCache) Expire(key string, duration time.Duration) error {
	return nil
}

func (r *redisCache) ListPush(key string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (r *redisCache) ListRange(key string, start string, size string) ([]interface{}, error) {
	return nil, nil
}

func (r *redisCache) Flush() {}
