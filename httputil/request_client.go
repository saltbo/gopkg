//  Copyright 2020 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package httputil

import (
	"fmt"
	"net/url"
)

type requestClient struct {
	*Resty

	resourcePath string
}

func NewClient(endpoint string, resourcePath string) *requestClient {
	return &requestClient{
		Resty: NewResty(fmt.Sprintf("http://%s", endpoint)),

		resourcePath: resourcePath,
	}
}

func (rc *requestClient) Find(name string, result interface{}) error {
	return rc.Get(fmt.Sprintf("%s/%s", rc.resourcePath, name), result)
}

func (rc *requestClient) FindAll(query url.Values, result interface{}) error {
	rs := rc.resourcePath
	if query != nil {
		rs += "?" + query.Encode()
	}

	return rc.Get(rs, result)
}

func (rc *requestClient) Create(spec interface{}) error {
	return rc.Resty.Create(rc.resourcePath, spec)
}

func (rc *requestClient) Update(spec interface{}) error {
	return rc.Resty.Update(rc.resourcePath, spec)
}

func (rc *requestClient) Delete(name string) error {
	return rc.Resty.Delete(fmt.Sprintf("%s/%s", rc.resourcePath, name))
}
