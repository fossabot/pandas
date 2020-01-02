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
package factory

import (
	"reflect"

	"github.com/cloustone/pandas/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	factories             map[string]Factory = make(map[string]Factory)
	factoryServingOptions *FactoryServingOptions
)

// Factory will create and manage object
type Factory interface {
	// Save a newly created object into factory
	Save(Owner, models.Model) (models.Model, error)

	// List return object sets accoroding to query
	List(Owner, *models.Query) ([]models.Model, error)

	// Get return a specified object
	Get(Owner, string) (models.Model, error)

	// Delete will delete specified object in factory
	Delete(Owner, string) error

	// Update update object in factory
	Update(Owner, models.Model) error

	// initialize is used internally to initialize factory
	initialize(*FactoryServingOptions) error
}

// NewFactory create and return model factory
func NewFactory(obj interface{}) Factory {
	name := reflect.TypeOf(obj).Name()
	if factory, found := factories[name]; found {
		factory.initialize(factoryServingOptions)
		return factory
	}
	logrus.Fatalf("unregistered model '%s' factory", name)
	return nil
}

// RegisterFactory will add model factory. user can also add customized model factory
func RegisterFactory(model interface{}, f Factory) {
	factories[reflect.TypeOf(model).Name()] = f
}

// Initialize will be called in startup to initialize all internal model factory
func Initialize(factoryServingOptions *FactoryServingOptions) {
	// Open pandas data based
	db, err := gorm.Open(factoryServingOptions.StorePath, "pandas.db")
	if err != nil {
		logrus.Fatalf("factory failed to open database")
	}
	defer db.Close()

	factoryServingOptions = factoryServingOptions

	RegisterFactory(models.Project{}, &projectFactory{})
	RegisterFactory(models.RuleChain{}, &rulechainFactory{})
	RegisterFactory(models.Workshop{}, &workshopFactory{})
	RegisterFactory(models.View{}, &viewFactory{})
	RegisterFactory(models.DeviceInProject{}, &deviceInProjectFactory{})

	for modelName, factory := range factories {
		if err := factory.initialize(factoryServingOptions); err != nil {
			logrus.Fatalf("factory '%s' initialize failed", modelName)
		}
	}
}
