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

type deviceInProjectFactory struct {
	modelDB *gorm.DB
}

func (pf *deviceInProjectFactory) initialize(factoryServingOptions *FactoryServingOptions) error {
	modelDB, err := gorm.Open(factoryServingOptions.StorePath, "pandas-devices.db")
	if err != nil {
		return err
	}
	modelDB.AutoMigrate(&models.Project{})
	pf.modelDB = modelDB
	return nil
}

func (pf *deviceInProjectFactory) Save(owner Owner, model models.Model) (models.Model, error) {
	deviceInProject := model.(*models.Project)
	deviceInProject.CreatedAt = time.Now()
	deviceInProject.LastUpdatedAt = time.Now()

	pf.modelDB.Save(deviceInProject)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return deviceInProject, nil
}

func (pf *deviceInProjectFactory) List(owner Owner, query *models.Query) ([]models.Model, error) {
	values := []*models.Project{}

	pf.modelDB.Where("userId = ?", owner.User()).Find(values)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}

	deviceInProjects := []models.Model{}
	for _, deviceInProject := range values {
		deviceInProjects = append(deviceInProjects, deviceInProject)
	}
	return deviceInProjects, nil
}

func (pf *deviceInProjectFactory) Get(owner Owner, deviceInProjectId string) (models.Model, error) {
	deviceInProject := models.Project{}

	pf.modelDB.Where("userId = ? AND deviceInProjectId = ?", owner.User(), deviceInProjectId).Find(&deviceInProject)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return &deviceInProject, nil
}

func (pf *deviceInProjectFactory) Delete(owner Owner, deviceInProjectID string) error {
	pf.modelDB.Delete(&models.Project{
		UserID: owner.User(),
		ID:     deviceInProjectID,
	})
	return getModelError(pf.modelDB)
}

func (pf *deviceInProjectFactory) Update(owner Owner, model models.Model) error {
	deviceInProject := model.(*models.Project)
	deviceInProject.LastUpdatedAt = time.Now()

	// Check wethere the deviceInProject exist
	if _, err := pf.Get(owner, deviceInProject.ID); err != nil {
		return err
	}
	pf.modelDB.Save(deviceInProject)
	return getModelError(pf.modelDB)
}
