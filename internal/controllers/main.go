package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/controllers/app"
	"github.com/sephix/htmx-player/internal/controllers/home"
	"github.com/sephix/htmx-player/internal/controllers/image"
	"github.com/sephix/htmx-player/internal/controllers/status"
)

func InitController(router *gin.Engine) {
	status.Status(router)

	router.GET("/login", renderLoginPage)
	router.POST("/login", handleLogin)
	router.GET("/image/:id", image.GetImageById)
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
	files := []string{
		"./templates/views/base.html",
		"./templates/views/loginPage.html",
		"./templates/components/header.html",
		"./templates/components/login.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("PLOP")
		tmpl.ExecuteTemplate(c.Writer, "views/base.html", nil)
	}
}

func handleLogin(c *gin.Context) {
	c.SetCookie("LOGGED", "true", 3600, "/", "localhost", false, true)

	if header := c.GetHeader("Hx-Request"); header == "true" {
		c.Header("HX-Push", "/")
		c.Header("Content-Type", "text/html")
		home.RenderHome(c)
	} else {
		c.Header("location", "/")
	}
	c.Status(http.StatusOK)
}
