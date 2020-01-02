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
package manifest

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type AdditionalInfo struct {
	Description string `json:"description" yaml:"description"`
	LayoutX     int32  `json:"layoutX" yaml:"layoutX"`
	LayoutY     int32  `json:"layoutY" yaml:"layoutY"`
}

type RuleChain struct {
	Name            string                 `json:"name" yaml:"name"`
	Id              string                 `json:"id" yaml"id"`
	FirstRuleNodeId string                 `json:"firstRuleNodeId" yaml:"firstRuleNodeid"`
	Root            bool                   `json:"root" yaml:"boot"`
	DebugMode       bool                   `json:"debugMode" yaml:"debugMode"`
	Configuration   map[string]interface{} `json:"configuration" yaml:"configuration"`
}

type Node struct {
	AdditionalInfo AdditionalInfo         `json:"additionalInfo" yaml:"additionalInfo"`
	Type           string                 `json:"type" yaml:"type"`
	Name           string                 `json:"name" yaml:"name"`
	DebugMode      bool                   `json:"debugMode" yaml:"debugMode"`
	Configuration  map[string]interface{} `json:"configuration" yaml:"configuration"`
}

type NodeConnection struct {
	FromIndex int    `json:"fromIndex" yaml:"fromIndex"`
	ToIndex   int    `json:"toIndex" yaml:"toIndex"`
	Type      string `json:"type" yaml:"type"`
}

type Metadata struct {
	firstNodeIndex       int32                 `json:"firstNodeIndex" yaml:"firstNodeIndex"`
	Nodes                []Node                `json:"nodes" yaml:"nodes"`
	Connections          []NodeConnection      `json:"connections" yaml:"connections"`
	RuleChainConnections []RuleChainConnection `json:"ruleChainConnections" yaml:"ruleChainConnections"`
}

type TargetRuleChainId struct {
	EntityType string `json:"entityType" yaml:"entityType"`
	Id         string `json:"id" yaml:"id"`
}

type RuleChainConnection struct {
	FromIndex         int               `json:"fromIndex" yaml:"fromIndex"`
	TargetRuleChainId TargetRuleChainId `json:"targetRuleChainId" yaml:"targetRuleChainId"`
	AdditionalInfo    AdditionalInfo    `json:"additionalInfo" yaml:"additionalInfo"`
	Type              string            `json:"type" yaml:"type"`
}

type Manifest struct {
	RuleChain RuleChain `json:"ruleChain" yaml:"ruleChain"`
	Metadata  Metadata  `json:"metadata" yaml:"metadata"`
}

func New(data []byte) (*Manifest, error) {
	m := &Manifest{}
	if err := json.Unmarshal(data, m); err != nil {
		logrus.WithError(err).Errorf("invalid node chain manifest file")
		return nil, err
	}
	return m, nil
}
