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
package grpc

import (
	"context"
	"net"

	"github.com/cloustone/pandas/rulechain/adaptors"
	logr "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

const Name = "grpc"

type AdaptorFactory struct{}

func (r AdaptorFactory) Create(servingOptions *adaptors.AdaptorOptions) (adaptors.Adaptor, error) {
	return newGrpcAdaptor(servingOptions)
}

type grpcAdaptor struct {
	context         context.Context
	shutdownFn      context.CancelFunc
	childRoutines   *errgroup.Group
	grpcServer      *grpc.Server
	adaptorObserver adaptors.AdaptorObserver
	servingOptions  *adaptors.AdaptorOptions
}

func newGrpcAdaptor(servingOptions *adaptors.AdaptorOptions) (adaptors.Adaptor, error) {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	var opts []grpc.ServerOption
	/*
		if servingOptions.IsTlsEnabled {
			creds, err := credentials.NewServerTLSFromFile(servingOptions.CertFile, servingOptions.KeyFile)
			if err != nil {
				log.Fatalf("failed to generate credentials %v", err)
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}
	*/

	return &grpcAdaptor{
		context:        childCtx,
		shutdownFn:     shutdownFn,
		childRoutines:  childRoutines,
		grpcServer:     grpc.NewServer(opts...),
		servingOptions: servingOptions,
	}, nil
}
func (r *grpcAdaptor) Options() *adaptors.AdaptorOptions {
	return nil
}

func (r *grpcAdaptor) Start() error {
	grpc_health_v1.RegisterHealthServer(r.grpcServer, health.NewServer())
	RegisterAdaptorServer(r.grpcServer, r)

	r.childRoutines.Go(func() error {
		listen, err := net.Listen("tcp", r.servingOptions.ServicePort)
		if err != nil {
			logr.Fatal(err)
		}
		logr.Infof("rpc service is listening on '%s'...", r.servingOptions.ServicePort)
		if err := r.grpcServer.Serve(listen); err != nil {
			logr.Fatal(err)
		}
		return nil
	})
	return nil
}

func (r *grpcAdaptor) Name() string { return Name }

func gerrf(err error, c codes.Code, format string, a ...interface{}) error {
	if err != nil && c != codes.OK {
		logr.WithError(err).Errorf(format, a...)
		return status.Errorf(c, format, a...)
	}
	return nil
}

func (r *grpcAdaptor) GracefulShutdown() error {
	r.grpcServer.Stop()
	r.shutdownFn()

	err := r.childRoutines.Wait()
	if err != nil && err != context.Canceled {
		logr.WithError(err).Errorf("rpc endpoint shutdown failed")
	}
	return nil
}

func (r *grpcAdaptor) PostMessage(ctx context.Context, in *PostMessageRequest) (*PostMessageResponse, error) {
	// broadcast message received from endpoint
	return &PostMessageResponse{}, nil
}

func (r *grpcAdaptor) RegisterObserver(obs adaptors.AdaptorObserver) {
	r.adaptorObserver = obs
}
