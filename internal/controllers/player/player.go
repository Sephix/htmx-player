package player

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
)

func RenderPlayer(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Param("id"))
	trackId, _ := strconv.Atoi(c.Query("track"))

	track := data.GetTrackId(trackId)
	artist := data.GetArtistByTrackId(trackId)

	data.AddAblumToPlaylist(int64(albumId), int64(trackId))

	c.Header("HX-Trigger", "playlist-update")
	c.HTML(http.StatusOK, "components/player/song", gin.H{
		"track":  track,
		"artist": artist,
	})
}
