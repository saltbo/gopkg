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
	host     string
	duration time.Duration
	respBody []byte
}

func NewMockServer() *mockServer {
	ms := new(mockServer)
	ms.Server = httptest.NewServer(ms)

	ms.host = ms.URL().Host
	ms.respBody = []byte("mock succ!")
	return ms
}

func (m *mockServer) SetHost(host string) {
	m.host = host
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

func (m *mockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host != m.host {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	m.req = r
	if m.duration > 0 {
		time.Sleep(m.duration)
	}

	for _, cookie := range r.Cookies() {
		http.SetCookie(w, cookie)
	}

	if _, err := w.Write(m.respBody); err != nil {
		log.Printf("[mockServer] response failed: %s", err)
		return
	}
}
