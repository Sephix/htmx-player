package artist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
	"github.com/sephix/htmx-player/internal/models"
)

func RenderAllArtist(c *gin.Context) {
	filterValue := c.Query("artist")
	artists := data.GetAllArtists(filterValue)
	c.HTML(http.StatusOK, "views/homePage", gin.H{
		"artists": artists,
		"nav":     []models.Nav{models.GetNav("Home", "", false), models.GetNav("Artists", "artist", true), models.GetNav("Albums", "album", false)},
		"search":  models.GetSearch("artist", filterValue, "/"),
	})
}

func RenderArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	artist := data.GetArtistById(id)

	filterValue := c.Query("album")
	albums := data.GetAlbumByArtistId(id, filterValue)

	c.HTML(http.StatusOK, "views/artistPage", gin.H{
		"artist": artist,
		"albums": albums,
		"nav":    []models.Nav{models.GetNav("Home", "", false), models.GetNav("Artists", "artist", true), models.GetNav("Albums", "album", false)},
		"search": models.GetSearch("album", filterValue, fmt.Sprintf("/artist/%v", id)),
	})
}
