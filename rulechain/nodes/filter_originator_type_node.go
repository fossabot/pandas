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

type originatorTypeFilterNode struct {
	bareNode
	Filters []string `json:"filters" yaml:"filters"`
}

type originatorFilterNodeFactory struct{}

func (f originatorFilterNodeFactory) Name() string     { return "OriginatorFilterNode" }
func (f originatorFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f originatorFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &originatorTypeFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
		Filters:  []string{},
	}
	return decodePath(meta, node)
}

func (n *originatorTypeFilterNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")

	//links := n.GetLinks()
	originatorType := msg.GetOriginator()

	for _, filter := range n.Filters {
		if originatorType == filter {
			return trueLabelNode.Handle(msg)
		}
	}
	// not found
	return falseLabelNode.Handle(msg)
}
