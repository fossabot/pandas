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
package mixer

import (
	"fmt"

	"github.com/cloustone/pandas/mixer/adaptors"
	"github.com/cloustone/pandas/pkg/broadcast"
	broadcast_util "github.com/cloustone/pandas/pkg/broadcast/util"
)

const (
	MIXER_NOTIFICATION_PATH = "mixer/adaptors"
)

// BuildAdaptorID create adaptor id with domain and protocol
// One domain has only one adaptor for a protocol
func BuildAdaptorID(domain string, protocol string) string {
	return fmt.Sprintf("mixer-%s-%s", domain, protocol)
}

// AsyncCreatorAdaptor notify mixer that a adaptor should be created
func AsyncCreateAdaptor(adaptorOptions *adaptors.AdaptorOptions) {
	broadcast_util.Notify(MIXER_NOTIFICATION_PATH, broadcast.OBJECT_CREATED,
		&Notification{
			Domain:         adaptorOptions.Domain,
			Protocol:       adaptorOptions.Protocol,
			AdaptorOptions: adaptorOptions,
		})
}

// AsyncDeleteAdaptor notify mixer that a adaptor should be deleted
func AsyncDeleteAdaptor(domain string, protocol string) {
	broadcast_util.Notify(MIXER_NOTIFICATION_PATH,
		broadcast.OBJECT_DELETED,
		&Notification{Domain: domain, Protocol: protocol},
	)
}
