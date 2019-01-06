package main

import (
	"MSA/auth"
	"MSA/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	router.Run()
}

func initializeRoutes() {
	// Используем middleware метод setUserStatus для каждого маршрута, чтобы установить флаг,
	// который указывает, был ли запрос от аутентифицированного пользователя или нет
	router.Use(auth.SetUserStatus())

	http.HandleFunc("/favicon.ico", faviconHandler)

	// Стартовая страница
	router.GET("/", handlers.ShowMainPage)

	userRoutes := router.Group("/auth")
	{
		userRoutes.GET("/register", auth.EnsureLoggedIn(), handlers.ShowRegistrationPage)
		userRoutes.POST("/register", auth.EnsureLoggedIn(), handlers.Register)

		userRoutes.GET("/login", handlers.ShowLoginPage)
		userRoutes.POST("/login", handlers.PerformLogin)

		userRoutes.GET("/logout", auth.EnsureLoggedIn(), handlers.Logout)
	}

	personalAreaRoutes := router.Group("/personal")
	{
		personalAreaRoutes.GET("/area", auth.EnsureLoggedIn(), handlers.ShowPersonalAreaPage)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "relative/path/to/favicon.ico")
}
