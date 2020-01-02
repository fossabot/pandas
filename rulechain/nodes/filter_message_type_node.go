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

type messageTypeFilterNode struct {
	bareNode
	MessageTypes []string `json:"messageTypes" yaml:"messageTypes"`
}

type messageTypeFilterNodeFactory struct{}

func (f messageTypeFilterNodeFactory) Name() string     { return "MessageTypeFilterNode" }
func (f messageTypeFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f messageTypeFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &messageTypeFilterNode{
		bareNode:     newBareNode(f.Name(), id, meta, labels),
		MessageTypes: []string{},
	}
	return decodePath(meta, node)
}

func (n *messageTypeFilterNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")
	if trueLabelNode == nil || falseLabelNode == nil {
		return fmt.Errorf("no true or false label linked node in %s", n.Name())
	}
	messageType := msg.GetType()

	// TODO: how to resolve user customized message type dynamically
	//userMessageType := msg.GetMetadata().GetKeyValue(n.Metadata().MessageTypeKey)
	userMessageType := "TODO"
	for _, filterType := range n.MessageTypes {
		if filterType == messageType || filterType == userMessageType {
			return trueLabelNode.Handle(msg)
		}
	}
	return falseLabelNode.Handle(msg)
}
