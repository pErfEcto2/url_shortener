package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pErfEcto2/url_shortener/internal/db"
	"github.com/pErfEcto2/url_shortener/internal/models"
	"net/http"
)

func SignupHandlerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupHandlerPost(c *gin.Context) {
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
