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
package options

import (
	"github.com/cloustone/pandas/dmms"
	modeloptions "github.com/cloustone/pandas/models/options"
	genericoptions "github.com/cloustone/pandas/pkg/server/options"
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	SecureServing   *genericoptions.SecureServingOptions
	LocationServing *genericoptions.LocationServingOptions
	ModelServing    *modeloptions.ServingOptions
	DeviceServing   *dmms.ServingOptions
	SSHKeyfile      string
	SSHUser         string
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		SecureServing:   genericoptions.NewSecureServingOptions("dmms"),
		LocationServing: genericoptions.NewLocationServingOptions(),
		ModelServing:    modeloptions.NewServingOptions(),
		DeviceServing:   dmms.NewServingOptions(),
	}
	return &s
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	s.SecureServing.AddFlags(fs)
	s.LocationServing.AddFlags(fs)
	s.ModelServing.AddFlags(fs)

	fs.StringVar(&s.SSHUser, "ssh-user", s.SSHUser,
		"If non-empty, use secure SSH proxy to the nodes, using this user name")

	fs.StringVar(&s.SSHKeyfile, "ssh-keyfile", s.SSHKeyfile,
		"If non-empty, use secure SSH proxy to the nodes, using this user keyfile")
}
