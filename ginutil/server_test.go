package ginutil

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

type mockResource struct {
}

func (m *mockResource) Register(router *gin.RouterGroup) {
	router.GET("/mock", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
}

func TestServer(t *testing.T) {
	rs := NewServer(":8057")
	rs.SetupGroupRS("/api", &mockResource{})
	rs.SetupPing()

	time.AfterFunc(time.Second, rs.Stop)
	assert.NoError(t, rs.Run())
}

func TestNewIndex(t *testing.T) {
	os.Mkdir("/tmp/gopkg", 0755)
	if !assert.NoError(t, ioutil.WriteFile("/tmp/gopkg/index.html", []byte("test"), 0755)) {
		return
	}

	rs := NewServer(":8057")
	rs.SetupIndex("/", NewIndex("/tmp/gopkg", func(c *gin.Context) {
		c.Header("test", "aaa")
	}))
	time.AfterFunc(time.Second, func() {
		resp, err := resty.New().R().Get("http://localhost:8057/")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		assert.Equal(t, []byte("test"), resp.Body())
		assert.Equal(t, "aaa", resp.Header().Get("test"))
		rs.Stop()
	})
	assert.NoError(t, rs.Run())
}
