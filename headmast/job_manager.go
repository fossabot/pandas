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

	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

const (
	HEADMAST_JOBS_PATH    = "/headmast/jobs"
	HEADMAST_WORKERS_PATH = "/headmast/workers"
	HEADMAST_KILLER_PATH  = "/headmast/killers"
)

// JobManager is light weight scheduler  service based on  etcd backend
type JobManager interface {
	AddJob(job *Job) error
	RemoveJob(jobID string) error
	KillJob(jobID string) error
	GetJob(jobID string) (*Job, error)
	UpdateJob(job *Job) error
}

// NewJobManager return manager instance
func NewJobManager(servingOptions *ServingOptions) JobManager {
	return newJobManager(servingOptions)
}

// jobManager is default manager implementation
type jobManager struct {
	servingOptions *ServingOptions
}

func newJobManager(servingOptions *ServingOptions) JobManager {
	// connection with etcd at first
	s := &jobManager{servingOptions: servingOptions}
	client := s.newEtcdClient()
	defer client.Close()
	return s
}

// newEtcdClient return client endpoint of etcd
func (manager *jobManager) newEtcdClient() *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{manager.servingOptions.EtcdEndpoints},
		DialTimeout: dialTimeout,
	})
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	return client
}

// AddJob post a new job on etcd
// Job id identifier is managed by client, so we must check wether the
// sample job already exist.
func (manager *jobManager) AddJob(job *Job) error {
	client := manager.newEtcdClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	if _, err := client.Get(ctx, HEADMAST_JOBS_PATH+job.ID); err == nil {
		logrus.Infof("same job '%s' already exist", job.ID)
		return nil
	}
	// Complte job information and marshal it
	job.Status = JOB_STATUS_CREATED
	buf, _ := job.MarshalBinary()

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	if _, err := client.Put(ctx, HEADMAST_JOBS_PATH, string(buf)); err != nil {
		logrus.WithError(err)
		return err
	}
	cancel()
	return nil
}

// RemoveJob remove specific job from job manager
func (manager *jobManager) RemoveJob(jobID string) error {
	return nil
}

func (manager *jobManager) KillJob(jobID string) error {
	job, err := manager.GetJob(jobID)
	if err != nil {
		return err
	}
	job.Status = JOB_STATUS_KILLED
	manager.UpdateJob(job)
	return nil
}

func (manager *jobManager) GetJob(jobID string) (*Job, error) {
	key := HEADMAST_JOBS_PATH + jobID

	client := manager.newEtcdClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Infof("job '%s' no exist", jobID)
		return nil, err
	}
	cancel()

	// get job and update job status
	job := NewJob()
	for _, ev := range resp.Kvs {
		job.UnmarshalBinary([]byte(ev.Value))
		break
	}
	return job, nil
}

func (manager *jobManager) UpdateJob(job *Job) error {
	key := HEADMAST_JOBS_PATH + job.ID

	client := manager.newEtcdClient()
	buf, _ := json.Marshal(&job)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.Put(ctx, key, string(buf))
	if err != nil {
		logrus.Infof("job '%s' no exist", job.ID)
		return err
	}
	cancel()
	return nil
}
