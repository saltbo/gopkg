package ginutil

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Resource interface {
	Register(router *gin.RouterGroup)
}

// RestServer
type RestServer struct {
	srv    *http.Server
	router *gin.Engine
}

func NewServer(addr string) (*RestServer, error) {
	router := gin.Default()
	return &RestServer{
		srv: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		router: router,
	}, nil
}

func (rs *RestServer) Use(middleware ...gin.HandlerFunc) {
	rs.router.Use(middleware...)
}

func (rs *RestServer) SetupResource(relativePath string, resources ...Resource) {
	for _, resource := range resources {
		resource.Register(rs.router.Group(relativePath))
	}
}

func (rs *RestServer) Run() error {
	log.Printf("[rest server started, listen %s]", rs.srv.Addr)
	if err := rs.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("[rest server listen failed: %v]", err)
	}

	return nil
}

func (rs *RestServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rs.srv.Shutdown(ctx); err != nil {
		log.Fatal("[rest server shutdown err:]", err)
	}

	log.Printf("[rest server exited.]")
}

func (rs *RestServer) SetupPing() {
	pingHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}

	rs.router.HEAD("/ping", pingHandler)
	rs.router.GET("/ping", pingHandler)
}
