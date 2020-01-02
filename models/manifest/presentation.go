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
package manifest

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Node struct {
	Id                 string      `json:"id"`
	Width              int32       `json:"width"`
	Height             int32       `json:"height"`
	X                  int32       `json:"x"`
	Y                  int32       `json:"y"`
	IsLeftConnectShow  bool        `json:"isLeftConnectShow"`
	IsRightConnectShow bool        `json:"isRightConnectShow"`
	Name               string      `json:"name"`
	IsSelect           bool        `json:"isSelect"`
	InitW              int32       `json:"initW"`
	InitH              int32       `json:"initH"`
	Type               string      `json:"type"`
	ContainNodes       []string    `json:"containNodes"`
	ClassType          string      `json:"classType"`
	Attrs              []*NodeAttr `json:"attrs"`
	Icon               string      `json:"icon"`
}

type NodeAttr struct {
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Value       interface{}       `json:"value"`
	Placeholder string            `json:"placeholder"`
	Disabled    bool              `json:"disabled"`
	Options     []*NodeAttrOption `json:"options"`
}

type NodeAttrOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type NodeSource struct {
	Id     string `json:"id"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	X      int32  `json:"x"`
	Y      int32  `json:"y"`
}

type Connector struct {
	Type       string     `json:"type"`
	TargetNode NodeSource `json:"targetNode"`
	SourceNode NodeSource `json:"sourceNode"`
	IsSelect   bool       `json:"isSelect"`
}

type ModelPresentation struct {
	Nodes      []*Node      `json:"nodes"`
	Connectors []*Connector `json:"connectors"`
}

func ParseModelPresentation(data []byte, dataType string) (*ModelPresentation, error) {
	mp := &ModelPresentation{}
	var err error

	switch dataType {
	case "json":
		err = json.Unmarshal(data, mp)
	case "yaml":
		err = yaml.Unmarshal(data, mp)
	default:
		return nil, fmt.Errorf("invalid model presentnation type '%s'", dataType)
	}
	if err != nil {
		logrus.WithError(err).Errorf("invalid model presentation")
		return nil, err
	}

	// Check model definition file's validaity
	if err := normalizeDeviceModelPresentation(mp); err != nil {
		logrus.WithError(err).Errorf("invalid model presentation")
		return nil, err
	}
	return mp, nil
}

func normalizeDeviceModelPresentation(mp *ModelPresentation) error {
	return nil
}
