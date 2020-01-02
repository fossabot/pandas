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
package proxy

import "github.com/cloustone/pandas/pkg/message"

const (
	AlarmMessageTopic = "iotx-foundry-alarm"
)

type AlarmTopic struct {
	TopicName string
	Alarm     *AlarmNotification `json:"alarm"`
}

func (p *AlarmTopic) Topic() string        { return AlarmMessageTopic }
func (p *AlarmTopic) SetTopic(name string) {}
func (p *AlarmTopic) Serialize(opt message.SerializeOption) ([]byte, error) {
	return message.Serialize(p, opt)
}
func (p *AlarmTopic) Deserialize(buf []byte, opt message.SerializeOption) error {
	return message.Deserialize(buf, opt, p)
}
