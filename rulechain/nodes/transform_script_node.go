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

type transformScriptNode struct {
	bareNode
	Script string `json:"script" yaml:"script"`
}

type transformScriptNodeFactory struct{}

func (f transformScriptNodeFactory) Name() string     { return "TransformScriptNode" }
func (f transformScriptNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }

func (f transformScriptNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &transformScriptNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *transformScriptNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine()
	newMessage, err := scriptEngine.ScriptOnMessage(msg, n.Script)
	if err != nil {
		return failureLabelNode.Handle(msg)
	}
	return successLabelNode.Handle(newMessage)
}
