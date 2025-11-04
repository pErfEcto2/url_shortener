package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pErfEcto2/url_shortener/internal/auth"
	"github.com/pErfEcto2/url_shortener/internal/db"
)

func ShortenerHandlerPost(c *gin.Context) {
	originUrl := c.Request.Header.Get("Referer")

	if !strings.Contains(originUrl, "user") {
		// from home page
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
	if ok := db.AddUrlToUser(url, user); !ok {
		return
	}

	fmt.Println(user)
}
