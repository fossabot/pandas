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
	"github.com/cloustone/pandas/pkg/server"
	"github.com/cloustone/pandas/rulechain"
	"github.com/cloustone/pandas/rulechain/grpc_rulechain_v1"
	"github.com/cloustone/pandas/rulechain/options"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ManagementServer struct {
	rulechain.RuleChainService
	server.GenericGrpcServer
}

func NewManagementServer(servingOptions *options.ServingOptions) *ManagementServer {
	s := &ManagementServer{
		RuleChainService: *rulechain.NewRuleChainService(servingOptions),
	}
	s.RegisterService = func() {
		if servingOptions.IsStandalone() {
			grpc_rulechain_v1.RegisterRuleChainServiceServer(s.Server, s)
		}
	}
	return s
}

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewServingOptions()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use:  "rulechain",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}

// Run runs the specified APIServer.  This should never exit.
func Run(servingOptions *options.ServingOptions, stopCh <-chan struct{}) error {

	NewManagementServer(servingOptions).Run(servingOptions.SecureServing)
	<-stopCh
	return nil
}
