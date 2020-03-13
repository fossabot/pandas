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
package server

import (
	"fmt"
	"net"

	serverOptions "github.com/cloustone/pandas/pkg/server/options"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// LifecycleHandler provides basic lifecycle methods that each microservice has
// to implement.
type LifecycleHandler interface {
	Start(*serverOptions.SecureServingOptions)
	Stop()
	GetListenerAddress() string
}

// Lifecycle implements the lifecycle operations for microservice including
// dynamic service registration.
type GenericGrpcServer struct {
	Listener        net.Listener
	Server          *grpc.Server
	RegisterService func()
}

// Start will start a gRPC microservice on a given port and run it either in
// foreground or background.
func (s *GenericGrpcServer) Run(servingOptions *serverOptions.SecureServingOptions) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servingOptions.BindPort))
	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to listening: %v", err)
	}
	log.Infof("starting service at %s", lis.Addr())
	var opts []grpc.ServerOption
	if servingOptions.IsTlsEnabled() {
		creds, err := credentials.NewServerTLSFromFile(
			servingOptions.ServerCert.CertKey.CertFile,
			servingOptions.ServerCert.CertKey.KeyFile,
		)
		if err != nil {
			log.Fatalf("failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	s.Listener = lis
	s.Server = grpc.NewServer(opts...)

	s.RegisterService()
	grpc_health_v1.RegisterHealthServer(s.Server, health.NewServer())
	go s.Server.Serve(lis)
}

// Stop will stop the gRPC microservice and the socket.
func (s *GenericGrpcServer) Stop() {
	if s.Listener != nil {
		log.Infof("Stopping service at %s", s.Listener.Addr())
	}
	if s.Server != nil {
		s.Server.GracefulStop()
	}
}

// GetListenerAddress will get the address and port the service is listening.
// Returns the empty string if the service is not running but the method is invoked.
func (s *GenericGrpcServer) GetListenerAddress() string {
	if s.Listener != nil {
		return s.Listener.Addr().String()
	}
	return ""
}
