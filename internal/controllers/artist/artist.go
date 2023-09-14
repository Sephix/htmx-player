package artist

import (
	"strconv"

	"github.com/gin-gonic/gin"
	artistcontent "github.com/sephix/htmx-player/internal/components/artist-content"
	"github.com/sephix/htmx-player/internal/data/artist"
)

var MockArtists []artist.Artist = []artist.Artist{
	{Name: "John Lennon", Img: "1"},
	{Name: "Paul McCartney", Img: "2"},
	{Name: "George Harrison", Img: "3"},
	{Name: "Ringo Star", Img: "4"},
	{Name: "The Beatles", Img: "5"},
	{Name: "The White Stripes", Img: "6"},
	{Name: "Jack White", Img: "7"},
	{Name: "Taylor Swift", Img: "8"},
	{Name: "The Red Hot Chili Peppers", Img: "9"},
	{Name: "John Frusciante", Img: "10"},
}

func RenderArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	artistcontent.ArtistContent(MockArtists[id-1]).Render(c.Request.Context(), c.Writer)
}
