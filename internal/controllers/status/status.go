package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(router *gin.Engine) {
	status := router.Group("/status")
	{
		status.GET("/ok", statusOk)
	}
}

func statusOk(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
