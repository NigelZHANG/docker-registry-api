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
	"net"
	"net/http"
	"time"

	"github.com/bluele/gcache"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Registry struct {
	restyClient *resty.Client
	*zap.SugaredLogger
	Cache gcache.Cache
}

func NewRegistry(logger *zap.SugaredLogger) *Registry {
	restyClient := resty.NewWithClient(NewHTTPClient())
	restyClient.SetDisableWarn(true)

	return &Registry{
		restyClient:   restyClient,
		SugaredLogger: logger,
		Cache:         gcache.New(100000).LRU().Build(),
	}
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: GetDefaultTransport(),
		Timeout:   30 * time.Second,
	}
}

func GetDefaultTransport() http.RoundTripper {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
