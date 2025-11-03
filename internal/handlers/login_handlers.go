package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pErfEcto2/url_shortener/internal/db"
	"github.com/pErfEcto2/url_shortener/internal/models"
)

func LoginHandlerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginHandlerPost(c *gin.Context) {
	var u models.User

	if err := c.ShouldBind(&u); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "something went wrong..."})
		return
	}

	if !db.HasUser(u) {
		c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "user does not exist"})
		return
	}

	dbUser := db.GetUser(u)

	if !dbUser.CompareHashedPasswords(u.Password) {
		c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.Username,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "something went wrong..."})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.HTML(http.StatusOK, "index.html", nil)
}
