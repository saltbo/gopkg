package testutil

import (
	"io/ioutil"
	"net/http"
)

type testSimpleMux struct {
}

func (mux testSimpleMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		return
	}

	b, _ := ioutil.ReadAll(request.Body)
	_, _ = writer.Write(b)
}
