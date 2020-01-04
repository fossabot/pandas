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
package broadcast

import (
	"github.com/spf13/pflag"
)

type ServingOptions struct {
	Method   string
	User     string
	Password string
	Hosts    string
}

func NewServingOptions() *ServingOptions {
	s := ServingOptions{
		Method: "inproc",
	}
	return &s
}

func (s *ServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Method, "broadcast", s.Method,
		"broadcastization method(inproc, rabbitmq, raft)")

	fs.StringVar(&s.User, "broadcast-user", s.User,
		"is broadcastization using rabbitmq, the user name must be specified.")

	fs.StringVar(&s.Password, "broadcast-pwd", s.Password,
		"if broadcastization using rabbitmq, the password must be specified)")

	fs.StringVar(&s.Hosts, "broadcast-hosts", s.Hosts,
		"if broadcastization using rabbitmq,  the hosts muste be specified)")
}
