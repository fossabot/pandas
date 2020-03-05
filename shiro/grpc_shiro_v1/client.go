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
package grpc_shiro_v1

import (
	"fmt"

	serveroptions "github.com/cloustone/pandas/pkg/server/options"
	logr "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	DefaultPort = 30014
	ServiceName = "shiro"
)

var (
	// LbsV1LocalAddress exposes the local address that is used if we run with DNS disabled
	V1LocalAddress = fmt.Sprintf(":%d", DefaultPort)
)

// Client is a client interface for interacting with the tuning service.
type Client interface {
	UserManager() UnifiedUserManagementClient
	Close() error
}

type client struct {
	userManager UnifiedUserManagementClient
	conn        *grpc.ClientConn
}

// NewClient create a new load-balanced client to talk to the Tuning
// service. If the dns_server config option is set to 'disabled', it will
// default to the pre-defined LocalPort of the service.
func NewClient() (Client, error) {
	return NewWithAddress(V1LocalAddress)
}

// NewWithAddress create a new load-balanced client to talk to the Tuning
// service. If the dns_server config option is set to 'disabled', it will
// default to the pre-defined LocalPort of the service.
func NewWithAddress(addr string) (Client, error) {
	address := serveroptions.GetServiceAddress(addr, ServiceName)
	caKey := serveroptions.GetCAKey()

	logr.Infof("NewWithAddress: address: %s, ca: %s", address, caKey)

	var dialOpts []grpc.DialOption

	if serveroptions.IsTLSEnabled() {
		creds, err := credentials.NewClientTLSFromFile(caKey, ServiceName)
		if err != nil {
			return nil, err
		}
		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithBlock()}
	} else {
		dialOpts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	}

	conn, err := grpc.Dial(address, dialOpts...)
	if err != nil {
		return nil, err
	}
	logr.Infof("Dial pms service ok")

	return &client{
		conn:        conn,
		userManager: NewUnifiedUserManagementClient(conn),
	}, nil
}

func (c *client) UserManager() UnifiedUserManagementClient {
	return c.userManager
}

func (c *client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
