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
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

const (
	NODE_CATEGORY_FILTER     = "filter"
	NODE_CATEGORY_ACTION     = "action"
	NODE_CATEGORY_ENRICHMENT = "enrichment"
	NODE_CATEGORY_TRANSFORM  = "transform"
	NODE_CATEGORY_EXTERNAL   = "external"
	NODE_CATEGORY_OTHERS     = "others"
)

// Factory is node's factory to create node based on metadata
// factory also manage node's metadta description which can be used by other
// service to present node in web
type Factory interface {
	Name() string
	Category() string
	Create(id string, meta Metadata) (Node, error)
	//Metadata() string
}

var (
	// allNodeFactories hold all node's factory
	allNodeFactories map[string]Factory = make(map[string]Factory)

	// allNodeCategories hold node's metadata by category
	allNodeCategories map[string][]string = make(map[string][]string)

	// allNodeConfigs hold node's config data using map to index node's metadata directlly
	allNodeConfigs map[string]string = make(map[string]string)
)

// RegisterFactory add a new node factory and classify its category for
// metadata description
func RegisterFactory(f Factory) {
	allNodeFactories[f.Name()] = f

	if allNodeCategories[f.Category()] == nil {
		allNodeCategories[f.Category()] = []string{}
	}
	allNodeCategories[f.Category()] = append(allNodeCategories[f.Category()], f.Name())
	configFile := AssetPath + "/" + f.Name() + ".js"
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.Fatalf("asset file '%s' no exist", configFile)
	}
	allNodeConfigs[f.Name()] = string(buf)
}

// NewNode is the only way to create a new node
func NewNode(nodeType string, id string, meta Metadata) (Node, error) {
	if f, found := allNodeFactories[nodeType]; found {
		return f.Create(id, meta)
	}
	return nil, fmt.Errorf("invalid node type '%s'", nodeType)
}

// GetAllNodeConfigs returan all node's static description used by user to list nodes
func GetAllNodeConfigs() map[string]string { return allNodeConfigs }

// GetCategoryNodes return specified category's all nodes
func GetCategoryNodes() map[string][]string { return allNodeCategories }

// GetNodeMeta return a node's static metadata
func GetNodeConfigs(name string) (string, error) {
	if c, found := allNodeConfigs[name]; found {
		return c, nil
	}
	return "", errors.New("not found")
}
