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
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/cloustone/pandas/cmd/dmms/app"
	"github.com/cloustone/pandas/cmd/dmms/app/options"
	"github.com/cloustone/pandas/pkg/util/flag"
	"github.com/cloustone/pandas/pkg/util/logs"
	"github.com/cloustone/pandas/pkg/util/wait"
	"github.com/spf13/pflag"
)

// inject by go build
var (
	Version   = "0.0.0"
	BuildTime = "2020-01-13-0802 UTC"
)

func init() {
	fmt.Println("Version:", Version)
	fmt.Println("BuildTime:", BuildTime)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)

	flag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := app.Run(s, wait.NeverStop); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
