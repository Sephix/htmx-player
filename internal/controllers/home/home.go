package home

import (
	"github.com/gin-gonic/gin"
	homecontent "github.com/sephix/htmx-player/internal/components/home-content"
	"github.com/sephix/htmx-player/internal/data/artist"
	homepage "github.com/sephix/htmx-player/internal/views/home-page"
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

func RenderHome(c *gin.Context) {

	if header := c.GetHeader("Hx-Request"); header == "true" {
		homecontent.MainContent(MockArtists).Render(c.Request.Context(), c.Writer)
	} else {
		homepage.HomePage(MockArtists).Render(c.Request.Context(), c.Writer)
	}
}
