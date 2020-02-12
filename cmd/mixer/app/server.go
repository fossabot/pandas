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
package app

import (
	"github.com/cloustone/pandas/cmd/mixer/app/options"
	"github.com/cloustone/pandas/mixer"
	"github.com/cloustone/pandas/mixer/grpc_mixer_v1"
	"github.com/cloustone/pandas/models/factory"
	_ "github.com/cloustone/pandas/pkg/readers/grpc"
	"github.com/cloustone/pandas/pkg/server"
	"github.com/gogo/protobuf/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ManagementServer struct {
	mixer.MixerService
	server.GenericGrpcServer
}

func NewManagementServer(runOptions *options.ServerRunOptions) *ManagementServer {
	s := &ManagementServer{}
	s.RegisterService = func() {
		grpc_mixer_v1.RegisterMixerServer(s.Server, s)
	}
	return s

}

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use:  "mixer",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}

// Run runs the specified APIServer.  This should never exit.
func Run(runOptions *options.ServerRunOptions, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	logrus.Infof("Version: %+v", version.Get())

	// Initialize object factory
	factory.Initialize(runOptions.ModelServing)

	NewManagementServer(runOptions).Run(runOptions.SecureServing)
	<-stopCh
	return nil
}
