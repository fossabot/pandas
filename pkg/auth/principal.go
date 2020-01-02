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
package auth

type Principal interface {
	UserId() string
	ProjectId() string
}

func NewPrincipal(userId, projectId string) Principal {
	return &defaultPrincipal{
		userId:    userId,
		projectId: projectId,
	}
}

type defaultPrincipal struct {
	userId    string
	projectId string
}

func (p *defaultPrincipal) UserId() string    { return p.userId }
func (p *defaultPrincipal) ProjectId() string { return p.projectId }
func (p *defaultPrincipal) WithUserId(userId string) Principal {
	p.userId = userId
	return p
}

func (p *defaultPrincipal) WithProjectId(projectId string) Principal {
	p.projectId = projectId
	return p
}
