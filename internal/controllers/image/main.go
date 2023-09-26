package image

import (
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	data "github.com/sephix/htmx-player/internal/data"
)

func GetImageById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := data.GetImageById(id)

	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Type", "image/jpeg")
	c.Status(http.StatusOK)
	jpeg.Encode(c.Writer, result, nil)
}
