/*
Copyright 2021 The Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/emicklei/go-restful"
	"github.com/nigel/docker-registry-api/pkg/apis"
	"go.uber.org/zap"
	"knative.dev/pkg/logging"
)

type App struct {
	sync.Once

	Context context.Context

	Logger *zap.SugaredLogger

	container *restful.Container
}

func NewApp() *App {
	return &App{}
}

func (a *App) init() {
	a.Once.Do(func() {
		a.Context = context.Background()

		a.Logger = logging.FromContext(a.Context)

		a.container = restful.NewContainer()
		// support multi level repo, example: project/test/repo:tag, repo name is test/repo
		a.container.Router(restful.RouterJSR311{})
	})
}

func (a *App) Run() error {
	a.init()

	defer func() {
		if a.Logger != nil {
			a.Logger.Sync()
		}
	}()

	ws := a.newWebService()
	a.container.Add(ws)

	r := apis.NewRegistry(a.Logger)
	r.RegisterRoute(a.Context, ws)

	port := 8100
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.container,
	}
	return srv.ListenAndServe()
}

func (a *App) newWebService() *restful.WebService {
	ws := &restful.WebService{}

	ws.Path("/apis")

	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	return ws
}
