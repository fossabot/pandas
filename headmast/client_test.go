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

package headmast

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateNewJob(t *testing.T) {
	Convey("TestCreateNewJob", t, func() {
		Convey("Should be return error when json.Marshal error", func() {
			monkey.Patch(json.Marshal, func(interface{}) ([]byte, error) {
				return nil, errors.New("marshal panic")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			err := cli.CreateNewJob(job)
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when json unmarshal success", func() {
			monkey.Patch(http.Post, func(string, string, io.Reader) (*http.Response, error) {
				return nil, errors.New("http post failed")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			err := cli.CreateNewJob(job)
			ShouldNotBeNil(err)
		})
	})
}
