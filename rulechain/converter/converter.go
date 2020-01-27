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
package converter

import (
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/rulechain/grpc_rulechain_v1"
)

func NewRuleChainModel(obj *grpc_rulechain_v1.RuleChain) *models.RuleChain {
	return &models.RuleChain{}
}

func NewRuleChain(model models.Model) *grpc_rulechain_v1.RuleChain {
	return &grpc_rulechain_v1.RuleChain{}
}

func NewRuleChains(objs []models.Model) []*grpc_rulechain_v1.RuleChain {
	return []*grpc_rulechain_v1.RuleChain{}
}
