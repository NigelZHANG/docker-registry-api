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

package apis

import (
	"context"
	"net/http"

	"github.com/emicklei/go-restful"
)

func (r *Registry) RegisterRoute(ctx context.Context, ws *restful.WebService) {
	ws.Route(
		ws.GET("/healthz").To(healthz),
	)

	// TODO:
	ws.Route(
		ws.GET("/projects").To(r.ListProjects).Returns(http.StatusOK, "OK", nil),
	)
}

func healthz(req *restful.Request, resp *restful.Response) {
	resp.WriteHeaderAndJson(http.StatusOK, map[string]string{"ok": "true"}, restful.MIME_JSON)
}
