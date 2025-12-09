package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandlerPost(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.Redirect(http.StatusFound, "/")
}
