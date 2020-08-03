//  Copyright 2020 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Resty struct {
	*resty.Client

	baseURL string
}

func NewResty(baseURL string) *Resty {
	return &Resty{Client: resty.New().SetDebug(true), baseURL: baseURL}
}

func (x *Resty) Get(resource string, v interface{}) error {
	resp, err := x.R().Get(x.baseURL + resource)
	if err := x.formatError(resp, err); err != nil {
		return err
	}

	return json.Unmarshal(resp.Body(), v)
}

func (x *Resty) Create(resource string, body interface{}) error {
	return x.formatError(x.R().SetBody(body).Post(x.baseURL + resource))
}

func (x *Resty) Update(resource string, body interface{}) error {
	return x.formatError(x.R().SetBody(body).Put(x.baseURL + resource))
}

func (x *Resty) Delete(resource string) error {
	return x.formatError(x.R().Delete(x.baseURL + resource))
}

func (x *Resty) formatError(response *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return BindJSONResponse(response.Body()).Error()
	}

	return nil
}
