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
	"errors"
	"net/http"
	"testing"

	"bou.ke/monkey"
	"github.com/cloustone/pandas/pkg/server/options"
	. "github.com/smartystreets/goconvey/convey"
	macaron "gopkg.in/macaron.v1"
)

func TestRun(t *testing.T) {
	Convey("TestRun", t, func() {
		Convey("Should be return error when Run accourd error", func() {
			monkey.Patch(http.ListenAndServe, func(string, http.Handler) error {
				return errors.New("tcp connected failed")
			})
			defer monkey.UnpatchAll()
			certkey := options.CertKey{
				CertFile: "certfile",
				KeyFile:  "keyfile",
			}
			generatablekeycert := options.GeneratableKeyCert{
				CertKey:       certkey,
				CACertFile:    "cacertfile",
				CertDirectory: "certdirectory",
				PairName:      "pariname",
			}
			secureservingoptions := options.SecureServingOptions{
				BindAddress: []byte("10.4.47.48"),
				BindPort:    8080,
				ServerCert:  generatablekeycert,
				// useLoopbackCfg: true,
			}
			servingoptions := ServingOptions{
				SecureServing:  &secureservingoptions,
				EtcdEndpoints:  "etcdendpoints",
				SchedulePolicy: "schedulepoints",
			}
			srv := NewHeadmastService(&servingoptions) // HeadmastService怎么定义？？
			err := srv.Run()
			ShouldNotBeNil(err)
		})
	})
}
func TestcerateJob(t *testing.T) {
	Convey("TestcreateJob", t, func() {
		Convey("Should be return error when New Etcd Client error", func() {
			monkey.Patch(JobManager.AddJob, func(*Job) error {
				return errors.New("New Etcd Client Failed")
			})
			defer monkey.UnpatchAll()
			job := Job{
				ID:      "1234",
				Domain:  "hello",
				Payload: []byte("world"),
				Status:  "unknown",
			}
			servingoptions := ServingOptions{}
			srv := NewHeadmastService(&servingoptions)
			a := &macaron.Context{}
			srv.createJob(a, job)
		})
	})
}
func TestdeleteJob(t *testing.T) {
	Convey("TestdeleteJob", t, func() {
		Convey("Should be return error when delete specified job from headmast error", func() {
			monkey.Patch(JobManager.RemoveJob, func(string) error {
				return errors.New("Delete Job Failed")
			})
			defer monkey.UnpatchAll()
			servingoptions := ServingOptions{}
			srv := NewHeadmastService(&servingoptions)
			a := &macaron.Context{}
			srv.deleteJob(a)
		})
	})
}
func TestgetJob(t *testing.T) {
	Convey("TestgetJob", t, func() {
		Convey("Should be return error when get job error", func() {
			monkey.Patch(JobManager.GetJob, func(string) error {
				return errors.New("Get Job Failed")
			})
			defer monkey.UnpatchAll()
			servingoptions := ServingOptions{}
			srv := NewHeadmastService(&servingoptions)
			a := &macaron.Context{}
			srv.getJob(a)
		})
	})
}
func TestgetJobs(t *testing.T) {
	Convey("TestgetJobs", t, func() {
		Convey("Should be return error when get jobs error", func() {
			monkey.Patch(JobManager.GetJobs, func() {
				return // GetJobs function doesn't have return value
			})
			defer monkey.UnpatchAll()
			servingoptions := ServingOptions{}
			srv := NewHeadmastService(&servingoptions)
			a := &macaron.Context{}
			srv.getJob(a)
		})
	})
}
func TestwatchJobPath(t *testing.T) {
	Convey("TestwatchJobPath", t, func() {
		Convey("Should be return error when mointor client's job path error", func() {
			monkey.Patch(workerManager.GetWorker, func(string) error {
				return errors.New("mointor path Failed")
			})
			defer monkey.UnpatchAll()
			servingoptions := ServingOptions{}
			srv := NewHeadmastService(&servingoptions)
			a := &macaron.Context{}
			srv.watchJobPath(a)
		})
	})
}
func TestcontrolJob(t *testing.T) {
	Convey("TestcontrolJob", t, func() {
		Convey("Should be return error when control job path error", func() {
			monkey.Patch(workerManager.GetWorker, func(string) error {
				return errors.New("Control job Failed")
			})
			defer monkey.UnpatchAll()
			servingoptions := ServingOptions{}
			srv := NewHeadmastService(&servingoptions)
			a := &macaron.Context{}
			srv.watchJobPath(a)
		})
	})
}
