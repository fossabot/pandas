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

type LocationServingOptions struct {
	// Provider is location engine name, baidu or othere
	Provider string

	// AK is access key for lbs service provider
	AK string

	// ServiceId is service id for lbs service provider
	ServiceId string
}

func NewLocationServingOptions() *LocationServingOptions {
	return &LocationServingOptions{
		Provider: "baidu",
	}
}

func (s *LocationServingOptions) Validate() []error {
	errors := []error{}
	if s.Provider == "" {
		errors = append(errors, fmt.Errorf("--lbs-provider must be assigned."))
	}

	return errors
}

func (s *LocationServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Provider, "lbs-provider", s.Provider, ""+
		"Location service provider must be specified")

	fs.StringVar(&s.AK, "lbs-ak", s.AK, ""+
		"Access Key provied by Location service provider must be specified")

	fs.StringVar(&s.ServiceId, "lbs-service-id", s.ServiceId, ""+
		"Service ID provied by ocation service provider must be specified")
}

func (s *LocationServingOptions) AddDeprecatedFlags(fs *pflag.FlagSet) {
}
