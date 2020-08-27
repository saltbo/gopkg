package ginutil

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetOrigin(c *gin.Context) string {
	scheme := "http"
	host := c.Request.Host
	forwardedHost := c.GetHeader("X-Forwarded-Host")
	if forwardedHost != "" {
		host = forwardedHost
	}
	forwardedProto := c.GetHeader("X-Forwarded-Proto")
	if forwardedProto == "https" {
		scheme = forwardedProto
	}

	return fmt.Sprintf("%s://%s", scheme, host)
}
