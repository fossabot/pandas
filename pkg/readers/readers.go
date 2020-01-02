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
package readers

import (
	"fmt"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/readers/grpc"
	"github.com/sirupsen/logrus"
)

type ReaderFactory interface {
	Create(map[string]interface{}) (models.Reader, error)
}

// allReaderFactories hold all embeded data source factory
var allReaderFactories map[string]ReaderFactory = make(map[string]ReaderFactory)

// RegisterReader register data source in datas or plugins
func RegisterReader(name string, factory ReaderFactory) {
	allReaderFactories[name] = factory
}

// GetAllReaders returan all data source names
func GetAllReaders() []string {
	names := []string{}
	for key, _ := range allReaderFactories {
		names = append(names, key)
	}
	return names
}

// DumpAllReaders dump all data sources
func DumpAllReaders() {
	for name, _ := range allReaderFactories {
		logrus.Infof("data source '%s' is registered", name)
	}
}

// NewReader create a new data source with source configuration
func NewReader(name string, c map[string]string) (models.Reader, error) {
	// convert config models
	configs := make(map[string]interface{})
	for k, v := range c {
		configs[k] = v
	}

	for name, factory := range allReaderFactories {
		if name == name && factory != nil {
			return factory.Create(configs)
		}
	}
	return nil, fmt.Errorf("unknown reader '%s'", name)
}

func init() {
	RegisterReader("grpc", grpc.ReaderFactory{})
}
