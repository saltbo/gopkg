package testutil

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

type mockServer struct {
	*httptest.Server

	req      *http.Request
	duration time.Duration
	respBody []byte
}

func NewMockServer() *mockServer {
	ms := new(mockServer)
	ms.Server = httptest.NewServer(http.HandlerFunc(ms.handler))
	ms.respBody = []byte("mock succ!")
	return ms
}

func (m *mockServer) SetDuration(d time.Duration) {
	m.duration = d
}

func (m *mockServer) SetResponseBody(body []byte) {
	m.respBody = body
}

func (m *mockServer) URL() *url.URL {
	u, _ := url.Parse(m.Server.URL)
	return u
}

func (m *mockServer) handler(w http.ResponseWriter, r *http.Request) {
	m.req = r
	if m.duration > 0 {
		time.Sleep(m.duration)
	}

	for _, cookie := range r.Cookies() {
		w.Header().Add("Set-Cookie", cookie.String())
	}

	_, err := w.Write(m.respBody)
	if err != nil {
		log.Printf("[mockServer] response failed: %s", err)
		return
	}
}
