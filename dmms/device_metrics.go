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

type deviceMetricsFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newDeviceMetricsFactory(servingOptions *modelsoptions.ServingOptions) factory.Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-devices.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.DeviceMetrics{})
	return &deviceMetricsFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *deviceMetricsFactory) Save(owner factory.Owner, obj models.Model) (models.Model, error) {
	deviceMetrics := obj.(*models.DeviceMetrics)
	deviceMetrics.CreatedAt = time.Now()
	pf.modelDB.Save(deviceMetrics)

	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	return deviceMetrics, nil
}

func (pf *deviceMetricsFactory) List(owner factory.Owner, query *models.Query) ([]models.Model, error) {
	deviceMetricss := []*models.DeviceMetrics{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(deviceMetricss)

	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	results := []models.Model{}
	for _, deviceMetrics := range deviceMetricss {
		results = append(results, deviceMetrics)
	}
	return results, nil
}

func (pf *deviceMetricsFactory) Get(factory.Owner, string) (models.Model, error) {
	return nil, nil
}

func (pf *deviceMetricsFactory) Delete(factory.Owner, string) error {
	return nil
}

func (pf *deviceMetricsFactory) Update(factory.Owner, models.Model) error {
	return nil
}
