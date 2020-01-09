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

type rulechainFactory struct {
	modelDB        *gorm.DB
	cache          cache.Cache
	servingOptions *modelsoptions.ServingOptions
}

func newRuleChainFactory(servingOptions *modelsoptions.ServingOptions) Factory {
	modelDB, err := gorm.Open(servingOptions.StorePath, "pandas-rulechains.db")
	if err != nil {
		logrus.Fatal(err)
	}
	modelDB.AutoMigrate(&models.RuleChain{})
	return &rulechainFactory{
		modelDB:        modelDB,
		cache:          cache.NewCache(servingOptions),
		servingOptions: servingOptions,
	}
}

func (pf *rulechainFactory) Save(owner Owner, obj models.Model) (models.Model, error) {
	rulechain := obj.(*models.RuleChain)
	rulechain.CreatedAt = time.Now()
	rulechain.LastUpdatedAt = time.Now()
	pf.modelDB.Save(rulechain)

	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return rulechain, nil
}

func (pf *rulechainFactory) List(owner Owner, query *models.Query) ([]models.Model, error) {
	rulechains := []*models.RuleChain{}
	pf.modelDB.Where("userId = ?", owner.User()).Find(rulechains)

	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	results := []models.Model{}
	for _, rulechain := range rulechains {
		results = append(results, rulechain)
	}
	return results, nil
}

func (pf *rulechainFactory) Get(owner Owner, rulechainID string) (models.Model, error) {
	rulechain := models.RuleChain{}

	pf.modelDB.Where("userId = ? AND chainId = ?", owner.User(), rulechainID).Find(&rulechain)
	if err := getModelError(pf.modelDB); err != nil {
		return nil, err
	}
	return &rulechain, nil
}

func (pf *rulechainFactory) Delete(owner Owner, rulechainID string) error {
	pf.modelDB.Delete(&models.RuleChain{
		UserID: owner.User(),
		ID:     rulechainID,
	})
	return getModelError(pf.modelDB)
}

func (pf *rulechainFactory) Update(owner Owner, obj models.Model) error {
	rulechain := obj.(*models.RuleChain)
	rulechain.LastUpdatedAt = time.Now()

	if _, err := pf.Get(owner, rulechain.ID); err != nil {
		return err
	}
	pf.modelDB.Save(rulechain)
	return getModelError(pf.modelDB)
}
