package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pErfEcto2/url_shortener/internal/handlers"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	_ = godotenv.Load()

	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("static/*.html")

	router.GET("/", handlers.RootHandler)

	router.GET("/signup", handlers.SignupHandlerGet)
	router.POST("/signup", handlers.SignupHandlerPost)

	router.GET("/login", handlers.LoginHandlerGet)
	router.POST("/login", handlers.LoginHandlerPost)

	router.Run("localhost:8080")
}
