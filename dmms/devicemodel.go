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

type deviceModelFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newDeviceModelFactory(servingOptions *modelsoptions.ServingOptions) factory.Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-dmms.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.DeviceModel{})
	return &deviceModelFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *deviceModelFactory) Save(owner factory.Owner, model models.Model) (models.Model, error) {
	devicemodel := model.(*models.DeviceModel)
	devicemodel.CreatedAt = time.Now()
	devicemodel.LastUpdatedAt = time.Now()

	pf.modelDB.Save(devicemodel)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	return devicemodel, nil
}

func (pf *deviceModelFactory) List(owner factory.Owner, query *models.Query) ([]models.Model, error) {
	values := []*models.DeviceModel{}

	pf.modelDB.Where("userID = ?", owner.User()).Find(values)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}

	devicemodels := []models.Model{}
	for _, devicemodel := range values {
		devicemodels = append(devicemodels, devicemodel)
	}
	return devicemodels, nil
}

func (pf *deviceModelFactory) Get(owner factory.Owner, ID string) (models.Model, error) {
	devicemodel := models.DeviceModel{}

	pf.modelDB.Where("userID = ? AND ID = ?", owner.User(), ID).Find(&devicemodel)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	return &devicemodel, nil
}

func (pf *deviceModelFactory) Delete(owner factory.Owner, ID string) error {
	pf.modelDB.Delete(&models.DeviceModel{
		UserID: owner.User(),
		ID:     ID,
	})
	return factory.Error(pf.modelDB)
}

func (pf *deviceModelFactory) Update(owner factory.Owner, model models.Model) error {
	devicemodel := model.(*models.DeviceModel)
	devicemodel.LastUpdatedAt = time.Now()

	// Check wethere the workshop exist
	if _, err := pf.Get(owner, devicemodel.ID); err != nil {
		return err
	}
	pf.modelDB.Save(devicemodel)
	return factory.Error(pf.modelDB)
}
