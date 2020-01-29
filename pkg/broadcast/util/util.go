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
package util

import (
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/pkg/broadcast"
	"github.com/cloustone/pandas/pkg/broadcast/inproc"
	"github.com/cloustone/pandas/pkg/broadcast/rabbitmq"
)

var (
	globalBroadcast broadcast.Broadcast
)

// NewSynchronizer will create broadcastizer according to environment's setting
func NewBroadcast(options *broadcast.ServingOptions) broadcast.Broadcast {
	switch options.Method {
	case rabbitmq.NAME:
		return rabbitmq.NewBroadcast(options.User, options.Password, options.Hosts)
	case inproc.NAME:
		return inproc.NewBroadcast()
	}
	return nil
}

// InitializeBroadcast initialize the global broadcast with a service based path
func InitializeBroadcast(servingOptions *broadcast.ServingOptions, rootPath string) {
	globalBroadcast = NewBroadcast(servingOptions).WithRootPath(rootPath)
	globalBroadcast.AsMember()
}

// RegisterObserver register an observer on the global broadcast
func RegisterObserver(obs broadcast.Observer, path string) {
	globalBroadcast.RegisterObserver(path, obs)
}

// Notify fire the global broadcast
func Notify(objectPath string, action string, obj models.Model) {
	param, _ := obj.MarshalBinary()
	globalBroadcast.Notify(broadcast.Notification{
		ObjectPath: objectPath,
		Action:     action,
		Param:      param,
	})
}
