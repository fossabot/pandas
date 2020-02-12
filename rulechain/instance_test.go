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
	"io/ioutil"
	"testing"

	"github.com/cloustone/pandas/rulechain/manifest"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNodeChains(t *testing.T) {
	manifestNormalSample, errs := ioutil.ReadFile("../manifest_sample.json")
	So(errs, ShouldBeNil)

	Convey("Construct node chanis", t, func() {
		manifest, err := manifest.New([]byte(manifestNormalSample))
		So(err, ShouldBeNil)

		_, errs := NewWithManifest(manifest)
		So(len(errs), ShouldEqual, 0)
	})
}
