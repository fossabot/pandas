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

package synchron

const (
	ActionCreated = "created"
	ActionUpdated = "updated"
	ActionDeleted = "deleted"
)

type Notification struct {
	Path   string      `json:"path"`
	Action string      `json:"action"`
	Param  interface{} `json:"param"`
}

type Observer interface {
	OnSynchronizationNotified(Synchronizer, Notification)
}

type Synchronizer interface {
	WithRootPath(string) Synchronizer
	AsMember()
	Notify(Notification)
	RegisterObserver(path string, obs Observer)
}
