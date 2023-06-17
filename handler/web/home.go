package web

import (
	"embed"
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type HomeWeb interface {
	Index(c *gin.Context)
}

type homeWeb struct {
	embed embed.FS
}

func NewHomeWeb(embed embed.FS) *homeWeb {
	return &homeWeb{embed}
}

func (h *homeWeb) Index(c *gin.Context) {
	var filepath = path.Join("views","main", "index.html")
	var header = path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(h.embed, filepath, header))


	err := tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
