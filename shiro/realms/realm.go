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
package realms

type Realm interface {
	Authenticate(principal *Principal) error
}

type RealmOptions struct {
	Name              string `json:"name"`
	CertFile          string `json:"certFile"`
	KeyFile           string `json:"keyFile"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	ServiceConnectURL string `json:"serviceConnectURL"`
	SearchDN          string `json:"searchDN"`
}
