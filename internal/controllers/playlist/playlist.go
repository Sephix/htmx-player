package playlist

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
)

func RenderPlaylist(c *gin.Context) {
	playlist := data.GetPlaylist()
	var current data.PlaylistElement
	for _, elem := range playlist {
		if elem.Current {
			current = elem
			break
		}
	}
	c.HTML(http.StatusOK, "components/playlistContent", gin.H{
		"playlist": playlist,
		"current":  current,
	})
}

func UpdatePlaylist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	filterValue, _ := strconv.Atoi(c.Query("track"))
	data.AddAblumToPlaylist(int64(id), int64(filterValue))
}
