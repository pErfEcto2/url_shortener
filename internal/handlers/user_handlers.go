package handlers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/pErfEcto2/url_shortener/internal/db"
	// "github.com/pErfEcto2/url_shortener/internal/models"
)

func UserHandelerGet(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	c.HTML(http.StatusOK, "user_page.html", nil)
}
