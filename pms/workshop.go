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
package pms

import (
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/cache"
	"github.com/cloustone/pandas/models/factory"
	modelsoptions "github.com/cloustone/pandas/models/options"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type workshopFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newWorkshopFactory(servingOptions *modelsoptions.ServingOptions) factory.Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-workshops.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.Project{})
	return &workshopFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *workshopFactory) Save(owner factory.Owner, model models.Model) (models.Model, error) {
	workshop := model.(*models.Workshop)
	workshop.CreatedAt = time.Now()
	workshop.LastUpdatedAt = time.Now()

	pf.modelDB.Save(workshop)
	if err := factory.ModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return workshop, nil
}

func (pf *workshopFactory) List(owner factory.Owner, query *models.Query) ([]models.Model, error) {
	values := []*models.Workshop{}

	pf.modelDB.Where("userID = ?", owner.User()).Find(values)
	if err := factory.ModelError(pf.modelDB); err != nil {
		return nil, err
	}

	workshops := []models.Model{}
	for _, workshop := range values {
		workshops = append(workshops, workshop)
	}
	return workshops, nil
}

func (pf *workshopFactory) Get(owner factory.Owner, ID string) (models.Model, error) {
	workshop := models.Workshop{}

	pf.modelDB.Where("userID = ? AND ID = ?", owner.User(), ID).Find(&workshop)
	if err := factory.ModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return &workshop, nil
}

func (pf *workshopFactory) Delete(owner factory.Owner, ID string) error {
	pf.modelDB.Delete(&models.Project{
		UserID: owner.User(),
		ID:     ID,
	})
	return factory.ModelError(pf.modelDB)
}

func (pf *workshopFactory) Update(owner factory.Owner, model models.Model) error {
	workshop := model.(*models.Project)
	workshop.LastUpdatedAt = time.Now()

	// Check wethere the workshop exist
	if _, err := pf.Get(owner, workshop.ID); err != nil {
		return err
	}
	pf.modelDB.Save(workshop)
	return factory.ModelError(pf.modelDB)
}
