package song

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "github.com/sephix/htmx-player/internal/data"
)

func GetSongById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := data.GetSongById(id)

	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Type", "audio/mpeg")
	c.Data(http.StatusOK, "audio/mpeg", result)
}
