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

package event

type QutoChangeEvent struct {
	BrokerId   string `json:"brokerId"`   // Broker identifier where event come from
	Type       uint32 `json:"type"`       // Event type
	ClientId   string `json:"clientID"`   // Client identifier where event come from
	ProductID  string `json:"productId"`  // Product identifier where event come from
	DeviceName string `json:"deviceName"` // DeviceName identifier where event come from
	Persistent bool   `json:"persistent"` // Whether the session is persistent
	QutoId     string `json:"qutoId"`
}

func (p *QutoChangeEvent) SetBrokerId(brokerId string)     { p.BrokerId = brokerId }
func (p *QutoChangeEvent) SetType(eventType uint32)        { p.Type = eventType }
func (p *QutoChangeEvent) SetClientId(clientID string)     { p.ClientId = clientID }
func (p *QutoChangeEvent) GetBrokerId() string             { return p.BrokerId }
func (p *QutoChangeEvent) GetType() uint32                 { return QutoChange }
func (p *QutoChangeEvent) GetClientId() string             { return p.ClientId }
func (p *QutoChangeEvent) Serialize() ([]byte, error)      { return nil, nil }
func (p *QutoChangeEvent) SetDeviceName(deviceName string) { p.DeviceName = deviceName }
func (p *QutoChangeEvent) SetProductId(productId string)   { p.ProductID = productId }
func (p *QutoChangeEvent) GetDeviceName() string           { return p.DeviceName }
func (p *QutoChangeEvent) GetProductId() string            { return p.ProductID }
