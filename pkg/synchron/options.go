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

import (
	"github.com/spf13/pflag"
)

type SyncServingOptions struct {
	Method   string
	User     string
	Password string
	Hosts    string
}

func NewSyncServingOptions() *SyncServingOptions {
	s := SyncServingOptions{
		Method: "inproc",
	}
	return &s
}

func (s *SyncServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Method, "sync-method", s.Method,
		"synchronization method(inproc, rabbitmq, raft)")

	fs.StringVar(&s.User, "sync-user", s.User,
		"is synchronization using rabbitmq, the user name must be specified.")

	fs.StringVar(&s.Password, "sync-pwd", s.Password,
		"if synchronization using rabbitmq, the password must be specified)")

	fs.StringVar(&s.Hosts, "sync-hosts", s.Hosts,
		"if synchronization using rabbitmq,  the hosts muste be specified)")
}
