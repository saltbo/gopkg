package ginutil

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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
