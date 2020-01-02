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
package nodes

import (
	"github.com/cloustone/pandas/models"
	"github.com/sirupsen/logrus"
)

const InputNodeName = "InputNode"

type inputNode struct {
	bareNode
	Script string `json:"script" yaml:"script"`
}

type inputNodeFactory struct{}

func (f inputNodeFactory) Name() string     { return "InputNode" }
func (f inputNodeFactory) Category() string { return NODE_CATEGORY_OTHERS }

func (f inputNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{}
	node := &inputNode{
		bareNode: newBareNode(InputNodeName, id, meta, labels),
	}
	return node, nil
}

func (n *inputNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	nodes := n.GetLinkedNodes()
	for _, node := range nodes {
		return node.Handle(msg)
	}
	return nil
}
