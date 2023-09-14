package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	homecontent "github.com/sephix/htmx-player/internal/components/home-content"
	"github.com/sephix/htmx-player/internal/controllers/app"
	"github.com/sephix/htmx-player/internal/controllers/home"
	"github.com/sephix/htmx-player/internal/controllers/status"
	loginpage "github.com/sephix/htmx-player/internal/views/login-page"
)

func InitController(router *gin.Engine) {
	status.Status(router)

	router.GET("/login", renderLoginPage)
	router.POST("/login", handleLogin)
	baseRouter := router.Group("/", cookieAuth)

	app.App(baseRouter)
}

func cookieAuth(c *gin.Context) {
	cookie, err := c.Cookie("LOGGED")
	if err != nil || cookie != "true" {
		c.Redirect(http.StatusFound, "/login")
	}
}

func renderLoginPage(c *gin.Context) {
	loginPage := loginpage.LoginPage()
	loginPage.Render(c.Request.Context(), c.Writer)
}

func handleLogin(c *gin.Context) {
	c.SetCookie("LOGGED", "true", 3600, "/", "localhost", false, true)

	if header := c.GetHeader("Hx-Request"); header == "true" {
		c.Header("HX-Push", "/")
		c.Header("Content-Type", "text/html")
		homecontent.MainContent(home.MockArtists).Render(c.Request.Context(), c.Writer)
	} else {
		c.Header("location", "/")
	}
	c.Status(http.StatusOK)
}
