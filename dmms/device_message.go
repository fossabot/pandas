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
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/cache"
	"github.com/cloustone/pandas/models/factory"
	modelsoptions "github.com/cloustone/pandas/models/options"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type deviceMessageFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newDeviceMessageFactory(servingOptions *modelsoptions.ServingOptions) factory.Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-devices.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.DeviceMessage{})
	return &deviceMessageFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *deviceMessageFactory) Save(owner factory.Owner, obj models.Model) (models.Model, error) {
	view := obj.(*models.DeviceMessage)
	view.CreatedAt = time.Now()
	pf.modelDB.Save(view)

	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	return view, nil
}

func (pf *deviceMessageFactory) List(owner factory.Owner, query *models.Query) ([]models.Model, error) {
	views := []*models.DeviceMessage{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(views)

	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	results := []models.Model{}
	for _, view := range views {
		results = append(results, view)
	}
	return results, nil
}

func (pf *deviceMessageFactory) Get(factory.Owner, string) (models.Model, error) {
	return nil, nil
}

func (pf *deviceMessageFactory) Delete(factory.Owner, string) error {
	return nil
}

func (pf *deviceMessageFactory) Update(factory.Owner, models.Model) error {
	return nil
}
