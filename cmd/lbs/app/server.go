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
	"github.com/cloustone/pandas/cmd/lbs/app/options"
	"github.com/cloustone/pandas/lbs"
	"github.com/cloustone/pandas/lbs/grpc_lbs_v1"
	lbsproxy "github.com/cloustone/pandas/lbs/proxy"
	"github.com/cloustone/pandas/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type LocationBasedServer struct {
	lbs.LbsService
	server.GenericGrpcServer
}

func NewLocationBasedServer(runOptions *options.ServerRunOptions) *LocationBasedServer {
	s := &LocationBasedServer{
		LbsService: lbs.LbsService{
			Proxy: lbsproxy.NewProxy(runOptions.LocationServing),
		},
	}
	s.RegisterService = func() {
		grpc_lbs_v1.RegisterLBSServer(s.Server, s)
	}
	return s
}

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use:  "lbs",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}

// Run runs the specified APIServer.  This should never exit.
func Run(runOptions *options.ServerRunOptions, stopCh <-chan struct{}) error {

	service := NewLocationBasedServer(runOptions)
	service.Run(runOptions.SecureServing)
	<-stopCh
	return nil
}
