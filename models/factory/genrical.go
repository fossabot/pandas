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
	"github.com/cloustone/pandas/models"
	modeloptions "github.com/cloustone/pandas/models/options"
	"github.com/jinzhu/gorm"
)

type genericalFactory struct {
	modelDB *gorm.DB
}

func (pf *genericalFactory) initialize(factoryServingOptions *modeloptions.ServingOptions) error {
	modelDB, err := gorm.Open(factoryServingOptions.StorePath, "pandas.db")
	if err != nil {
		return err
	}
	modelDB.AutoMigrate(&models.Project{})
	pf.modelDB = modelDB
	return nil
}

func (pf *genericalFactory) Save(owner Owner, model models.Model) (models.Model, error) {
	pf.modelDB.Save(model)
	if err := Error(pf.modelDB); err != nil {
		return nil, err
	}
	return nil, nil // TODO
}

func (pf *genericalFactory) List(owner Owner, query *models.Query) ([]models.Model, error) {
	values := []models.Model{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(values)
	if err := Error(pf.modelDB); err != nil {
		return nil, err
	}
	return values, nil
}

func (pf *genericalFactory) Get(owner Owner, ID string) (models.Model, error) {
	/*
		var value interface{}
		pf.modelDB.Where("userId = ? AND Id = ?", owner.User(), ID).Find(value)
		if err := Error(pf.modelDB); err != nil {
			return nil, err
		}
		return value, nil
	*/
	return nil, nil
}

func (pf *genericalFactory) Delete(owner Owner, ID string) error {
	return Error(pf.modelDB)
}

func (pf *genericalFactory) Update(owner Owner, model models.Model) error {
	if _, err := pf.Get(owner, "TODO"); err != nil {
		return err
	}
	pf.modelDB.Save(model)
	return Error(pf.modelDB)
}
