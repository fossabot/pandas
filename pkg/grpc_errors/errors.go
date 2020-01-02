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
package grpcerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	NotExist           = status.Errorf(codes.NotFound, "object no exist")
	InvalidArgument    = status.Errorf(codes.InvalidArgument, "invalid argument")
	AlreadyExist       = status.Errorf(codes.AlreadyExists, "already exist")
	FailedPrecondition = status.Errorf(codes.FailedPrecondition, "failed precondition")
	Internal           = status.Errorf(codes.Internal, "internal error")
	OK                 = status.Errorf(codes.OK, "")
	Unknown            = status.Errorf(codes.Unknown, "unknown error")
)
