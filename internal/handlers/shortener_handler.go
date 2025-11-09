package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pErfEcto2/url_shortener/internal/auth"
	"github.com/pErfEcto2/url_shortener/internal/db"
	"github.com/pErfEcto2/url_shortener/internal/models"
	"github.com/sym01/htmlsanitizer"
)

func ShortenerHandlerPost(c *gin.Context) {
	originUrl := c.Request.Header.Get("Referer")

	if !strings.Contains(originUrl, "user") {
		url := c.PostForm("original_url")
		sanitizedUrl, ok := htmlsanitizer.DefaultURLSanitizer(url)
		if !ok {
			return
		}

		shortenedUrl, ok := db.GetShortenedUrlByUrl(sanitizedUrl)
		if !ok {
			shortenedUrl, ok = db.AddUrlToUser(sanitizedUrl, models.User{Username: "system"})
			if !ok {
				c.Redirect(http.StatusMovedPermanently, "/")
				return
			}
		}

		c.HTML(http.StatusOK, "index_answer.html", gin.H{"shortenedUrl": shortenedUrl})
		return
	}

	// from user page
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	user, ok := auth.IsValidTokenString(tokenString)
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	url := c.PostForm("original_url")
	sanitizedUrl, ok := htmlsanitizer.DefaultURLSanitizer(url)
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/user")
	}
	db.AddUrlToUser(sanitizedUrl, user)

	c.Redirect(http.StatusMovedPermanently, "/user")
}
