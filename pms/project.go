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

type projectFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newProjectFactory(servingOptions *modelsoptions.ServingOptions) factory.Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-workshops.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.Project{})
	return &projectFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *projectFactory) Save(owner factory.Owner, model models.Model) (models.Model, error) {
	project := model.(*models.Project)
	project.CreatedAt = time.Now()
	project.LastUpdatedAt = time.Now()

	pf.modelDB.Save(project)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	return project, nil
}

func (pf *projectFactory) List(owner factory.Owner, query *models.Query) ([]models.Model, error) {
	values := []*models.Project{}

	pf.modelDB.Where("userId = ?", owner.User()).Find(values)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}

	projects := []models.Model{}
	for _, project := range values {
		projects = append(projects, project)
	}
	return projects, nil
}

func (pf *projectFactory) Get(owner factory.Owner, projectId string) (models.Model, error) {
	project := models.Project{}

	pf.modelDB.Where("userId = ? AND projectId = ?", owner.User(), projectId).Find(&project)
	if err := factory.Error(pf.modelDB); err != nil {
		return nil, err
	}
	return &project, nil
}

func (pf *projectFactory) Delete(owner factory.Owner, projectID string) error {
	pf.modelDB.Delete(&models.Project{
		UserID: owner.User(),
		ID:     projectID,
	})
	return factory.Error(pf.modelDB)
}

func (pf *projectFactory) Update(owner factory.Owner, model models.Model) error {
	project := model.(*models.Project)
	project.LastUpdatedAt = time.Now()

	// Check wethere the project exist
	if _, err := pf.Get(owner, project.ID); err != nil {
		return err
	}
	pf.modelDB.Save(project)
	return factory.Error(pf.modelDB)
}
