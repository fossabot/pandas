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

type messageGeneratorNode struct {
	bareNode
	DetailBuilderScript string `json:"detail_builder_script" yaml:"detail_builder_script"`
	FrequenceInSecond   int32  `json:"frequency" yaml:"frequency"`
}

type messageGeneratorNodeFactory struct{}

func (f messageGeneratorNodeFactory) Name() string     { return "MessageGeneratorNode" }
func (f messageGeneratorNodeFactory) Category() string { return NODE_CATEGORY_ACTION }

func (f messageGeneratorNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Created", "Updated"}
	node := &messageGeneratorNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *messageGeneratorNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	createdLabelNode := n.GetLinkedNode("Created")
	updatedLabelNode := n.GetLinkedNode("Updated")
	if createdLabelNode == nil || updatedLabelNode == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}

	return nil
}
