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
package inproc

import "github.com/cloustone/pandas/pkg/broadcast"

const NAME = "inproc"

type InprocBroadcast struct {
	observers []broadcast.Observer
}

func NewBroadcast() broadcast.Broadcast {
	return &InprocBroadcast{
		observers: []broadcast.Observer{},
	}
}

func (s *InprocBroadcast) AsMember()                                    {}
func (s *InprocBroadcast) WithRootPath(path string) broadcast.Broadcast { return s }
func (s *InprocBroadcast) Notify(n broadcast.Notification) {
	for _, observer := range s.observers {
		observer.Onbroadcast(s, n)
	}
}

func (s *InprocBroadcast) RegisterObserver(path string, obs broadcast.Observer) {
	s.observers = append(s.observers, obs)
}
