package ginutil

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	httpSrv  *http.Server
	indexMap map[string]*Index
}

func NewServer(addr string) *RestServer {
	router := gin.Default()
	return &RestServer{
		Engine: router,
		httpSrv: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		indexMap: make(map[string]*Index),
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

func (rs *RestServer) SetupIndex(relativePath string, index *Index) {
	rs.indexMap[relativePath] = index
}

func (rs *RestServer) SetupStatic(relativePath string, staticDir string) {
	if staticDir == "" {
		return
	}

	router := &rs.Engine.RouterGroup
	if relativePath != "/" {
		router = router.Group(relativePath)
	}

	_, rootDirName := filepath.Split(staticDir)
	staticLoader := func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		if info.IsDir() && info.Name() != rootDirName {
			router.Static(info.Name(), path)
		}

		return nil
	}

	if err := filepath.Walk(staticDir, staticLoader); err != nil {
		log.Fatalln(err)
	}
}

func (rs *RestServer) SetupPing() {
	pingHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}

	rs.HEAD("/ping", pingHandler)
	rs.GET("/ping", pingHandler)
}

func (rs *RestServer) setupNoRouter() {
	rs.NoRoute(func(c *gin.Context) {
		if index, ok := rs.matchIndex(c.Request.URL.Path); ok {
			index.run(c)
			return
		}
	})
}

func (rs *RestServer) matchIndex(path string) (*Index, bool) {
	for key, index := range rs.indexMap {
		if strings.HasPrefix(path, key) {
			return index, true
		}
	}

	return nil, false
}

func (rs *RestServer) Run() error {
	rs.setupNoRouter()

	log.Printf("[rest server started, listen %s]", rs.httpSrv.Addr)
	if err := rs.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("[rest server listen failed: %v]", err)
	}

	return nil
}

func (rs *RestServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rs.httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("[rest server shutdown err:]", err)
	}

	log.Printf("[rest server exited.]")
}
