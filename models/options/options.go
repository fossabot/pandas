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
package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

const (
	KCacheNone  = "none"
	KCacheRedis = "redis"
	KCacheLocal = "local"
)

// ServingOptions
type ServingOptions struct {
	// StorePath is backend storage connect url
	StorePath         string
	Cache             string
	CacheConnectedUrl string
}

func NewServingOptions() *ServingOptions {
	return &ServingOptions{
		StorePath:         "sqlite3",
		Cache:             KCacheNone,
		CacheConnectedUrl: "127.0.0.1:6379",
	}
}

func (s *ServingOptions) Validate() []error {
	errors := []error{}

	if s.StorePath == "" {
		errors = append(errors, fmt.Errorf("--models-store-path is not specified"))
	}

	return errors
}

func (s *ServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.StorePath, "models-store-path", s.StorePath, "The backend storage connect url.")
	fs.StringVar(&s.Cache, "cache", s.Cache, "cache method for models backend, options(none, local, redis.")
	fs.StringVar(&s.CacheConnectedUrl, "cache-connected-url", s.CacheConnectedUrl, "The backend cache connectd url.")
}
