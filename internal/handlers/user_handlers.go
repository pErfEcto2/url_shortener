package handlers

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/pErfEcto2/url_shortener/internal/models"
)

func map2slice(m map[string]string) [][]string {
	var res [][]string

	for k, v := range m {
		res = append(res, []string{k, v})
	}
	return res
}

func NewUserHandlerGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if !ok {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		urlsPairs := map2slice(user.(models.User).Urls)
		slices.Reverse(urlsPairs)
		c.HTML(http.StatusOK, "user_page.html", gin.H{"urls": urlsPairs})
	}
}
