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
package rulechain

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNodeChains(t *testing.T) {
	Convey("Construct node chanis", t, func() {
		manifest, err := ParseManifest([]byte(manifestNormalSample))
		So(err, ShouldBeNil)

		chains, err := NewNodeChainsWithManifest(manifest)
		So(err, ShouldBeNil)
		So(len(chains), ShouldEqual, 1)
	})

	Convey("Handling message", t, func() {
		manifest, err := ParseManifest([]byte(manifestNormalSample))
		So(err, ShouldBeNil)

		chains, err := NewNodeChainsWithManifest(manifest)
		So(err, ShouldBeNil)
		So(len(chains), ShouldEqual, 1)

		vals := map[string]interface{}{}
		metadata := newTestingMetadata(vals)
		payload := []byte(`{"hello":"world", "who":"jenson"}`)
		msg := newTestingMessage("M0001", "device", MESSAGE_TYPE_POST_ATTRIBUTES_REQUEST, payload, metadata)

		err = chains[0].ApplyData(msg)
		So(err, ShouldBeNil)
	})
}
