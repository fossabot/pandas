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
package server

import (
	"github.com/cloustone/pandas/apimachinery/restapi/operations/rulechain"
	"github.com/cloustone/pandas/models"

	"github.com/go-openapi/runtime/middleware"
)

func GetRuleChains(params rulechain.GetRuleChainsParams, principal *models.Principal) middleware.Responder {
	return &rulechain.GetRuleChainsOK{}
}

func GetRuleChain(params rulechain.GetRuleChainParams, principal *models.Principal) middleware.Responder {
	return &rulechain.GetRuleChainOK{}
}

func SaveRuleChain(params rulechain.SaveRuleChainParams, principal *models.Principal) middleware.Responder {
	return &rulechain.SaveRuleChainOK{}
}

func GetRuleChainMetadata(params rulechain.GetRuleChainMetadataParams, principal *models.Principal) middleware.Responder {
	return &rulechain.GetRuleChainMetadataOK{}
}

func SaveRuleChainMetadata(params rulechain.SaveRuleChainMetadataParams, principal *models.Principal) middleware.Responder {
	return &rulechain.SaveRuleChainMetadataOK{}
}

func DeleteRuleChain(params rulechain.DeleteRuleChainParams, principal *models.Principal) middleware.Responder {
	return &rulechain.DeleteRuleChainOK{}
}

func DownloadRuleChain(params rulechain.DownloadRuleChainParams, principal *models.Principal) middleware.Responder {
	return &rulechain.DownloadRuleChainOK{}
}

func UploadRuleChain(params rulechain.UploadRuleChainParams, principal *models.Principal) middleware.Responder {
	return &rulechain.UploadRuleChainOK{}
}

func SetRootRuleChain(params rulechain.SetRootRuleChainParams, principal *models.Principal) middleware.Responder {
	return &rulechain.SetRootRuleChainOK{}
}
