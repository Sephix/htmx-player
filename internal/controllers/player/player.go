package player

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
)

func RenderPlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	track := data.GetTrackId(id)
	artist := data.GetArtistByTrackId(id)

	fmt.Println("Render song:", track.Title)

	c.HTML(http.StatusOK, "components/player/song", gin.H{
		"track":  track,
		"artist": artist,
	})
}
