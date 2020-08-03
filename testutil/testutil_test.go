package testutil

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGETRequest(t *testing.T) {
	w := TestRequest("GET", "/abc").Debug().Do()
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGETRequestWithMockServer(t *testing.T) {
	s := NewMockServer()
	defer s.Close()

	w := TestRequest("GET", s.URL().String()).Mux(s).Do()
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPOSTRequest(t *testing.T) {
	host := "test.localhost"
	body := []byte("test body")
	cookies := []*http.Cookie{
		{Name: "abc", Value: "123"},
		{Name: "bcd", Value: "456"},
	}

	s := NewMockServer()
	s.SetHost(host)
	s.SetDuration(time.Second)
	s.SetResponseBody(body)
	defer s.Close()

	tr := TestRequest("POST", "/abc")
	w := tr.Mux(s).Host(host).Header("X-T", "1").Cookie(cookies...).Body(body).Do()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, w.Body.Bytes())
}

func TestErrorHost(t *testing.T) {
	host := "test.localhost"
	s := NewMockServer()
	s.SetHost(host)
	s.SetDuration(time.Second)
	defer s.Close()

	w := TestRequest("GET", "/abc").Mux(s).Do()
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}
