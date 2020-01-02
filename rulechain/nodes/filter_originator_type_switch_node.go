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
	"fmt"

	"github.com/cloustone/pandas/models"
	"github.com/sirupsen/logrus"
)

type originatorTypeSwitchNode struct {
	bareNode
}

type originatorTypeSwitchNodeFactory struct{}

func (f originatorTypeSwitchNodeFactory) Name() string     { return "OriginatorTypeSwitchNode" }
func (f originatorTypeSwitchNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f originatorTypeSwitchNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{}
	node := &originatorTypeSwitchNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *originatorTypeSwitchNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	nodes := n.GetLinkedNodes()
	originatorType := msg.GetOriginator()

	for label, node := range nodes {
		if originatorType == label {
			return node.Handle(msg)
		}
	}
	// not found
	return fmt.Errorf("%s no label to handle message", n.Name())
}
