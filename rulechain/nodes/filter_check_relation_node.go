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

const (
	RelationTypeContains    = "Contains"
	RelationTypeNotContains = "NotContains"
)

type checkRelationFilterNode struct {
	bareNode
	Direction    string   `json:"direction" yaml:"direction"`
	RelationType string   `json:"relationType" yaml:"relationType"`
	InstanceType string   `json:"instanceType" yaml:"instanceType"`
	Values       []string `json:"values" yaml:"values"`
}

type checkRelationFilterNodeFactory struct{}

func (f checkRelationFilterNodeFactory) Name() string     { return "CheckRelationFilterNode" }
func (f checkRelationFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f checkRelationFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &checkRelationFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
		Values:   []string{},
	}
	return decodePath(meta, node)
}

func (n *checkRelationFilterNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")

	// direction := msg.GetDirection()
	attr := msg.GetMetadata().GetKeyValue(n.InstanceType)
	switch n.RelationType {
	case RelationTypeContains:
		for _, val := range n.Values {
			// specified attribute exist in names
			if attr == val {
				return trueLabelNode.Handle(msg)
			}
		}
		// not found
		return falseLabelNode.Handle(msg)

	case RelationTypeNotContains:
		for _, val := range n.Values {
			// specified attribute exist in names
			if attr == val {
				return falseLabelNode.Handle(msg)
			}
		}
		// not found
		return trueLabelNode.Handle(msg)
	}
	return nil
}
