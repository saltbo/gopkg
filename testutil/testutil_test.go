package testutil

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETRequest(t *testing.T) {
	s := NewMockServer()
	defer s.Close()

	w := TestRequest("GET", "/abc").Do()
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPOSTRequest(t *testing.T) {
	s := NewMockServer()
	defer s.Close()

	body := []byte("test body")
	w := TestRequest("POST", "/abc").Body(body).Do()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, w.Body.Bytes())
}
