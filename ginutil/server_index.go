package ginutil

import "github.com/gin-gonic/gin"

type Index struct {
	filepath string
	handlers gin.HandlersChain
}

func NewIndex(filepath string, hooks ...gin.HandlerFunc) *Index {
	return &Index{filepath: filepath, handlers: hooks}
}

func (i *Index) run(c *gin.Context) {
	for _, handler := range i.handlers {
		handler(c)
	}
	if c.IsAborted() {
		return
	}

	c.File(i.filepath)
}
