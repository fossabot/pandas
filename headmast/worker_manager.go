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

const (
	HEADMAST_WORKER_PATH = "/headmast/workers"
)

const (
	HEADMAST_CHANGES_ADDED   = "added"
	HEADMAST_CHANGES_DELETED = "deleted"
)

// Worker represent worker node on which the job is run, worker monitor its
// working path and recive job that assigned to it. the worker ID is a UUID
// that created by worker, when worker watch it working path, it will post the
// worker ID to server.
type Worker struct {
	ID          string   `json:"id"`
	WorkingJobs []string `json:"workingJobs"`
	KillingJobs []string `json:"killingJobs"`
}

func (w Worker) WorkingPath() string                             { return fmt.Sprintf("/headmast/workers/%s/jobs", w.ID) }
func (w Worker) KillerPath() string                              { return fmt.Sprintf("/headmast/workers/%s/killers", w.ID) }
func (w *Worker) RetrieveJobs(jobCh chan *Job, errCh chan error) {}

type WorkersObserver func(worker *Worker, reason string)

// WorkerManager monitors worker path on etcd and jobs to workers
// WorkerManager also manager all nodes info, if none node exist, the worker's job
// will be assign to other works
type WorkerManager interface {
	GetWorker(wid string) (*Worker, error)
	UpdateWorkers([]*Worker)
	GetWorkers() []*Worker
	RemoveWorker(wid string)
	RegisterObserver(WorkersObserver)
}

// NewWorkerManager return worker manager instance, one service should only have one
// schduler
func NewWorkerManager(servingOptions *ServingOptions) WorkerManager {
	return newWorkerManager(servingOptions)
}

// workerManager is default worker manager implementation
type workerManager struct {
	servingOptions  *ServingOptions
	workersObserver WorkersObserver
}

// newWorkerManager return default worker manager instance
func newWorkerManager(servingOptions *ServingOptions) WorkerManager {
	manager := &workerManager{
		servingOptions: servingOptions,
	}
	go manager.watchWorkerChanged()
	return manager
}

// watchWorkerChanged will monitor etcd '/headmast/workers' for worker's change
func (manager *workerManager) watchWorkerChanged() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{manager.servingOptions.EtcdEndpoints},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		logrus.Fatalf("connect failed, err:", err)
		return
	}
	logrus.Println("worker manage connect with etcd '/headdmast/workers' successfully")
	defer cli.Close()

	for true {
		rch := cli.Watch(context.Background(), HEADMAST_WORKER_PATH)
		for resp := range rch {
			for _, ev := range resp.Events {
				logrus.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				worker := Worker{}
				if err := json.Unmarshal(ev.Kv.Value, &worker); err != nil {
					logrus.WithError(err)
					break
				}
				reason := HEADMAST_CHANGES_ADDED
				if ev.Type == clientv3.EventTypeDelete {
					reason = HEADMAST_CHANGES_DELETED
				}
				manager.workersObserver(&worker, reason)
			}
		}
	}
}

// UpdateWorker update worker with specific jobs
func (manager *workerManager) UpdateWorkers(workers []*Worker) {
}

// GetWorker return specific worker node on etcd path '/headmast/workers/%s`
func (manager *workerManager) GetWorker(wid string) (*Worker, error) {
	return nil, nil
}

// GetWorker return all availables worker nodes in etcd path
// '/headmast/workers/`
func (manager *workerManager) GetWorkers() []*Worker {
	return nil
}
func (manager *workerManager) RemoveWorker(wid string) {
}

func (manager *workerManager) RegisterObserver(observer WorkersObserver) {
	manager.workersObserver = observer
}
