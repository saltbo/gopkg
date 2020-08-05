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
	rs, err := NewServer(":8057")
	assert.NoError(t, err)

	rs.SetupPing()
	rs.SetupResource("/api", &mockResource{})

	time.AfterFunc(time.Second, rs.Stop)
	assert.NoError(t, rs.Run())
}
