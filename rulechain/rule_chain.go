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
package rulechain

import (
	"fmt"
	"strconv"

	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/rulechain/manifest"
	"github.com/cloustone/pandas/rulechain/nodes"
	"github.com/sirupsen/logrus"
)

type RuleChain interface {
	Name() string
	ApplyMessage(models.Message) error
}

// RuleChain manage all nodes in this chain, validate and apply data
// Only one input node exist in chain as precondition, and with many output nodes
// Relations within nodes is maintained by link object
type ruleChain struct {
	name            string                 `json:"name" yaml:"name"`
	firstRuleNodeId string                 `json:"firstRuleNodeId" yaml:"firstRuleNodeid"`
	root            bool                   `json:"root" yaml:"boot"`
	debugMode       bool                   `json:"debugMode" yaml:"debugMode"`
	configuration   map[string]interface{} `json:"configuration" yaml:"configuration"`
	nodes           map[string]nodes.Node  `json:"nodes" yaml:"nodes"`
}

func NewRuleChain(data []byte) (RuleChain, error) {
	manifest, err := manifest.New(data)
	if err != nil {
		logrus.WithError(err).Errorf("invalidi manifest file")
		return nil, err
	}
	return NewWithManifest(manifest)
}

// NewWithManifest create rule chain by user's manifest file
func NewWithManifest(m *manifest.Manifest) (RuleChain, error) {
	r := &ruleChain{
		name:            m.RuleChain.Name,
		firstRuleNodeId: m.RuleChain.FirstRuleNodeId,
		root:            m.RuleChain.Root,
		debugMode:       m.RuleChain.DebugMode,
		configuration:   m.RuleChain.Configuration,
		nodes:           make(map[string]nodes.Node),
	}
	// Create All nodes
	for _, n := range m.Metadata.Nodes {
		metadata := nodes.NewMetadataWithValues(n.Configuration).With("debugMode", r.debugMode)
		node, err := nodes.NewNode(n.Type, n.Name, metadata)
		if err != nil {
			return nil, err
		}
		if _, found := r.nodes[n.Name]; found {
			return nil, fmt.Errorf("node '%s' already exist in rulechain '%s'", n.Name, m.RuleChain.Name)
		}
		r.nodes[n.Name] = node
	}

	// Create All node connections
	for _, conn := range m.Metadata.Connections {
		originalNode, found := r.nodes[strconv.Itoa(conn.FromIndex)]
		if !found {
			return nil, fmt.Errorf("original node '%s' no exist in rulechain '%s'", originalNode.Name, m.RuleChain.Name)
		}
		targetNode, found := r.nodes[strconv.Itoa(conn.ToIndex)]
		if !found {
			return nil, fmt.Errorf("target node '%s' no exist in rulechain '%s'", targetNode.Name, m.RuleChain.Name)
		}
		originalNode.AddLinkedNode(conn.Type, targetNode)
	}
	// some labels must be satisified
	for name, node := range r.nodes {
		targetNodes := node.GetLinkedNodes()
		mustLabels := node.MustLabels()
		for _, label := range mustLabels {
			if _, found := targetNodes[label]; !found {
				return nil, fmt.Errorf("the label '%s' in node '%s' no exist'", label, name)
			}
		}
	}

	return r, nil
}

func (r *ruleChain) Name() string { return r.name }

// ApplyMessage push message into node chain, so that data processing will be triggered
func (r *ruleChain) ApplyMessage(msg models.Message) error {
	if node, found := r.nodes[r.firstRuleNodeId]; found {
		return node.Handle(msg)
	}
	return fmt.Errorf("node chain '%s' has no valid node", r.Name)
}
