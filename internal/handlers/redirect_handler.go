package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type original_url_getter interface {
	OriginalURLByShortened(shortenedURL string) (string, error)
}

func NewRedirectHandlerGet(db original_url_getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Param("uri")

		shortenedURL := "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + uri

		originalURL, err := db.OriginalURLByShortened(shortenedURL)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		c.Redirect(http.StatusMovedPermanently, originalURL)

	}
}
