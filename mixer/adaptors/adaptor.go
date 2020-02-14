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

import "github.com/cloustone/pandas/models"

type Adaptor interface {
	Name() string
	Options() *AdaptorOptions
	Start() error
	GracefulShutdown() error
	WithMessageBuilder(MessageBuilder) Adaptor
}

type AdaptorFactory interface {
	Create(*AdaptorOptions) (Adaptor, error)
}

type AdaptorOptions struct {
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	IsProvider   bool   `json:"isProvider"`
	ServicePort  string `json:"servicePort"`
	IsTLSEnabled bool   `json:"isTlsEnabled"`
	ConnectURL   string `json:"connectURL"`
	CertFile     []byte `json:"certFile"`
	KeyFile      []byte `json:"keyFile"`
}

type MessageBuilder interface {
	ConstructMessage(payload []byte) (models.Message, error)
}
