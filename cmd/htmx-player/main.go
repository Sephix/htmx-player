package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/controllers"
	"github.com/sephix/htmx-player/internal/data"
)

func main() {

	r := gin.Default()

	funcMap := r.FuncMap
	funcMap["add"] = (func(a, b int) int {
		return a + b
	})
	funcMap["parseDuration"] = (func(duration int64) string {
		return fmt.Sprintf("%02d:%02d", duration/60, duration%60)
	})
	funcMap["parseAlbumDuration"] = (func(tracks []data.Track) string {
		duration := 0
		for _, track := range tracks {
			duration += int(track.Duration)
		}
		fmt.Println("Duration:", duration)
		hours := duration / 3600
		if hours < 1 {
			hours = 1
		}
		minutes := (duration - (3600 / hours)) / 60
		fmt.Println("minutes:", minutes)
		return fmt.Sprintf("%d h %d", hours, minutes)
	})

	r.SetFuncMap(funcMap)
	r.Use(gin.Logger())
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*/**")

	controllers.InitController(r)

	r.Run("0.0.0.0:6942")

}
