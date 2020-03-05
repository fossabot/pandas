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
package adaptors

import (
	"fmt"
)

type AdaptorObserver interface {
	OnAdaptorMessageAvailable(Adaptor, []byte)
	OnAdaptorError(Adaptor)
}

type Adaptor interface {
	Options() *AdaptorOptions
	Start() error
	GracefulShutdown() error
	RegisterObserver(AdaptorObserver)
}

type AdaptorFactory interface {
	Create(*AdaptorOptions) (Adaptor, error)
}

type AdaptorOptions struct {
	Domain       string   `json:"domain"`
	Protocol     string   `json:"protocol"`
	Name         string   `json:"name"`
	IsProvider   bool     `json:"isProvider"`
	ServicePort  string   `json:"servicePort"`
	IsTLSEnabled bool     `json:"isTlsEnabled"`
	ConnectURL   string   `json:"connectURL"`
	CertFile     []byte   `json:"certFile"`
	KeyFile      []byte   `json:"keyFile"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Endpoints    []string `json:"endpoints"`
}

func NewAdaptorOptions() *AdaptorOptions {
	return &AdaptorOptions{
		IsProvider:   false,
		IsTLSEnabled: false,
		Endpoints:    []string{},
	}
}

// BuildAdaptorID create adaptor id with domain and protocol
// One domain has only one adaptor for a protocol
func BuildAdaptorID(domain string, protocol string) string {
	return fmt.Sprintf("mixer-%s-%s", domain, protocol)
}
