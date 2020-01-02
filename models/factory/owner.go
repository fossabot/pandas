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
package factory

// Owner represent the object owner

type Owner interface {
	WithProject(string) Owner
	WithWorkshop(string) Owner
	User() string
	Project() string
	Workshop() string
}

type objectOwner struct {
	userId     string
	projectId  string
	workshopId string
}

func NewOwner(userId string) Owner {
	return &objectOwner{userId: userId}
}

func (o *objectOwner) WithProject(projectId string) Owner {
	o.projectId = projectId
	return o
}

func (o *objectOwner) WithWorkshop(wid string) Owner {
	o.workshopId = wid
	return o
}

func (o *objectOwner) User() string     { return o.userId }
func (o *objectOwner) Project() string  { return o.projectId }
func (o *objectOwner) Workshop() string { return o.workshopId }
