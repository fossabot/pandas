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

import "fmt"

const (
	HEADERMAST_WORKER_PATH = "/headmast/workers"
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

func (w Worker) WorkingPath() string { return fmt.Sprintf("/headmast/workers/%s", w.ID) }
func (w Worker) KillerPath() string  { return fmt.Sprintf("/headmast/workers/%s/killers", w.ID) }

const (
	HEADMAST_CHANGES_ADDED   = "added"
	HEADMAST_CHANGES_DELETED = "deleted"
)

type WorkersObserver func(worker *Worker, reason string)

// WorkerManager monitors worker path on etcd and jobs to workers
// WorkerManager also manager all nodes info, if none node exist, the worker's job
// will be assign to other works
type WorkerManager interface {
	GetWorker(wid string) *Worker
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
	return &workerManager{
		servingOptions: servingOptions,
	}
}

// UpdateWorker update worker with specific jobs
func (manager *workerManager) UpdateWorkers(workers []*Worker) {
}

// GetWorker return specific worker node on etcd path '/headmast/workers/%s`
func (manager *workerManager) GetWorker(wid string) *Worker {
	return nil
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
