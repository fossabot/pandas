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
	"github.com/cloustone/pandas/pkg/broadcast"
	broadcast_util "github.com/cloustone/pandas/pkg/broadcast/util"
	"github.com/sirupsen/logrus"
)

// MixerService manage all data source adaptors
type MixerService struct {
	adaptorPool *adaptorPool
}

// NewMixerManagmentService return service instance
func NewMixerManagementService() *MixerService {
	s := &MixerService{
		adaptorPool: newAdaptorPool(),
	}
	broadcast_util.RegisterObserver(s, MIXER_NOTIFICATION_PATH)
	return s
}

// buildAdaptorOptions
func buildAdaptorOptions(c *pb.AdaptorOptions) *adaptors.AdaptorOptions {
	return &adaptors.AdaptorOptions{
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

// OnBroadcase will be notified when rulechain model object is changed
func (s *MixerService) Onbroadcast(b broadcast.Broadcast, notify broadcast.Notification) {
	notification := Notification{}
	if err := notification.UnmarshalBinary(notify.Param); err != nil {
		logrus.Errorf("unmarshal rulechain notifications '%s' failed", notify.ObjectPath)
		return
	}
	adaptorOptions := notification.AdaptorOptions

	switch notify.Action {
	case broadcast.OBJECT_CREATED:
		adaptor := s.adaptorPool.getAdaptorWithOptions(adaptorOptions)
		if adaptor != nil {
			s.adaptorPool.incAdaptorRef(adaptor)
			return
		}
		// If no existed adaptor found, create a new adaptor and save it into pool
		adaptor, err := NewAdaptor(adaptorOptions)
		if err != nil {
			logrus.Errorf("adding mixer adaptor('%s') failed", adaptorOptions.Domain)
			return
		}
		s.adaptorPool.addAdaptor(adaptor)

	case broadcast.OBJECT_UPDATED:
	}
}

// CreateAdaptor create a new adaptor, and will reuse it if same reader already exist
func (s *MixerService) CreateAdaptor(ctx context.Context, in *pb.CreateAdaptorRequest) (*pb.CreateAdaptorResponse, error) {
	adaptorOptions := buildAdaptorOptions(in.AdaptorOptions)
	adaptor := s.adaptorPool.getAdaptorWithOptions(adaptorOptions)
	if adaptor != nil {
		s.adaptorPool.incAdaptorRef(adaptor)
		return &pb.CreateAdaptorResponse{AdaptorID: adaptor.Name()}, nil
	}
	// If no existed adaptor found, create a new adaptor and save it into pool
	adaptor, err := NewAdaptor(adaptorOptions)
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}
	s.adaptorPool.addAdaptor(adaptor)
	return &pb.CreateAdaptorResponse{AdaptorID: adaptor.Name()}, nil

}

// DeleteAdaptor decrease reference of adaptor, and remove the adaptor if reference count is zero
func (s *MixerService) DeleteAdaptor(ctx context.Context, in *pb.DeleteAdaptorRequest) (*pb.DeleteAdaptorResponse, error) {
	adaptorID := BuildAdaptorID(in.Domain, in.Protocol)
	adaptor := s.adaptorPool.getAdaptor(adaptorID)
	if adaptor != nil {
		ref := s.adaptorPool.decAdaptorRef(adaptor)
		if ref <= 0 {
			s.adaptorPool.removeAdaptor(adaptor)
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

// GetAdaptors return domain's all adaptors
func (s *MixerService) GetAdaptors(ctx context.Context, in *pb.GetAdaptorsRequest) (*pb.GetAdaptorsResponse, error) {
	adaptors := s.adaptorPool.getAdaptors(in.Domain)
	allAdaptorOptions := []*pb.AdaptorOptions{}

	for _, adaptor := range adaptors {
		options := adaptor.Options()
		adaptorOptions := &pb.AdaptorOptions{
			Domain:       options.Domain,
			Protocol:     options.Protocol,
			Name:         options.Name,
			IsProvider:   options.IsProvider,
			ServicePort:  options.ServicePort,
			ConnectURL:   options.ConnectURL,
			IsTLSEnabled: options.IsTLSEnabled,
			KeyFile:      options.KeyFile,
			CertFile:     options.CertFile,
		}
		allAdaptorOptions = append(allAdaptorOptions, adaptorOptions)
	}
	return &pb.GetAdaptorsResponse{AdaptorOptions: allAdaptorOptions}, nil
}
