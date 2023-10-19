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

func PlaySongFromPlaylist(c *gin.Context) {
	trackId, _ := strconv.Atoi(c.Param("id"))
	data.SetCurrentPlaylistTrack(int64(trackId))
	playlist := data.GetPlaylist()
	var current data.PlaylistElement
	for _, elem := range playlist {
		if elem.Current {
			current = elem
			break
		}
	}
	c.Header("HX-Trigger", "play-song")
	c.HTML(http.StatusOK, "components/playlistContent", gin.H{
		"playlist": playlist,
		"current":  current,
	})
}
func RenderPlaylistPreview(c *gin.Context) {
	trackId := data.GetCurrentPlaylistTrack()
	if trackId != 0 {
		track := data.GetTrackId(int(trackId))
		c.HTML(http.StatusOK, "components/playlist/preview", track)
	} else {
		c.HTML(http.StatusOK, "components/playlist/preview", nil)
	}
}
func PlayNextSong(c *gin.Context) {
	data.SetNextTrackPlaylist()
	c.Header("HX-Trigger", "play-song")
	c.Status(http.StatusOK)
}
