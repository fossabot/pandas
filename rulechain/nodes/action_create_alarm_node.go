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

type createAlarmNode struct {
	bareNode
	DetailBuilderScript string `json:"detailBuilderScript" yaml:"detailBuilderScript"`
	AlarmType           string `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity       string `json:"alarmSeverity" yaml:"alarmSeverity"`
	Propagate           string `json:"propagate" yaml:"propagate"`
	AlarmStartTime      string `json:"alarmStartTime" yaml:"alarmStartTime"`
	AlarmEndTime        string `json:"alarmEndTime" yaml:"alarmEndTime"`
}

type createAlarmNodeFactory struct{}

func (f createAlarmNodeFactory) Name() string     { return "CreateAlarmNode" }
func (f createAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f createAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Created", "Updated", "Failure"}
	node := &createAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *createAlarmNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	node1 := n.GetLinkedNode("Created")
	node2 := n.GetLinkedNode("Updated")
	node3 := n.GetLinkedNode("Failure")
	if node1 == nil || node2 == nil || node3 == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}

	return nil
}
