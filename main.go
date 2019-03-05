package main

import (
	"MSA/auth"
	"MSA/data"
	"MSA/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	initUser()
	router.Run()
}

func initUser() {
	user := data.User{Email: "test3@hse.ru", FirstName: "name", LastName: "lastname"}
	user.Password, _ = data.GenerateNewPassword()
	fmt.Println(user.Password)
	user.RegisterNewUser()
}

func initializeRoutes() {
	// Используем middleware метод setUserStatus для каждого маршрута, чтобы установить флаг,
	// который указывает, был ли запрос от аутентифицированного пользователя или нет
	router.Use(auth.SetUserStatus())

	http.HandleFunc("/favicon.ico", faviconHandler)

	// Стартовая страница
	router.GET("/", handlers.ShowMainPage)

	userProf := router.Group("/prof")
	{
		userProf.GET("/login", handlers.ShowLoginPage)

		userProf.GET("/logout", auth.EnsureLoggedIn(), handlers.Logout)

		userProf.GET("/create-task", auth.EnsureLoggedIn(), handlers.ShowTaskCreationPage)

		userProf.POST("/create-task-successful", auth.EnsureLoggedIn(), handlers.CreateTask)

		userProf.GET("/get-tasks", auth.EnsureLoggedIn(), handlers.ShowTaskGettingPage)

		userProf.GET("/registration", auth.EnsureLoggedIn(), handlers.ShowRegistrationPage)

		userProf.POST("/registration-successful", auth.EnsureLoggedIn(), handlers.Registration)

		userProf.GET("/forum", auth.EnsureLoggedIn(), handlers.ShowForumPage)

		userProf.GET("/personal-area", auth.EnsureLoggedIn(), handlers.ShowPersonalAreaPage)
		userProf.POST("/personal-area", auth.EnsureLoggedIn(), handlers.PerformLogin)
	}

	userStudent := router.Group("/student")
	{
		userStudent.POST("/", handlers.ShowStudentPage)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "relative/path/to/favicon.ico")
}
