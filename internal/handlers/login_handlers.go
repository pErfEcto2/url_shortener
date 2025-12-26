package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pErfEcto2/url_shortener/internal/models"
)

func NewLoginHandlerGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)

	}
}

type user_getter interface {
	UserByUsername(username string) models.User
}

func NewLoginHandlerPost(db user_getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User

		if err := c.ShouldBind(&u); err != nil {
			c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "something went wrong..."})
			return
		}

		dbUser := db.UserByUsername(u.Username)
		if dbUser.Username == "" || dbUser.Password == "" {
			c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "user does not exist"})
			return
		}

		if !dbUser.CompareHashedPasswords(u.Password) {
			c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "invalid username or password"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": dbUser.Username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.HTML(http.StatusBadRequest, "login.html", map[string]string{"Error": "something went wrong..."})
			return
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		c.Redirect(http.StatusFound, "/user")

	}
}
