package ginutil

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SRI struct {
	Pattern  string
	Handlers gin.HandlersChain
}

type SimpleRouter struct {
	mm []SRI
}

func NewSimpleRouter() *SimpleRouter {
	return &SimpleRouter{
		make([]SRI, 0),
	}
}

func (si *SimpleRouter) Route(relativePath string, handlerFunc ...gin.HandlerFunc) {
	si.mm = append(si.mm, SRI{
		Pattern:  relativePath,
		Handlers: handlerFunc,
	})
}

func (si *SimpleRouter) StaticIndex(pattern, dir string) {
	si.Route(pattern, func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if !strings.Contains(accept, "html") {
			return
		}

		c.File(dir + "/index.html")
	})
}

func (si *SimpleRouter) StaticFsIndex(pattern string, fs http.FileSystem) {
	si.Route(pattern, func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if !strings.Contains(accept, "html") {
			return
		}

		c.FileFromFS("/", fs)
	})
}

func (si *SimpleRouter) Handler(c *gin.Context) {
	for _, sri := range si.mm {
		if !strings.HasPrefix(c.Request.URL.Path, sri.Pattern) {
			continue
		}

		for _, handler := range sri.Handlers {
			handler(c)

			if c.IsAborted() {
				return
			}
		}
		break
	}
}
