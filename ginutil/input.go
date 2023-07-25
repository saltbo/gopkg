package ginutil

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamInt(c *gin.Context, name string) int {
	v, _ := strconv.Atoi(c.Param(name))
	return v
}

func ParamInt64(c *gin.Context, name string) int64 {
	v, _ := strconv.ParseInt(c.Param(name), 10, 64)
	return v
}

func QueryInt(c *gin.Context, name string) int {
	v, _ := strconv.Atoi(c.Query(name))
	return v
}

func QueryInt64(c *gin.Context, name string) int64 {
	v, _ := strconv.ParseInt(c.Query(name), 10, 64)
	return v
}
