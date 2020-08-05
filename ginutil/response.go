package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// JSONResponse
func NewJSONResponse(data interface{}) *JSONResponse {
	return &JSONResponse{Code: 0, Msg: "ok", Data: data}
}

// JSON
func JSON(c *gin.Context) {
	c.JSON(http.StatusOK, NewJSONResponse(nil))
}

// JSONError
func JSONError(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, JSONResponse{
		Code: status,
		Msg:  err.Error(),
	})
	c.Error(err)
}

// JSONData
func JSONData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewJSONResponse(data))
}

// JSONList
func JSONList(c *gin.Context, list interface{}, total int64) {
	c.JSON(http.StatusOK, NewJSONResponse(gin.H{
		"list":  list,
		"total": total,
	}))
}

// BadRequest
func BadRequest(c *gin.Context, err error) {
	JSONError(c, http.StatusBadRequest, err)
}

// Unauthorized
func Unauthorized(c *gin.Context, err error) {
	JSONError(c, http.StatusUnauthorized, err)
}

// Forbidden
func Forbidden(c *gin.Context, err error) {
	JSONError(c, http.StatusForbidden, err)
}

// ServerError
func ServerError(c *gin.Context, err error) {
	JSONError(c, http.StatusInternalServerError, err)
}
