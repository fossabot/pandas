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

	pb "github.com/cloustone/pandas/mixer/grpc_mixer_v1"
)

type MixerService struct {
}

func NewMixerManagementService() *MixerService {
	return &MixerService{}
}

func (s *MixerService) CreateAgent(ctx context.Context, in *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error) {
	return nil, nil
}
func (s *MixerService) DeleteAgent(ctx context.Context, in *pb.DeleteAgentRequest) (*pb.DeleteAgentResponse, error) {
	return nil, nil
}
