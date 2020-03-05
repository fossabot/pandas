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
	"github.com/cloustone/pandas/apimachinery/restapi/operations/user"
	"github.com/cloustone/pandas/models"
	"github.com/cloustone/pandas/shiro/grpc_shiro_v1"
	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// LoginUser POST /users/login
func LoginUser(params user.LoginUserParams) middleware.Responder {
	client, _ := grpc_shiro_v1.NewClient()
	defer client.Close()

	req := &grpc_shiro_v1.AuthenticateRequest{
		Username: params.Username,
		Password: params.Password,
	}
	resp, err := client.UserManager().Authenticate(params.HTTPRequest.Context(), req)
	if err != nil {
		if grpc.Code(err) == codes.InvalidArgument || grpc.Code(err) == codes.NotFound {
			return user.NewLoginUserBadRequest()
		}
		// return error500("shiro  service call failed", err)
	}

	return user.NewLoginUserOK().WithPayload(
		&models.LoginToken{
			TokenType:   "bearer",
			AccessToken: resp.Token,
		})
}

// LogoutUser GET /users/logout
func LogoutUser(params user.LogoutUserParams, principal *models.Principal) middleware.Responder {
	return user.NewLogoutUserDefault(200)
}
