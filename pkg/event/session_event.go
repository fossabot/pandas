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

// SessionCreateEvent
type SessionCreateEvent struct {
	BrokerId   string `json:"brokerId"`   // Broker identifier where event come from
	Type       uint32 `json:"type"`       // Event type
	ProductID  string `json:"productId"`  // Product identifier where event come from
	DeviceName string `json:"deviceName"` // DeviceName identifier where event come from
	ClientId   string `json:"clientID"`   // Client identifier where event come from
	Persistent bool   `json:"persistent"` // Whether the session is persistent
	Retain     bool   `json:"retain"`
}

func (p *SessionCreateEvent) SetBrokerId(brokerId string)     { p.BrokerId = brokerId }
func (p *SessionCreateEvent) SetDeviceName(deviceName string) { p.DeviceName = deviceName }
func (p *SessionCreateEvent) SetProductId(productId string)   { p.ProductID = productId }
func (p *SessionCreateEvent) SetType(eventType uint32)        { p.Type = eventType }
func (p *SessionCreateEvent) SetClientId(clientID string)     { p.ClientId = clientID }
func (p *SessionCreateEvent) GetBrokerId() string             { return p.BrokerId }
func (p *SessionCreateEvent) GetType() uint32                 { return SessionCreate }
func (p *SessionCreateEvent) GetClientId() string             { return p.ClientId }
func (p *SessionCreateEvent) Serialize() ([]byte, error)      { return nil, nil }
func (p *SessionCreateEvent) GetDeviceName() string           { return p.DeviceName }
func (p *SessionCreateEvent) GetProductId() string            { return p.ProductID }

// SessionDestroyEvent
type SessionDestroyEvent struct {
	BrokerId   string `json:"brokerId"`   // Broker identifier where event come from
	Type       uint32 `json:"type"`       // Event type
	ProductID  string `json:"productId"`  // Product identifier where event come from
	DeviceName string `json:"deviceName"` // DeviceName identifier where event come from
	ClientId   string `json:"clientID"`   // Client identifier where event come from
	Persistent bool   `json:"persistent"` // Whether the session is persistent
	Retain     bool   `json:"retain"`
}

func (p *SessionDestroyEvent) SetBrokerId(brokerId string)     { p.BrokerId = brokerId }
func (p *SessionDestroyEvent) SetType(eventType uint32)        { p.Type = eventType }
func (p *SessionDestroyEvent) SetDeviceName(deviceName string) { p.DeviceName = deviceName }
func (p *SessionDestroyEvent) SetProductId(productId string)   { p.ProductID = productId }
func (p *SessionDestroyEvent) SetClientId(clientID string)     { p.ClientId = clientID }
func (p *SessionDestroyEvent) GetBrokerId() string             { return p.BrokerId }
func (p *SessionDestroyEvent) GetType() uint32                 { return SessionDestroy }
func (p *SessionDestroyEvent) GetClientId() string             { return p.ClientId }
func (p *SessionDestroyEvent) Serialize() ([]byte, error)      { return nil, nil }
func (p *SessionDestroyEvent) GetDeviceName() string           { return p.DeviceName }
func (p *SessionDestroyEvent) GetProductId() string            { return p.ProductID }

// SessionResumeEvent
type SessionResumeEvent struct {
	BrokerId   string `json:"brokerId"`   // Broker identifier where event come from
	Type       uint32 `json:"type"`       // Event type
	DeviceName string `json:"deviceName"` // DeviceName identifier where event come from
	ClientId   string `json:"clientID"`   // Client identifier where event come from
	ProductID  string `json:"productId"`  // Product identifier where event come from
}

func (p *SessionResumeEvent) SetBrokerId(brokerId string)     { p.BrokerId = brokerId }
func (p *SessionResumeEvent) SetType(eventType uint32)        { p.Type = eventType }
func (p *SessionResumeEvent) SetClientId(clientID string)     { p.ClientId = clientID }
func (p *SessionResumeEvent) GetBrokerId() string             { return p.BrokerId }
func (p *SessionResumeEvent) GetType() uint32                 { return SessionResume }
func (p *SessionResumeEvent) GetClientId() string             { return p.ClientId }
func (p *SessionResumeEvent) Serialize() ([]byte, error)      { return nil, nil }
func (p *SessionResumeEvent) SetDeviceName(deviceName string) { p.DeviceName = deviceName }
func (p *SessionResumeEvent) SetProductId(productId string)   { p.ProductID = productId }
func (p *SessionResumeEvent) GetDeviceName() string           { return p.DeviceName }
func (p *SessionResumeEvent) GetProductId() string            { return p.ProductID }
