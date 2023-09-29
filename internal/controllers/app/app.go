package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/controllers/album"
	"github.com/sephix/htmx-player/internal/controllers/artist"
	"github.com/sephix/htmx-player/internal/controllers/home"
)

func App(router *gin.RouterGroup) {
	router.GET("", home.RenderHome)
	router.GET("artist", artist.RenderAllArtist)
	router.GET("artist/:id", artist.RenderArtist)
	router.GET("album/:id", album.RenderAlbum)
}
