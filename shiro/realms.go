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
	"encoding/json"
	"io/ioutil"

	"github.com/cloustone/pandas/shiro/realms"
	"github.com/cloustone/pandas/shiro/realms/ldap"
	"github.com/sirupsen/logrus"
)

func NewRealm(realmOptions *realms.RealmOptions) realms.Realm {
	switch realmOptions.Name {
	case ldap.AdaptorName:
		realm, err := ldap.NewLdapRealm(realmOptions)
		if err != nil {
			logrus.WithError(err).Fatalf("invalid realm '%s' options", ldap.AdaptorName)
		}
		return realm
	default:
		logrus.Fatalf("invalid realm names")
	}
	return nil
}

func NewRealmOptionsWithFile(fullFilePath string) []*realms.RealmOptions {
	buf, err := ioutil.ReadFile(fullFilePath)
	if err != nil {
		logrus.WithError(err).Fatalf("open realms config file failed")
	}
	realmOptions := []*realms.RealmOptions{}
	if err := json.Unmarshal(buf, &realmOptions); err != nil {
		logrus.WithError(err).Fatalf("illegal realm config file")
	}
	if len(realmOptions) == 0 {
		logrus.Fatalf("no realms are specified")
	}
	return realmOptions
}
