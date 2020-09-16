//go:generate go-bindata -pkg ui -o bindata.go --prefix react/build/ react/build/...
//go:generate go fmt .

package ui

import (
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func RenderIndex(c *gin.Context) {
	idx, err := Asset("index.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.Data(http.StatusOK, "text/html", idx)
}

func ServeUI(c *gin.Context) {
	fileName := strings.TrimPrefix(c.Request.URL.Path, "/")

	file, err := Asset(fileName)
	if err != nil {
		RenderIndex(c)
		return
	}

	ct := mime.TypeByExtension(path.Ext(fileName))
	c.Data(http.StatusOK, ct, file)
}
