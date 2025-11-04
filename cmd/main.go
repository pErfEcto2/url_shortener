package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pErfEcto2/url_shortener/internal/auth"
	"github.com/pErfEcto2/url_shortener/internal/handlers"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	_ = godotenv.Load()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("static/*.html")

	router.GET("/", handlers.RootHandlerGet)

	router.GET("/signup", handlers.SignupHandlerGet)
	router.POST("/signup", handlers.SignupHandlerPost)

	router.GET("/login", handlers.LoginHandlerGet)
	router.POST("/login", handlers.LoginHandlerPost)

	router.GET("/user", auth.Authorize, handlers.UserHandelerGet)

	router.POST("/shorten", handlers.ShortenerHandlerPost)

	router.Run(host + ":" + port)
}
