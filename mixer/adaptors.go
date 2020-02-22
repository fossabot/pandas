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
package mixer

import (
	"fmt"

	"github.com/cloustone/pandas/mixer/adaptors"
	"github.com/cloustone/pandas/mixer/adaptors/grpc"
	"github.com/cloustone/pandas/mixer/adaptors/mqtt"
	"github.com/sirupsen/logrus"
)

// allAdaptorFactories hold all embeded adaptors factory
var allAdaptorFactories map[string]adaptors.AdaptorFactory = make(map[string]adaptors.AdaptorFactory)

// RegisterAdaptor register new adaptor facotory
func RegisterAdaptor(name string, factory adaptors.AdaptorFactory) {
	allAdaptorFactories[name] = factory
}

// GetFactories return all adaptor factories name
func GetFactories() []string {
	names := []string{}
	for key, _ := range allAdaptorFactories {
		names = append(names, key)
	}
	return names
}

// DumpAllAdaptors dump all adaptors
func DumpAllAdaptors() {
	for name, _ := range allAdaptorFactories {
		logrus.Infof("data source '%s' is registered", name)
	}
}

// NewAdaptor create a new adaptor with specified options
func NewAdaptor(adaptorOptions *adaptors.AdaptorOptions) (adaptors.Adaptor, error) {
	for name, factory := range allAdaptorFactories {
		if name == adaptorOptions.Name && factory != nil {
			return factory.Create(adaptorOptions)
		}
	}
	return nil, fmt.Errorf("unknown adaptor '%s'", adaptorOptions.Name)
}

func init() {
	RegisterAdaptor("grpc", grpc.AdaptorFactory{})
	RegisterAdaptor("mqtt", mqtt.AdaptorFactory{})
}
