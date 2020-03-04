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
package ldap

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap"
)

type Options struct {
	Addr         string
	BindUserName string
	BindPassword string
	SearchDN     string
}

type LdapRealm struct {
	conn   *ldap.Conn
	option Options
}

func NewLdapRealm(options Options) (*LDAPService, error) {
	conn, err := ldap.Dial("tcp", options.Addr)
	if err != nil {
		return nil, err
	}

	err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, err
	}

	err = conn.Bind(options.BindUserName, options.BindPassword)
	if err != nil {
		return nil, err
	}

	return &LdapRealm{conn: conn, options: options}, nil
}

func (l *ldap) Authenticate(principal *realm.Principal) error {
	searchRequest := ldap.NewSearchRequest(
		l.SearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=inetOrgPerson)(mail=%s))", principal.UserName),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Conn.Search(searchRequest)
	if err != nil {
		return err
	}

	if len(sr.Entries) != 1 {
		return fmt.Errorf("User does not exist or too many entries returned")
	}

	userDN := sr.Entries[0].DN
	err = l.Conn.Bind(userDN, principal.Password)
	if err != nil {
		return err
	}

	err = l.Conn.Bind(l.Config.BindUserName, l.Config.BindPassword)
	if err != nil {
		return err
	}

	return nil
}
