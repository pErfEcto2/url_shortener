package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pErfEcto2/url_shortener/internal/auth"
	"github.com/pErfEcto2/url_shortener/internal/db/memory"
	"github.com/pErfEcto2/url_shortener/internal/handlers"
)

func main() {

	_ = godotenv.Load()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	gin.SetMode(os.Getenv("MODE"))

	db := memory.NewMemoryDB()
	_ = db

	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("static/*.html")

	router.GET("/", handlers.NewRootHandlerGet())

	router.GET("/signup", handlers.NewSignupHandlerGet())
	router.POST("/signup", handlers.NewSignupHandlerPost(db))

	router.GET("/login", handlers.NewLoginHandlerGet())
	router.POST("/login", handlers.NewLoginHandlerPost(db))

	router.GET("/user", auth.Authorize, handlers.NewUserHandlerGet())

	router.POST("/shorten", handlers.NewShortenerHandlerPost(db))

	router.POST("/logout", handlers.NewLogoutHandlerPost())

	router.GET("/:uri", handlers.NewRedirectHandlerGet(db))

	router.Run(host + ":" + port)
}
