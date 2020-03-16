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
	"github.com/cloustone/pandas/cmd/pms/app/options"
	"github.com/cloustone/pandas/models/factory"
	"github.com/cloustone/pandas/pkg/server"
	"github.com/cloustone/pandas/pms"
	"github.com/cloustone/pandas/pms/grpc_pms_v1"
	"github.com/cloustone/pandas/scada"
	"github.com/cloustone/pandas/scada/grpc_scada_v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ProjectManagementServer struct {
	pms.ProjectManagementService
	scada.ScadaService
	server.GenericGrpcServer
}

func NewProjectManagementServer() *ProjectManagementServer {
	s := &ProjectManagementServer{}
	s.RegisterService = func() {
		grpc_pms_v1.RegisterProjectManagementServer(s.Server, s)
		grpc_scada_v1.RegisterScadaServer(s.Server, s)
	}
	return s
}

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use:  "pms",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}

// Run runs the specified APIServer.  This should never exit.
func Run(runOptions *options.ServerRunOptions, stopCh <-chan struct{}) error {

	// Initialize object factory
	factory.Initialize(runOptions.ModelServing)

	NewProjectManagementServer().Run(runOptions.SecureServing)
	<-stopCh
	return nil
}
