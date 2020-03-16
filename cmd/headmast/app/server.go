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
	"github.com/cloustone/pandas/cmd/headmast/app/options"
	"github.com/cloustone/pandas/headmast"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ManagementServer struct {
	headmast.HeadmastService
}

func NewManagementServer(runOptions *options.ServerRunOptions) *ManagementServer {
	s := &ManagementServer{
		HeadmastService: *headmast.NewHeadmastService(runOptions.HeadmastServingOptions),
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

	NewManagementServer(runOptions).Run()
	<-stopCh
	return nil
}
