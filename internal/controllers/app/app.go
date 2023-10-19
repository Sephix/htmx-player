package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/controllers/album"
	"github.com/sephix/htmx-player/internal/controllers/artist"
	"github.com/sephix/htmx-player/internal/controllers/home"
	"github.com/sephix/htmx-player/internal/controllers/like"
	"github.com/sephix/htmx-player/internal/controllers/player"
	"github.com/sephix/htmx-player/internal/controllers/playlist"
)

func App(router *gin.RouterGroup) {
	router.GET("", home.RenderHome)

	router.GET("artist", artist.RenderAllArtist)
	router.GET("artist/:id", artist.RenderArtist)

	router.GET("album/:id", album.RenderAlbum)

	router.GET("player/:id", player.RenderPlayer)
	router.GET("player/song/:id", player.PlaySong)
	router.GET("player/current", player.PlayCurrentSong)

	router.PUT("track/like/:id", like.UpdateTrackLike)
	router.PUT("track/like/current/:id", like.UpdateCurrentTrackLike)
	router.GET("track/like/:id", like.IsTrackLike)

	router.PUT("playlist/album/:id", playlist.UpdatePlaylist)
	router.PUT("playlist/track/:id", playlist.UpdatePlaylist)
	router.GET("playlist/song/:id", playlist.PlaySongFromPlaylist)
	router.GET("playlist/preview", playlist.RenderPlaylistPreview)
	router.POST("playlist/next", playlist.PlayNextSong)
	router.GET("playlist/current", playlist.RenderPlaylist)
	router.GET("playlist", playlist.RenderPlaylist)
}
