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
		Endpoints: []string{manager.servingOptions.EtcdEndpoints},
		//Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	fmt.Println(manager.servingOptions.EtcdEndpoints)
	if err != nil {
		logrus.WithError(err).Fatal("watch worker changed failed")
		return
	}
	logrus.Println("worker manage connect with etcd '/headdmast/workers' successfully")
	defer cli.Close()

	for {
		rch := cli.Watch(context.Background(), HEADMAST_WORKER_PATH)
		for resp := range rch {
			for _, ev := range resp.Events {
				logrus.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				worker := &Worker{
					context: WorkerContext{EtcdEndpoints: []string{manager.servingOptions.EtcdEndpoints}},
				}
				if err := json.Unmarshal(ev.Kv.Value, &worker); err != nil {
					logrus.WithError(err)
					break
				}
				reason := HEADMAST_CHANGES_ADDED
				if ev.Type == clientv3.EventTypeDelete {
					reason = HEADMAST_CHANGES_DELETED
				}
				manager.workersObserver(worker, reason)
			}
		}
	}
}

// newEtcdClient return client endpoint of etcd
func (manager *workerManager) newEtcdClient() *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{manager.servingOptions.EtcdEndpoints},
		DialTimeout: dialTimeout,
	})
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	return client
}

// UpdateWorker update worker with specific jobs
func (manager *workerManager) UpdateWorkers(workers []*Worker) {
	client := manager.newEtcdClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	for _, worker := range workers {
		buf, _ := worker.MarshalBinary()
		if _, err := client.Put(ctx, HEADMAST_WORKER_PATH+worker.ID, string(buf)); err != nil {
			logrus.WithError(err)
		}
	}
	cancel()
}

// GetWorker return specific worker node on etcd path '/headmast/workers/%s`
func (manager *workerManager) GetWorker(wid string) (*Worker, error) {
	client := manager.newEtcdClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := client.Get(ctx, HEADMAST_WORKER_PATH+"/"+wid)
	if err != nil {
		logrus.WithError(err).Errorf("getting worker from manager failed")
		return nil, err
	}
	defer cancel()
	worker := NewWorker(
		WorkerContext{EtcdEndpoints: []string{manager.servingOptions.EtcdEndpoints}})
	for _, ev := range resp.Kvs {
		if err := worker.UnmarshalBinary([]byte(ev.Value)); err != nil {
			logrus.WithError(err).Errorf("can not unmarshal worker object")
			return nil, err
		}
		break
	}

	return worker, nil
}

// GetWorker return all availables worker nodes in etcd path
// '/headmast/workers/`
func (manager *workerManager) GetWorkers() []*Worker {
	workers := []*Worker{}

	client := manager.newEtcdClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := client.Get(ctx, HEADMAST_WORKER_PATH+"/")
	if err != nil {
		logrus.WithError(err).Errorf("getting workers from manager failed")
		return workers
	}
	defer cancel()
	for _, ev := range resp.Kvs {
		worker := NewWorker(
			WorkerContext{EtcdEndpoints: []string{manager.servingOptions.EtcdEndpoints}})
		if err := worker.UnmarshalBinary([]byte(ev.Value)); err != nil {
			logrus.WithError(err).Errorf("can not unmarshal worker object")
			continue
		}
		workers = append(workers, worker)
	}

	return workers

}
func (manager *workerManager) RemoveWorker(wid string) {
	client := manager.newEtcdClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.Delete(ctx, HEADMAST_WORKER_PATH+"/"+wid)
	if err != nil {
		logrus.WithError(err).Errorf("getting workers from manager failed")
		return
	}
	cancel()
}

func (manager *workerManager) RegisterObserver(observer WorkersObserver) {
	manager.workersObserver = observer
}
