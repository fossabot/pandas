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

import (
	"encoding/json"
	"io/ioutil"
)

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
}

func NewRealm(realmOptions *RealmOptions) Realm {
	switch realmOptions.Name {
	case "l2dap":
	default:
		return nil
	}
}

func NewRealmOptionsWithFile(fullFilePath string) []*RealmOptions {
	buf, err := ioutil.ReadAll(fullFilePath)
	if err != nil {
		logrus.WithError(err).Fatalf("open realms config file failed")
	}
	realmOptions := []*RealmOptions{}
	if err := json.Unmarshal(buf, &realmOptions); err != nil {
		logurs.WithError(err).Fatalf("illegal realm config file")
	}
	if len(realmOptions) == 0 {
		logrus.Fatalf("no realms are specified")
	}
	return realmOptions
}
