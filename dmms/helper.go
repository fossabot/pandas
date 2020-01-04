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
package dmms

import (
	"errors"

	"github.com/cloustone/pandas/models/factory"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcError return grpc error according to models errors
func grpcError(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, factory.ErrObjectNotFound) {
		return status.Errorf(codes.NotFound, "%w", err)
	} else if errors.Is(err, factory.ErrObjectAlreadyExist) {
		return status.Errorf(codes.AlreadyExists, "%w", err)
	} else if errors.Is(err, factory.ErrObjectInvalidArg) {
		return status.Errorf(codes.InvalidArgument, "%w", err)
	} else {
		return status.Errorf(codes.Internal, "%s", err)
	}
}
