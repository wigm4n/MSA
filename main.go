package main

import (
	"MSA/auth"
	"MSA/data"
	"MSA/handlers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

var router *gin.Engine

func main() {

	// ЭТО ДЛЯ СТАРОЙ ВЕРСИИ, РАСКОММЕНТИРОВАТЬ, ЕСЛИ ХОЧЕШЬ СТАРЫЕ ПУТИ
	//router = gin.Default()
	//router.LoadHTMLGlob("templates/*")
	//initializeRoutes()
	//initUser()
	//router.Run()

	// ЭТО ДЛЯ НОВОЙ ВЕРСИИ, РАСКОММЕНТИРОВАТЬ, ЕСЛИ ХОЧЕШЬ ПРИВЯЗАТЬ ФРОНТ
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", fs)
	http.HandleFunc("/test", handler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func initUser() {
	user := data.User{Email: "test3@hse.ru", FirstName: "name", LastName: "lastname"}
	user.Password, _ = data.GenerateNewPassword()
	fmt.Println(user.Password)
	user.RegisterNewUser()
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	type Task struct {
		Count  int
		Length int
		From   int
		To     int
		Type   string
		Name   string
	}

	var newTask Task

	newTask.Count, _ = strconv.Atoi(r.FormValue("HowMuch"))
	newTask.Length, _ = strconv.Atoi(r.FormValue("Size"))
	newTask.From, _ = strconv.Atoi(r.FormValue("Min"))
	newTask.To, _ = strconv.Atoi(r.FormValue("Max"))
	newTask.Type = r.FormValue("TaskType")
	newTask.Name = r.FormValue("TaskName")

	fmt.Println(newTask)

	taskJson, _ := json.Marshal(newTask)
	w.Write(taskJson)
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
		//дописать  auth.EnsureLoggedIn() и отладить
		userProf.POST("/personal-area", handlers.PerformLogin)
	}

	userStudent := router.Group("/student")
	{
		userStudent.POST("/", handlers.ShowStudentPage)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "relative/path/to/favicon.ico")
}
