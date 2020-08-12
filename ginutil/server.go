package ginutil

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Resource interface {
	Register(router *gin.RouterGroup)
}

// RestServer
type RestServer struct {
	*gin.Engine

	srv *http.Server
}

func NewServer(addr string) *RestServer {
	router := gin.Default()
	return &RestServer{
		Engine: router,
		srv: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (rs *RestServer) SetupRS(resources ...Resource) {
	for _, resource := range resources {
		resource.Register(&rs.RouterGroup)
	}
}

func (rs *RestServer) SetupGroupRS(relativePath string, resources ...Resource) {
	for _, resource := range resources {
		resource.Register(rs.Group(relativePath))
	}
}

func (rs *RestServer) SetupSwagger() {
	rs.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (rs *RestServer) SetupPing() {
	pingHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}

	rs.HEAD("/ping", pingHandler)
	rs.GET("/ping", pingHandler)
}

func (rs *RestServer) SetupIndex(indexDir string) {
	rs.LoadHTMLGlob(filepath.Join(indexDir, "index.html"))
	rs.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

func (rs *RestServer) SetupStatic(relativePath string, staticDir string) {
	if staticDir == "" {
		return
	}

	router := &rs.Engine.RouterGroup
	if relativePath != "/" {
		router = router.Group(relativePath)
	}

	staticLoader := func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		if info.IsDir() && info.Name() != staticDir {
			router.Static(info.Name(), path)
		}

		return nil
	}

	if err := filepath.Walk(staticDir, staticLoader); err != nil {
		log.Fatalln(err)
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
