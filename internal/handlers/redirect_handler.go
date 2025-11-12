package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pErfEcto2/url_shortener/internal/db"
)

func RedirectHandlerGet(c *gin.Context) {
	uri := c.Param("uri")

	shortenedURL := "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + uri
	originalURL, err := db.GetOriginalURLByShortened(shortenedURL)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
