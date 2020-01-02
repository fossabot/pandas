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
	"errors"

	"github.com/cloustone/pandas/models"
	"github.com/sirupsen/logrus"
)

var (
	AssetPath = "../rulechain/assets"
)

type Node interface {
	Name() string
	Id() string
	Metadata() Metadata
	MustLabels() []string
	Handle(models.Message) error

	AddLinkedNode(label string, node Node)
	GetLinkedNode(label string) Node
	GetLinkedNodes() map[string]Node
}

type bareNode struct {
	name   string
	id     string
	nodes  map[string]Node
	meta   Metadata
	labels []string
}

func newBareNode(name string, id string, meta Metadata, labels []string) bareNode {
	return bareNode{
		name:   name,
		id:     id,
		nodes:  make(map[string]Node),
		meta:   meta,
		labels: labels,
	}
}

func (n *bareNode) Name() string                          { return n.name }
func (n *bareNode) WithId(id string)                      { n.id = id }
func (n *bareNode) Id() string                            { return n.id }
func (n *bareNode) MustLabels() []string                  { return n.labels }
func (n *bareNode) AddLinkedNode(label string, node Node) { n.nodes[label] = node }

func (n *bareNode) GetLinkedNode(label string) Node {
	if node, found := n.nodes[label]; found {
		return node
	}
	logrus.Error("no label '%s' in node '%s'", label, n.name)
	return nil
}

func (n *bareNode) GetLinkedNodes() map[string]Node { return n.nodes }

func (n *bareNode) Metadata() Metadata { return n.meta }

func (n *bareNode) Handle(models.Message) error { return errors.New("not implemented") }

func decodePath(meta Metadata, n Node) (Node, error) {
	if err := meta.DecodePath(n); err != nil {
		return n, err
	}
	return n, nil
}
