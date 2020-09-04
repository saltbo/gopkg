package ginutil

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupEmbedAssets(rg *gin.RouterGroup, fs http.FileSystem, relativePaths ...string) {
	handler := func(c *gin.Context) {
		c.FileFromFS(strings.TrimPrefix(c.Request.URL.Path, rg.BasePath()), fs)
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
