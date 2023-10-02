package like

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sephix/htmx-player/internal/data"
)

func UpdateTrackLike(c *gin.Context) {
	time.Sleep(500 * time.Millisecond)
	id, _ := strconv.Atoi(c.Param("id"))
	liked := data.UpdateLike(int64(id))
	c.Header("HX-Trigger", "liking-update")
	c.HTML(http.StatusOK, "components/track/like", liked)
}

func UpdateCurrentTrackLike(c *gin.Context) {
	time.Sleep(500 * time.Millisecond)
	id, _ := strconv.Atoi(c.Param("id"))
	liked := data.UpdateLike(int64(id))
	c.Header("HX-Trigger", "liking-current-update")
	c.HTML(http.StatusOK, "components/track/like", liked)
}

func IsTrackLike(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	liked := data.IsTrackLiked(int64(id))
	c.HTML(http.StatusOK, "components/track/like", liked)
}
