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
	"fmt"
	"net/http"
	"time"

	"github.com/go-macaron/binding"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	macaron "gopkg.in/macaron.v1"
)

const (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

type HeadmastService struct {
	servingOptions *ServingOptions
	context        context.Context
	shutdownFn     context.CancelFunc
	childRoutines  *errgroup.Group
	macaron        *macaron.Macaron
	httpsrv        *http.Server
	jobManager     JobManager
	workerManager  WorkerManager
	jobScheduler   JobScheduler
}

// NewHeadmastService manage http rest api server to handle client's request
// JobManager and JobScheduler is backend for job's control
func NewHeadmastService(servingOptions *ServingOptions) *HeadmastService {
	jobManager := NewJobManager(servingOptions)
	workerManager := NewWorkerManager(servingOptions)
	jobScheduler := NewJobScheduler(servingOptions, jobManager, workerManager)

	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)
	s := &HeadmastService{
		servingOptions: servingOptions,
		context:        childCtx,
		shutdownFn:     shutdownFn,
		childRoutines:  childRoutines,
		jobManager:     jobManager,
		workerManager:  workerManager,
		jobScheduler:   jobScheduler,
	}

	r := macaron.New()
	r.SetAutoHead(true)
	r.Use(macaron.Renderer())
	r.Post("/api/v1/jobs/", binding.Bind(Job{}), s.createJob)
	r.Delete("/api/v1/jobs/:jobid", s.deleteJob)
	r.Get("/api/v1/jobs/:jobid", s.getJob)
	r.Get("/api/v1/jobs", s.getJobs)
	r.Get("/api/v1/watch/:jobid/", s.watchJobPath)
	r.Post("/api/v1/jobs/:jobid/:action", s.controlJob)

	addr := fmt.Sprintf(":%d", servingOptions.SecureServing.BindPort)
	s.httpsrv = &http.Server{Addr: addr, Handler: r}
	s.macaron = r
	return s
}

func (s *HeadmastService) Run() error {
	s.childRoutines.Go(func() error {
		return s.httpsrv.ListenAndServe()
	})
	return nil
}

// createJob add a new job to etcd
// The new job will be uploaded to /jobs/
func (s *HeadmastService) createJob(ctx *macaron.Context, job Job) {
	s.jobManager.AddJob(&job)
	ctx.JSON(200, nil)
}

// deleteJob delete specified job from headmast
func (s *HeadmastService) deleteJob(ctx *macaron.Context) {
	s.jobManager.RemoveJob(ctx.Query("jobid"))
	ctx.JSON(200, nil)
}

// getJob return specific job detail
func (s *HeadmastService) getJob(ctx *macaron.Context) {
	jobID := ctx.Query("jobid")
	job, err := s.jobManager.GetJob(jobID)
	if err != nil {
		logrus.WithError(err)
		ctx.JSON(500, nil)
		return
	}
	ctx.JSON(200, job)
}

// getJobs return all jobs in headamast
func (s *HeadmastService) getJobs(ctx *macaron.Context) {
	jobs := s.jobManager.GetJobs()
	ctx.JSON(200, jobs)
}

// watchJob is used by headmast client to monitor client's job path
func (s *HeadmastService) watchJobPath(ctx *macaron.Context) {
	worker, err := s.workerManager.GetWorker(ctx.Query("workerid"))
	if err != nil {
		logrus.WithError(err)
		ctx.JSON(500, err)
		return
	}
	errCh := make(chan error, 1)
	jobCh := make(chan *Job, 1)
	defer close(errCh)
	defer close(jobCh)

	go worker.RetrieveJobs(jobCh, errCh)
	for {
		select {
		case job := <-jobCh:
			ctx.JSON(200, job)
		case err := <-errCh:
			ctx.JSON(504, err)
			return
		}
	}
	ctx.JSON(200, nil)
}

// controlJob
func (s *HeadmastService) controlJob(ctx *macaron.Context) {
	ctx.Query("jobid")
	ctx.Query("action")
}
