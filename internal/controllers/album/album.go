package album

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
	"github.com/sephix/htmx-player/internal/models"
)

func RenderAlbum(c *gin.Context) {
	time.Sleep(250 * time.Millisecond)
	id, _ := strconv.Atoi(c.Param("id"))

	tracks := data.GetTrackAlbumId(id)
	album := data.GetAlbumById(id)
	artist := data.GetArtistByAlbumId(id)

	c.HTML(http.StatusOK, "views/albumPage", gin.H{
		"artist": artist,
		"album":  album,
		"tracks": tracks,
		"nav":    []models.Nav{getNav("Home", "", false), getNav("Artists", "artist", false), getNav("Albums", "album", true)},
		"search": nil,
	})
}

func getNav(title, link string, isActive bool) models.Nav {
	return models.Nav{
		Title:    title,
		Link:     link,
		IsActive: isActive,
	}
}
func getSearch(field, value, url string) models.Search {
	return models.Search{
		Field: field,
		Value: value,
		Url:   url,
	}
}
