package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/controllers/app"
	"github.com/sephix/htmx-player/internal/controllers/image"
	"github.com/sephix/htmx-player/internal/controllers/login"
	"github.com/sephix/htmx-player/internal/controllers/song"
	"github.com/sephix/htmx-player/internal/controllers/status"
)

func InitController(router *gin.Engine) {
	status.Status(router)

	loginGroup := router.Group("/login")
	login.LoginController(loginGroup)
	router.GET("/image/:id", image.GetImageById)
	router.GET("/song/:id", song.GetSongById)
	baseRouter := router.Group("/", cookieAuth)

	app.App(baseRouter)
}

func cookieAuth(c *gin.Context) {
	cookie, err := c.Cookie("LOGGED")
	if err != nil || cookie != "true" {
		c.Redirect(http.StatusFound, "/login")
	}
}
