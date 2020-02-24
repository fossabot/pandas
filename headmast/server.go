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
	"golang.org/x/sync/errgroup"
	macaron "gopkg.in/macaron.v1"
)

const (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

type HeadmastService struct {
	context       context.Context
	shutdownFn    context.CancelFunc
	childRoutines *errgroup.Group
	macaron       *macaron.Macaron
	httpsrv       *http.Server
	jobManager    JobManager
}

func NewHeadmastService(servingOptions *ServingOptions) *HeadmastService {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)
	s := &HeadmastService{
		context:       childCtx,
		shutdownFn:    shutdownFn,
		childRoutines: childRoutines,
		jobManager:    NewJobManager(servingOptions),
	}

	r := macaron.New()
	r.SetAutoHead(true)
	r.Use(macaron.Renderer())
	r.Post("/api/v1/jobs/", binding.Bind(Job{}), s.createJob)
	r.Delete("/api/v1/jobs/:jobid", s.deleteJob)
	r.Get("/api/v1/jobs/:jobid", s.getJob)
	r.Get("/api/v1/jobs?domain=:domain", s.getJobs)
	r.Get("/api/v1/watch?path=:path", s.watchJobPath)
	r.Post("/api/v1/jobs/:jobid/:action", s.controlJob)

	addr := fmt.Sprintf(":%s", servingOptions.SecureServing.BindPort)
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

func (s *HeadmastService) deleteJob(ctx *macaron.Context)    {}
func (s *HeadmastService) getJob(ctx *macaron.Context)       {}
func (s *HeadmastService) getJobs(ctx *macaron.Context)      {}
func (s *HeadmastService) watchJobPath(ctx *macaron.Context) {}
func (s *HeadmastService) controlJob(ctx *macaron.Context)   {}
