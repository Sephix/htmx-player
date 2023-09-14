package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/controllers"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*/**")

	controllers.InitController(r)

	r.Run("0.0.0.0:6942")

}
