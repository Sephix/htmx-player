package home

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
	"github.com/sephix/htmx-player/internal/models"
)

func RenderHome(c *gin.Context) {
	fmt.Println("HOME")
	filterValue := c.Query("artist")
	artists := data.GetAllArtists(filterValue)
	c.HTML(http.StatusOK, "views/homePage", gin.H{
		"artists": artists,
		"nav":     []models.Nav{{"Home", "", true}, {"Artists", "artist", false}, {"Albums", "album", false}},
		"search":  models.Search{"artist", filterValue, "/"},
	})
}
