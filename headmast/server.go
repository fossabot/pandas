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
	"net/http"

	"github.com/go-macaron/binding"
	"golang.org/x/sync/errgroup"
	macaron "gopkg.in/macaron.v1"
)

type HeadmastService struct {
	context       context.Context
	shutdownFn    context.CancelFunc
	childRoutines *errgroup.Group
	macaron       *macaron.Macaron
	httpsrv       *http.Server
}

func NewHeadmastService() *HeadmastService {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	r := macaron.New()
	r.SetAutoHead(true)
	r.Use(macaron.Renderer())

	r.Post("/api/v1/jobs/", binding.Bind(Job{}), addNewJob)
	r.Delete("/api/v1/jobs/:jobid", removeJob)
	r.Get("/api/v1/jobs/:jobid", getJob)
	r.Get("/api/v1/jobs?domain=:domain", getJobs)
	r.Get("/api/v1/watch?path=:path", watchJobPath)
	r.Post("/api/v1/jobs/:jobid/:action", controlJob)

	addr := "localhost:80"
	httpsrv := &http.Server{Addr: addr, Handler: r}
	return &HeadmastService{
		context:       childCtx,
		shutdownFn:    shutdownFn,
		childRoutines: childRoutines,
		macaron:       r,
		httpsrv:       httpsrv,
	}
}

func (s *HeadmastService) Run() error {
	s.childRoutines.Go(func() error {
		return s.httpsrv.ListenAndServe()
	})
	return nil
}

func addNewJob(ctx *macaron.Context, job Job) { ctx.JSON(200, nil) }
func removeJob(ctx *macaron.Context)          {}
func getJob(ctx *macaron.Context)             {}
func getJobs(ctx *macaron.Context)            {}
func watchJobPath(ctx *macaron.Context)       {}
func controlJob(ctx *macaron.Context)         {}
