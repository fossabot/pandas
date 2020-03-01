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
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

// WorkerContext hold working context
type WorkerContext struct {
	EtcdEndpoints []string
}

// Worker represent worker node on which the job is run, worker monitor its
// working path and recive job that assigned to it. the worker ID is a UUID
// that created by worker, when worker watch it working path, it will post the
// worker ID to server.
type Worker struct {
	context     WorkerContext `json:"-"`
	ID          string        `json:"id"`
	WorkingJobs []string      `json:"workingJobs"`
	KillingJobs []string      `json:"killingJobs"`
}

func NewWorker(ctx WorkerContext) *Worker {
	return &Worker{context: ctx}
}

func (w Worker) WorkingPath() string { return fmt.Sprintf("/headmast/workers/%s/jobs", w.ID) }
func (w Worker) KillerPath() string  { return fmt.Sprintf("/headmast/workers/%s/killers", w.ID) }
func (w *Worker) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

func (w *Worker) UnmarshalBinary(buf []byte) error {
	return json.Unmarshal(buf, w)
}
func (w *Worker) RetrieveJobs(jobCh chan *Job, errCh chan error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   w.context.EtcdEndpoints,
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		errCh <- err
		return
	}
	defer cli.Close()

	for {
		rch := cli.Watch(context.Background(), w.WorkingPath())
		for resp := range rch {
			for _, ev := range resp.Events {
				logrus.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				job := &Job{}
				if err := json.Unmarshal(ev.Kv.Value, job); err != nil {
					logrus.WithError(err)
					errCh <- err
					return
				}
				jobCh <- job
			}
		}
	}
}
