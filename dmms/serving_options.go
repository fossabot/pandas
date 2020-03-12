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
package dmms

import (
	"fmt"

	modeloptions "github.com/cloustone/pandas/models/options"
	"github.com/spf13/pflag"
)

// ServingOptions
type ServingOptions struct {
	ServingOptions  modeloptions.ServingOptions
	DeviceModelPath string
}

func NewServingOptions() *ServingOptions {
	return &ServingOptions{
		//DeviceModelPath: "./pandas/models",
		DeviceModelPath: "./models",
	}
}

func (s *ServingOptions) Validate() []error {
	errors := []error{}

	if s.DeviceModelPath == "" {
		errors = append(errors, fmt.Errorf("--device-models-path is not specified"))
	}
	return errors
}

func (s *ServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.DeviceModelPath, "device-models-path", s.DeviceModelPath, "The device models path")
}
