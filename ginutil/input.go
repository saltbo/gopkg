package ginutil

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamInt(c *gin.Context, name string) int {
	v, _ := strconv.Atoi(c.Param(name))
	return v
}

func QueryInt(c *gin.Context, name string) int {
	v, _ := strconv.Atoi(c.Query(name))
	return v
}
