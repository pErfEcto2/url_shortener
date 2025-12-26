package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pErfEcto2/url_shortener/internal/auth"
	"github.com/pErfEcto2/url_shortener/internal/db/memory"
	"github.com/pErfEcto2/url_shortener/internal/handlers/login"
	"github.com/pErfEcto2/url_shortener/internal/handlers/logout"
	"github.com/pErfEcto2/url_shortener/internal/handlers/redirect"
	"github.com/pErfEcto2/url_shortener/internal/handlers/root"
	"github.com/pErfEcto2/url_shortener/internal/handlers/shortener"
	"github.com/pErfEcto2/url_shortener/internal/handlers/signup"
	"github.com/pErfEcto2/url_shortener/internal/handlers/user_page"
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

	router.GET("/", root.NewRootHandlerGet())

	router.GET("/signup", signup.NewSignupHandlerGet())
	router.POST("/signup", signup.NewSignupHandlerPost(db))

	router.GET("/login", login.NewLoginHandlerGet())
	router.POST("/login", login.NewLoginHandlerPost(db))

	router.GET("/user", auth.Authorize, user_page.NewUserHandlerGet())

	router.POST("/shorten", shortener.NewShortenerHandlerPost(db))

	router.POST("/logout", logout.NewLogoutHandlerPost())

	router.GET("/:uri", redirect.NewRedirectHandlerGet(db))

	router.Run(host + ":" + port)
}
