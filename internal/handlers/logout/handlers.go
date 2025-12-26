package logout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLogoutHandlerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("Authorization", "", -1, "", "", false, true)

		c.Redirect(http.StatusFound, "/")

	}
}
