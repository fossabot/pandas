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
	"time"

	"github.com/cloustone/pandas/models"
	"github.com/sirupsen/logrus"
)

const ClearAlarmNodeName = "ClearAlarmNode"

type clearAlarmNodeFactory struct{}

type clearAlarmNode struct {
	bareNode
	DetailBuilderScript string     `json:"detailBuilderScript" yaml:"detailBuilderScript"`
	AlarmType           string     `json:"alarmType" yaml:"alarmType"`
	AlarmSeverity       string     `json:"alarmSeverity" yaml:"alarmSeverity"`
	Propagate           string     `json:"propagate" yaml:"propagate"`
	AlarmStartTime      *time.Time `json:"alarmStartTime" yaml:"alarmStartTime"`
	AlarmEndTime        *time.Time `json:"alarmEndTime" yaml:"alarmEndTime"`
}

func (f clearAlarmNodeFactory) Name() string     { return ClearAlarmNodeName }
func (f clearAlarmNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f clearAlarmNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Created", "Updated", "Failure"}
	node := &clearAlarmNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *clearAlarmNode) Handle(msg models.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	return nil
}
