package ginutil

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

var defaultEmbedFS http.FileSystem

func EmbedFS() http.FileSystem {
	var once sync.Once
	once.Do(func() {
		efs, err := fs.New()
		if err != nil {
			log.Fatalln(err)
		}

		defaultEmbedFS = efs
	})

	return defaultEmbedFS
}

func SetupEmbedAssets(rg *gin.RouterGroup, relativePaths ...string) {
	handler := func(c *gin.Context) {
		http.FileServer(EmbedFS()).ServeHTTP(c.Writer, c.Request)
	}

	for _, relativePath := range relativePaths {
		urlPattern := relativePath
		if urlPattern != "/" {
			urlPattern = path.Join(relativePath, "/*filepath")
		}

		rg.GET(urlPattern, handler)
		rg.HEAD(urlPattern, handler)
	}
}

func SetupStaticAssets(rg *gin.RouterGroup, dir string) {
	_, rootDirName := filepath.Split(dir)
	staticLoader := func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		if info.IsDir() && info.Name() != rootDirName {
			rg.Static(info.Name(), path)
			return nil
		}

		if info.Name() == "index.html" {
			rg.StaticFile("/", path)
		}

		return nil
	}

	if err := filepath.Walk(dir, staticLoader); err != nil {
		log.Fatalln(err)
	}
}
