package home

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
	"github.com/sephix/htmx-player/internal/models"
)

func RenderHome(c *gin.Context) {
	filterValue := c.Query("artist")
	artists := data.GetAllArtists(filterValue)
	files := []string{
		"./templates/views/base.html",
		"./templates/views/homePage.html",
		"./templates/components/header.html",
		"./templates/components/nav.html",
		"./templates/components/homeContent.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		tmpl.ExecuteTemplate(c.Writer, "views/base.html", gin.H{
			"artists": artists,
			"nav":     []models.Nav{{"Home", "", true}, {"Artists", "artist", false}, {"Albums", "album", false}},
			"search":  models.Search{"artist", filterValue, "/"},
		})
	}
}
