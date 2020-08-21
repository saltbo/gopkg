package ginutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/saltbo/gopkg/httputil"
)

type ResponseFunc func(c *gin.Context, err error)
type RedirectFunc func(c *gin.Context, location string)

func TestJSONError(t *testing.T) {
	rhs := map[int]ResponseFunc{
		http.StatusBadRequest:          JSONBadRequest,
		http.StatusUnauthorized:        JSONUnauthorized,
		http.StatusForbidden:           JSONForbidden,
		http.StatusInternalServerError: JSONServerError,
	}
	for status, rh := range rhs {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		rh(c, fmt.Errorf("test"))
		assert.Equal(t, status, w.Code)
	}
}

func TestRedirect(t *testing.T) {
	rhs := map[int]RedirectFunc{
		http.StatusFound:             FoundRedirect,
		http.StatusMovedPermanently:  MovedRedirect,
		http.StatusTemporaryRedirect: TemporaryRedirect,
	}
	for status, rh := range rhs {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "http://example.com", nil)
		rh(c, "http://saltbo.cn")
		c.Writer.WriteHeaderNow()
		assert.Equal(t, status, w.Code)
		assert.Equal(t, "http://saltbo.cn", w.Header().Get("Location"))
	}
}

func TestJSONData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	JSONData(c, "test")
	assert.Equal(t, http.StatusOK, w.Code)
	jr := httputil.BindJSONResponse(w.Body.Bytes())
	assert.Equal(t, "test", jr.Data)
}

func TestJSONList(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	JSONList(c, []string{"test1", "test2"}, 2)
	assert.Equal(t, http.StatusOK, w.Code)
	jr := httputil.BindJSONResponse(w.Body.Bytes())
	data := jr.Data.(map[string]interface{})
	assert.Equal(t, []interface{}{"test1", "test2"}, data["list"].([]interface{}))
	assert.Equal(t, 2, int(data["total"].(float64)))
}

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	JSON(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Cookie(c, "test", "aaa", 10)
	cookie := w.Result().Cookies()[0]
	assert.Equal(t, "test", cookie.Name)
	assert.Equal(t, "aaa", cookie.Value)
	assert.Equal(t, 10, cookie.MaxAge)
}
