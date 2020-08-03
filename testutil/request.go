package testutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

var defaultMux = &mockServer{}

type testRequest struct {
	req *http.Request
	mux http.Handler

	debug bool
}

func TestRequest(method, path string) *testRequest {
	req, _ := http.NewRequest(method, path, nil)
	return &testRequest{
		req: req,
		mux: defaultMux,
	}
}

func (tr *testRequest) Debug() *testRequest {
	tr.debug = true
	return tr
}

func (tr *testRequest) Mux(mux http.Handler) *testRequest {
	tr.mux = mux
	return tr
}

func (tr *testRequest) Host(host string) *testRequest {
	tr.req.Host = host
	return tr
}

func (tr *testRequest) Body(body []byte) *testRequest {
	return tr.BodyReader(bytes.NewBuffer(body))
}

func (tr *testRequest) BodyReader(br io.Reader) *testRequest {
	tr.req.Body = ioutil.NopCloser(br)
	return tr
}

func (tr *testRequest) Header(k, v string) *testRequest {
	tr.req.Header.Set(k, v)
	return tr
}

func (tr *testRequest) Cookie(cookies ...*http.Cookie) *testRequest {
	for _, cookie := range cookies {
		tr.req.AddCookie(cookie)
	}

	return tr
}

func (tr *testRequest) Do() *httptest.ResponseRecorder {
	if tr.debug {
		b, err := httputil.DumpRequest(tr.req, true)
		if err != nil {
			fmt.Printf("httputil.DumpRequest failed: %s", err)
			return nil
		}

		fmt.Printf("========= Test Request =========\n%s\n", string(b))
	}

	w := httptest.NewRecorder()
	tr.mux.ServeHTTP(w, tr.req)
	return w
}
