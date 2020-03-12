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
	"errors"
	"fmt"
	"reflect"

	"github.com/cloustone/pandas/models"
	modeloptions "github.com/cloustone/pandas/models/options"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var (
	ErrObjectNotFound       = errors.New("object not found")
	ErrObjectAlreadyExist   = errors.New("object already exist")
	ErrObjectInvalidArg     = errors.New("invalid object args")
	ErrFactoryInternalError = errors.New("object factory internal")
)

var (
	factories map[string]Factory = make(map[string]Factory)
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
}

// NewFactory create and return model factory
func NewFactory(obj interface{}) Factory {
	name := obj.(string)
	//name := reflect.TypeOf(obj).Name()
	//fmt.Println(obj)
	//fmt.Println(reflect.TypeOf(obj))
	if factory, found := factories[name]; found {
		return factory
	}
	logrus.Fatalf("unregistered model '%s' factory", name)
	return nil
}

// RegisterFactory will add model factory. user can also add customized model factory
func RegisterFactory(model interface{}, f Factory) {
	factories[reflect.TypeOf(model).Name()] = f
	fmt.Println(reflect.TypeOf(model).Name())
}

// Initialize will be called in startup to initialize all internal model factory
func Initialize(servingOptions *modeloptions.ServingOptions) {
	// Open pandas data based
	db, err := gorm.Open(servingOptions.StorePath, "pandas.db")
	fmt.Println(servingOptions.StorePath)
	fmt.Println(err)
	if err != nil {
		logrus.Fatalf("factory failed to open database")
	}
	defer db.Close()

}

func Error(db *gorm.DB) error {
	if errs := db.GetErrors(); len(errs) >= 0 {
		switch errs[0] {
		case gorm.ErrRecordNotFound:
			return ErrObjectNotFound
		case gorm.ErrInvalidSQL:
			return ErrObjectInvalidArg
		default:
			return ErrFactoryInternalError
		}
	}
	return nil
}

func NewCacheID(owner Owner, additionals ...string) string {
	id := owner.User()
	if owner.Project() != "" {
		id += "_" + owner.Project()
	}
	for _, additionalID := range additionals {
		id += "_" + additionalID
	}
	return id
}
