package artist

import (
	"fmt"
	"net/http"
	"strconv"

	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
)

func RenderArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	artist := data.GetArtistById(id)
	albums := data.GetAlbumByArtistId(id)

	fmt.Println("Alb img: ", albums[0].Img)

	if header := c.GetHeader("Hx-Request"); header == "true" {
		c.HTML(http.StatusOK, "components/artistContent.html", gin.H{
			"artist": artist,
			"albums": albums,
		})
	} else {
		files := []string{
			"./templates/views/base.html",
			"./templates/views/artistPage.html",
			"./templates/components/header.html",
			"./templates/components/artistContent.html",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			tmpl.ExecuteTemplate(c.Writer, "views/base.html", gin.H{
				"artist": artist,
				"albums": albums,
			})
		}
	}
}
