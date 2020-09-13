//go:generate go-bindata -pkg main -o bindata.go --prefix ui/build/ ui/build/...
//go:generate go fmt .

package main

import (
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveUI(c *gin.Context) {
	fileName := strings.TrimPrefix(c.Request.URL.Path, "/")

	file, err := Asset(fileName)
	if err != nil {
		idx, err := Asset("index.html")
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Something went wrong")
			return
		}
		c.Data(http.StatusOK, "text/html", idx)
		return
	}

	ct := mime.TypeByExtension(path.Ext(fileName))
	c.Data(http.StatusOK, ct, file)
}
