package artist

import (
	"fmt"
	"strconv"

	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
	"github.com/sephix/htmx-player/internal/models"
)

func RenderAllArtist(c *gin.Context) {
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
			"nav":     []models.Nav{{"Home", "", false}, {"Artists", "artist", true}, {"Albums", "album", false}},
			"search":  models.Search{"artist", filterValue, "/"},
		})
	}
}

func RenderArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	artist := data.GetArtistById(id)

	filterValue := c.Query("album")
	albums := data.GetAlbumByArtistId(id, filterValue)

	files := []string{
		"./templates/views/base.html",
		"./templates/views/artistPage.html",
		"./templates/components/header.html",
		"./templates/components/nav.html",
		"./templates/components/artistContent.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		tmpl.ExecuteTemplate(c.Writer, "views/base.html", gin.H{
			"artist": artist,
			"albums": albums,
			"nav":    []models.Nav{{"Home", "", false}, {"Artists", "artist", true}, {"Albums", "album", false}},
			"search": models.Search{"album", filterValue, fmt.Sprintf("/artist/%v", id)},
		})
	}
}
