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
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultVersion          = "1.0"
	defaultIcon             = "default_device.png"
	modelDefinitionKindName = "deviceModel"
)

type ModelDefinition struct {
	Name              string             `json:"name" yaml:"name"`
	Kind              string             `json:"kind" yaml:"kind"`
	Description       string             `json:"description" yaml:"description"`
	Domain            string             `json:"domain" yaml:"domain"`
	Version           string             `json:"version" yaml:"version"`
	IsLogical         bool               `json:"logical" yaml:"logical"`
	IsCompound        bool               `json:"compound" yaml:"compound"`
	Icon              string             `json:"icon" yaml:"icon"`
	Endpoints         []*Endpoint        `json:"endpoints" yaml:"endpoints"`
	DataModels        []*DataModel       `json:"dataModels" yaml:"dataModels"`
	ChildDeviceModels []ChildDeviceModel `json:"childDeviceModels" yaml:"childDeviceModels"`
}

type Endpoint struct {
	Path       string `json:"path" yaml:"path"`
	DataModel  string `json:"dataModel" yaml:"dataModel"`
	Permission string `json:"permission" yaml:"permission"`
}

type DataModel struct {
	Name   string       `json:"name" yaml:"name"`
	Fields []*DataField `json:"fields" yaml:"fields"`
}

type DataField struct {
	Key          string `json:"key" yaml:"key"`
	Type         string `json:"type" yaml:"type"`
	DefaultValue string `json:"defaultValue" yaml:"defaultValue"`
}
type ChildDeviceModel struct {
	Name string `json:"name" yaml:"name"`
}

func ParseModelDefinition(data []byte, dataType string) (*ModelDefinition, error) {
	md := &ModelDefinition{}
	var err error

	switch dataType {
	case "json":
		err = json.Unmarshal(data, md)
	case "yaml":
		err = yaml.Unmarshal(data, md)
	default:
		return nil, fmt.Errorf("invalid model file type '%s'", dataType)
	}
	if err != nil {
		logrus.WithError(err).Errorf("parse model defintion file failed")
		return nil, err
	}

	// Check model definition file's validaity
	if err := normalizeDeviceModelDefinition(md); err != nil {
		logrus.WithError(err).Errorf("invalid model defintion")
		return nil, err
	}
	return md, nil
}

func normalizeDeviceModelDefinition(model *ModelDefinition) error {
	if model.Name == "" || model.Kind == "" {
		return errors.New("Model name or kind is not specified")
	}
	if model.Kind != modelDefinitionKindName {
		return errors.New("Model defintion kind is not device model")
	}
	if model.Version == "" {
		model.Version = defaultVersion
	}
	if model.Icon == "" {
		model.Icon = defaultIcon
	}
	// Check data model defintion
	dataModels := make(map[string]bool)
	for _, dataModel := range model.DataModels {
		if dataModel.Name == "" {
			return errors.New("data model name is not specified")
		}
		if len(dataModel.Fields) == 0 {
			return fmt.Errorf("there are no fields for data model '%s'", dataModel.Name)
		}
		if _, found := dataModels[dataModel.Name]; found {
			return fmt.Errorf("data model '%s' alread exist", dataModel.Name)
		}
		dataModels[dataModel.Name] = true
		fields := make(map[string]bool)
		for _, field := range dataModel.Fields {
			if field.Key == "" || field.Type == "" {
				return fmt.Errorf("data filed key or type is not specifed")
			}
			if _, found := fields[field.Key]; found {
				return fmt.Errorf("data filed '%s' already exist in '%s'", field.Key, dataModel.Name)
			}
			fields[field.Key] = true
		}
	}
	//  Endpoints
	endpoints := make(map[string]bool)
	for _, ep := range model.Endpoints {
		if ep.Path == "" {
			return fmt.Errorf("no path for endpoint of device model '%s'", model.Name)
		}
		if _, found := endpoints[ep.Path]; found {
			return fmt.Errorf("path '%s' already exist in device model '%s'", ep.Path, model.Name)
		}
		if _, found := dataModels[ep.DataModel]; !found {
			return fmt.Errorf("invalid data model '%s'", ep.DataModel)
		}
		if ep.Permission == "" {
			ep.Permission = "all"
		}
		switch ep.Permission {
		case "all", "readonly", "write":
		default:
			return fmt.Errorf("invalid permission for endpoint '%s'", ep.Path)
		}
		endpoints[ep.Path] = true
	}
	// Child
	childs := make(map[string]bool)
	for _, child := range model.ChildDeviceModels {
		if child.Name == "" {
			return fmt.Errorf("child device name is null")
		}
		if _, found := childs[child.Name]; found {
			return fmt.Errorf("child device '%s' already exist", child.Name)
		}
		childs[child.Name] = true
	}
	return nil
}
