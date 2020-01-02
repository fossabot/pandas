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
package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

func GetDNS() string {
	if flag := pflag.Lookup("dns"); flag != nil {
		return flag.Value.String()
	}
	return "disabled"
}

func IsTLSEnabled() bool {
	flag := pflag.Lookup("tls-ca-file")
	return flag != nil
}

func GetCAKey() string {
	if flag := pflag.Lookup("tls-ca-key"); flag != nil {
		return flag.Value.String()
	}
	return ""
}

func GetPodNamespace() string {
	if flag := pflag.Lookup("pod.namespace"); flag != nil {
		return flag.Value.String()
	}
	return "default"
}

func GetServiceAddress(addr string, serviceName string) string {
	address := fmt.Sprintf("pandas.%s.%s.svc.cluster.local:80", serviceName, GetPodNamespace())

	dnsServer := GetDNS()
	if dnsServer == "disabled" { // for local testing without DNS server
		address = addr
	} else if dnsServer == "docker-compose" {
		address = fmt.Sprintf("%s:%s", serviceName, addr)
	}
	return address
}
