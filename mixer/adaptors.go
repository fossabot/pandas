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
	"github.com/sirupsen/logrus"
)

// allAdaptorFactories hold all embeded data source factory
var allAdaptorFactories map[string]adaptors.AdaptorFactory = make(map[string]adaptors.AdaptorFactory)

// RegisterAdaptor register data source in datas or plugins
func RegisterAdaptor(name string, factory adaptors.AdaptorFactory) {
	allAdaptorFactories[name] = factory
}

// GetAllAdaptors returan all data source names
func GetAllAdaptors() []string {
	names := []string{}
	for key, _ := range allAdaptorFactories {
		names = append(names, key)
	}
	return names
}

// DumpAllAdaptors dump all data sources
func DumpAllAdaptors() {
	for name, _ := range allAdaptorFactories {
		logrus.Infof("data source '%s' is registered", name)
	}
}

// NewAdaptor create a new data source with source configuration
func NewAdaptor(name string, servingOptions *adaptors.AdaptorOptions) (adaptors.Adaptor, error) {
	for name, factory := range allAdaptorFactories {
		if name == name && factory != nil {
			return factory.Create(servingOptions)
		}
	}
	return nil, fmt.Errorf("unknown reader '%s'", name)
}

func init() {
	RegisterAdaptor("grpc", grpc.AdaptorFactory{})
}
