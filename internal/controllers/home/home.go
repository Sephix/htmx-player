package home

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
)

func RenderHome(c *gin.Context) {
	artists := data.GetAllArtists()
	if header := c.GetHeader("Hx-Request"); header == "true" {
		c.HTML(200, "components/homeContent.html", artists)
	} else {
		files := []string{
			"./templates/views/base.html",
			"./templates/views/homePage.html",
			"./templates/components/header.html",
			"./templates/components/homeContent.html",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			tmpl.ExecuteTemplate(c.Writer, "views/base.html", artists)
		}
	}
}
