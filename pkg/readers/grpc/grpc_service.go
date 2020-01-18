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

	"github.com/cloustone/pandas/models"
	logr "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type ReaderFactory struct{}

func (r ReaderFactory) Create(configs map[string]interface{}) (models.Reader, error) {
	return newGrpcReader(configs)
}

type grpcReader struct {
	context       context.Context
	shutdownFn    context.CancelFunc
	childRoutines *errgroup.Group
	grpcServer    *grpc.Server
	configs       map[string]interface{}
	observer      models.ReaderObserver
}

func newGrpcReader(configs map[string]interface{}) (models.Reader, error) {
	/*
		var opts []grpc.ServerOption

		if config.IsTLSEnabled() {
			config.FatalOnAbsentKey(config.ServerCertKey)
			config.FatalOnAbsentKey(config.ServerPrivateKey)

			creds, err := credentials.NewServerTLSFromFile(config.GetServerCert(), config.GetServerPrivateKey())
			if err != nil {
				logr.Fatalf("Failed to generate credentials %v", err)
				return nil, err
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}

		rootCtx, shutdownFn := context.WithCancel(context.Background())
		childRoutines, childCtx := errgroup.WithContext(rootCtx)

		return &grpcReader{
			context:       childCtx,
			shutdownFn:    shutdownFn,
			childRoutines: childRoutines,
			grpcServer:    grpc.NewServer(opts...),
			configs:       configs,
		}, nil
	*/
	return nil, nil
}

func (r *grpcReader) Config() map[string]interface{} { return r.configs }

func (g *grpcReader) Start() error {
	grpc_health_v1.RegisterHealthServer(g.grpcServer, health.NewServer())
	RegisterReaderServer(g.grpcServer, g)

	g.childRoutines.Go(func() error {
		port := "4001" //TODO:
		listen, err := net.Listen("tcp", port)
		if err != nil {
			logr.Fatal(err)
		}
		logr.Infof("rpc service is listening on '%s'...", port)
		if err := g.grpcServer.Serve(listen); err != nil {
			logr.Fatal(err)
		}
		return nil
	})
	return nil
}

func (r *grpcReader) Name() string { return "grpc-reader" }

func gerrf(err error, c codes.Code, format string, a ...interface{}) error {
	if err != nil && c != codes.OK {
		logr.WithError(err).Errorf(format, a...)
		return status.Errorf(c, format, a...)
	}
	return nil
}

func (r *grpcReader) GracefulShutdown() error {
	r.grpcServer.Stop()
	r.shutdownFn()

	err := r.childRoutines.Wait()
	if err != nil && err != context.Canceled {
		logr.WithError(err).Errorf("rpc endpoint hutdown failed")
	}
	return nil
}

func (r *grpcReader) PostMessage(ctx context.Context, in *PostMessageRequest) (*PostMessageResponse, error) {
	if r.observer != nil {
		r.observer.OnDataAvailable(r, in.Payload, nil)
	}
	return &PostMessageResponse{}, nil
}

func (r *grpcReader) RegisterObserver(obs models.ReaderObserver) {
	r.observer = obs
}
