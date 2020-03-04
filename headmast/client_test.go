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
	"io/ioutil"
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
func TestDeleteJob(t *testing.T) {
	Convey("TestDeletsJob", t, func() {
		Convey("Should be return error when delete error", func() {
			monkey.Patch(http.NewRequest, func(string, string, io.Reader) (*http.Request, error) {
				return nil, errors.New("http request failed")
			})
			jobID := "1234"
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			err := cli.DeleteJob(jobID)
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when delete job", func() {
			monkey.Patch(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
				return nil, errors.New("read failed")
			})
			jobID := "1234"
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			err := cli.DeleteJob(jobID)
			ShouldNotBeNil(err)
		})
	})
}
func TestGetJob(t *testing.T) {
	Convey("TestGetJob", t, func() {
		Convey("Should be return error when Get job error", func() {
			monkey.Patch(http.NewRequest, func(string, string, io.Reader) (*http.Request, error) {
				return nil, errors.New("http request failed")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJob(job.ID)
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when Get job", func() {
			monkey.Patch(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
				return nil, errors.New("read failed")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJob(job.ID)
			ShouldNotBeNil(err)
		})
	})
}
func TestGetJobsWithDomain(t *testing.T) {
	Convey("TestGetJobsWithDomain", t, func() {
		Convey("Should be return error when Get jobs with domain error", func() {
			monkey.Patch(http.NewRequest, func(string, string, io.Reader) (*http.Request, error) {
				return nil, errors.New("http request failed")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJobsWithDomain(job.Domain)
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when Get jobs with domain", func() {
			monkey.Patch(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
				return nil, errors.New("read failed")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJobsWithDomain(job.Domain)
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when with Json unmarshal", func() {
			monkey.Patch(json.Unmarshal, func([]byte, interface{}) error {
				return errors.New("Json unmarshal failed")
			})
			job := &Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "10.4.47.62"})
			_, err := cli.GetJobsWithDomain(job.Domain)
			ShouldNotBeNil(err)
		})
	})
}
func TestGetJobs(t *testing.T) {
	Convey("TestGetJobs", t, func() {
		Convey("Should be return error when Get jobs error", func() {
			monkey.Patch(http.NewRequest, func(string, string, io.Reader) (*http.Request, error) {
				return nil, errors.New("http request failed")
			})
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJobs()
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when Get jobs", func() {
			monkey.Patch(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
				return nil, errors.New("read failed")
			})
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJobs()
			ShouldNotBeNil(err)
		})
		Convey("Should be return nil when with Json unmarshal", func() {
			monkey.Patch(json.Unmarshal, func([]byte, interface{}) error {
				return errors.New("Json unmarshal failed")
			})
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			_, err := cli.GetJobs()
			ShouldNotBeNil(err)
		})
	})
}
func TestWatchPathHandler(t *testing.T) {
	Convey("TestWatchPathHandler", t, func() {
		Convey("Should be return error when watch patch handler", func() {
			monkey.Patch(http.NewRequest, func(string, string, io.Reader) (*http.Request, error) {
				return nil, errors.New("http request faild")
			})
			l := "liyanpeng"
			p := []byte(l)
			job := WatchPathHandler(func(l, p))
			defer monkey.UnpatchAll()
			cli := NewClient(&ClientOptions{ServerAddr: "localhost"})
			err := cli.WatchJobPath(job.jobpath, job)
			ShouldNotBeNil(err)
		})
	})
}
