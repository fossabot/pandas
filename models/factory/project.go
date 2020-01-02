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
	"github.com/jinzhu/gorm"
)

type projectFactory struct {
	modelDB *gorm.DB
}

func (pf *projectFactory) initialize(factoryServingOptions *FactoryServingOptions) error {
	modelDB, err := gorm.Open(factoryServingOptions.StorePath, "pandas-project.db")
	if err != nil {
		return err
	}
	modelDB.AutoMigrate(&models.Project{})
	pf.modelDB = modelDB
	return nil
}

func (pf *projectFactory) Save(owner Owner, model models.Model) (models.Model, error) {
	project := model.(*models.Project)
	project.CreatedAt = time.Now()
	project.LastUpdatedAt = time.Now()

	pf.modelDB.Save(project)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return project, nil
}

func (pf *projectFactory) List(owner Owner, query *models.Query) ([]models.Model, error) {
	values := []*models.Project{}

	pf.modelDB.Where("userId = ?", owner.User()).Find(values)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}

	projects := []models.Model{}
	for _, project := range values {
		projects = append(projects, project)
	}
	return projects, nil
}

func (pf *projectFactory) Get(owner Owner, projectId string) (models.Model, error) {
	project := models.Project{}

	pf.modelDB.Where("userId = ? AND projectId = ?", owner.User(), projectId).Find(&project)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return &project, nil
}

func (pf *projectFactory) Delete(owner Owner, projectID string) error {
	pf.modelDB.Delete(&models.Project{
		UserID: owner.User(),
		ID:     projectID,
	})
	return getModelError(pf.modelDB)
}

func (pf *projectFactory) Update(owner Owner, model models.Model) error {
	project := model.(*models.Project)
	project.LastUpdatedAt = time.Now()

	// Check wethere the project exist
	if _, err := pf.Get(owner, project.ID); err != nil {
		return err
	}
	pf.modelDB.Save(project)
	return getModelError(pf.modelDB)
}
