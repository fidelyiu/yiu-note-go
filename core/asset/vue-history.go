package asset

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
	return &binaryFileSystem{
		fs,
	}
}

func SetVueHistory(r *gin.Engine) {
	// go-bindata -o=core/asset/asset.go -pkg=asset dist/...
	r.Use(static.Serve("/", BinaryFileSystem("dist")))
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})
}
