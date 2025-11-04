package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pErfEcto2/url_shortener/internal/db"
	"github.com/pErfEcto2/url_shortener/internal/models"
)

func Authorize(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	dbUser, ok := IsValidTokenString(tokenString)
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	c.Set("user", dbUser)
	c.Next()
}

func IsValidTokenString(tokenString string) (models.User, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["sub"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return models.User{}, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.User{}, false
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return models.User{}, false
	}

	dbUser := db.GetUserByUsername(claims["sub"].(string))
	if dbUser.Username == "" || dbUser.Password == "" {
		return models.User{}, false
	}

	return dbUser, true
}
