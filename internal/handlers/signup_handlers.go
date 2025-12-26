package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pErfEcto2/url_shortener/internal/models"
)

func NewSignupHandlerGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)

	}
}

type check_and_add_user interface {
	HasUserByUsername(string) bool
	AddUser(models.User) error
}

func NewSignupHandlerPost(db check_and_add_user) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User
		if err := c.ShouldBind(&u); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", map[string]string{"Error": "something went wrong..."})
			return
		}

		if db.HasUserByUsername(u.Username) {
			c.HTML(http.StatusBadRequest, "signup.html", map[string]string{"Error": "user already exists"})
			return
		}

		if err := u.HashPassword(); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", map[string]string{"Error": "something went wrong..."})
			return
		}

		db.AddUser(u)

		c.Redirect(http.StatusMovedPermanently, "/login")

	}
}
