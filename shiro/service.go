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
package shiro

import (
	"context"

	pb "github.com/cloustone/pandas/shiro/grpc_shiro_v1"
	"github.com/cloustone/pandas/shiro/options"
)

// UnifiedUserManagementService manage user's authentication and authorization
type UnifiedUserManagementService struct {
	servingOptions *options.ServingOptions
}

// UnifiedUserManagementService  return service instance
func NewUnifiedUserManagementService(servingOptions *options.ServingOptions) *UnifiedUserManagementService {
	s := &UnifiedUserManagementService{
		servingOptions: servingOptions,
	}
	return s
}

// Authenticate authenticate the principal in specific domain realm
func (s *UnifiedUserManagementService) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {

	return nil, nil
}

// AddDomainRealm adds specific realm
func (s *UnifiedUserManagementService) AddDomainRealm(ctx context.Context, in *pb.AddDomainRealmRequest) (*pb.AddDomainRealmResponse, error) {
	return nil, nil
}

// GetRolePermissions return role's dynamica route path
func (s *UnifiedUserManagementService) GetRolePermissions(ctx context.Context, in *pb.GetRolePermissionsRequest) (*pb.GetRolePermissionsResponse, error) {
	return nil, nil
}
