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
package mixer

import (
	"context"
	"fmt"

	"github.com/cloustone/pandas/mixer/adaptors"
	pb "github.com/cloustone/pandas/mixer/grpc_mixer_v1"
)

// MixerService manage all data source adaptors
type MixerService struct {
	adaptorPool *adaptorPool
}

// NewMixerManagmentService return service instance
func NewMixerManagementService() *MixerService {
	return &MixerService{
		adaptorPool: newAdaptorPool(),
	}
}

// buildAdaptorOptions
func buildAdaptorOptions(userID string, c *pb.AdaptorOptions) *adaptors.AdaptorOptions {
	return &adaptors.AdaptorOptions{
		UserID:       userID,
		Name:         c.Name,
		Protocol:     c.Protocol,
		IsProvider:   c.IsProvider,
		ServicePort:  c.ServicePort,
		ConnectURL:   c.ConnectURL,
		IsTLSEnabled: c.IsTLSEnabled,
		KeyFile:      c.KeyFile,
		CertFile:     c.CertFile,
	}
}

// CreateAdaptor create a new adaptor, and will reuse it if same reader already exist
func (s *MixerService) CreateAdaptor(ctx context.Context, in *pb.CreateAdaptorRequest) (*pb.CreateAdaptorResponse, error) {
	adaptorOptions := buildAdaptorOptions(in.UserID, in.AdaptorOptions)
	adaptor := s.adaptorPool.getAdaptorWithOptions(adaptorOptions)
	if adaptor != nil {
		s.adaptorPool.incAdaptorRef(adaptor.Name())
		return &pb.CreateAdaptorResponse{AdaptorID: adaptor.Name()}, nil
	}
	// If no existed adaptor found, create a new adaptor and save it into pool
	adaptor, err := NewAdaptor(adaptorOptions)
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}
	s.adaptorPool.addAdaptor(in.UserID, adaptor)
	return &pb.CreateAdaptorResponse{AdaptorID: adaptor.Name()}, nil

}

// DeleteAdaptor decrease reference of adaptor, and remove the adaptor if reference count is zero
func (s *MixerService) DeleteAdaptor(ctx context.Context, in *pb.DeleteAdaptorRequest) (*pb.DeleteAdaptorResponse, error) {
	adaptor := s.adaptorPool.getAdaptor(in.UserID, in.AdaptorID)
	if adaptor != nil {
		ref := s.adaptorPool.decAdaptorRef(adaptor.Name())
		if ref <= 0 {
			s.adaptorPool.removeAdaptor(in.UserID, adaptor.Name())
			adaptor.GracefulShutdown()
		}
	}
	return &pb.DeleteAdaptorResponse{}, nil
}

// GetAdaptorFactories return all available adaptor facotory's name
func (s *MixerService) GetAdaptorFactories(ctx context.Context, in *pb.GetAdaptorFactoriesRequest) (*pb.GetAdaptorFactoriesResponse, error) {
	factoryNames := GetFactories()
	return &pb.GetAdaptorFactoriesResponse{AdaptorFactoryNames: factoryNames}, nil
}

// GetAdaptors return user's all adaptors
func (s *MixerService) GetAdaptors(ctx context.Context, in *pb.GetAdaptorsRequest) (*pb.GetAdaptorsResponse, error) {
	return nil, nil
}
