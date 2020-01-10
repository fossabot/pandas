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
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/models/cache"
	modelsoptions "github.com/cloustone/pandas/models/options"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type viewFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newViewFactory(servingOptions *modelsoptions.ServingOptions) Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-projects.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.View{})
	return &viewFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *viewFactory) Save(owner Owner, obj models.Model) (models.Model, error) {
	view := obj.(*models.View)
	view.CreatedAt = time.Now()
	pf.modelDB.Save(view)

	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	// update cache
	pf.cache.Set(newCacheID(owner, view.ID), view)
	return view, nil
}

func (pf *viewFactory) List(owner Owner, query *models.Query) ([]models.Model, error) {
	views := []*models.Project{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(views)

	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	results := []models.Model{}
	for _, view := range views {
		results = append(results, view)
	}
	return results, nil
}

func (pf *viewFactory) Get(owner Owner, viewID string) (models.Model, error) {
	view := models.View{}
	if err := pf.cache.Get(newCacheID(owner, viewID), &view); err == nil {
		return &view, nil
	}
	pf.modelDB.Where("userId = ? AND projectId = ?", owner.User(), viewID).Find(&view)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}

	return &view, nil
}

func (pf *viewFactory) Delete(owner Owner, viewID string) error {
	pf.modelDB.Delete(&models.View{
		ProjectID: owner.Project(),
		ID:        viewID,
	})
	pf.cache.Delete(newCacheID(owner, viewID))
	return nil
}

func (pf *viewFactory) Update(Owner, models.Model) error {
	return nil
}
