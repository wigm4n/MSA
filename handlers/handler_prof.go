package handlers

import (
	"MSA/auth"
	"MSA/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ShowTaskCreationPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Создание задания"},
		"prof-task-creation.html")
}

func ShowTaskGettingPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Просмотр заданий"},
		"prof-task-creation.html")
}

func ShowForumPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Форум"},
		"prof-forum.html")
}

func ShowRegistrationPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Регистрация нового преподавателя"},
		"prof-registration.html")
}

func CreateTask(c *gin.Context) {
	count, _ := strconv.Atoi(c.PostForm("count"))
	length, _ := strconv.Atoi(c.PostForm("length"))
	from, _ := strconv.Atoi(c.PostForm("from"))
	to, _ := strconv.Atoi(c.PostForm("to"))

	res, _ := data.SomeCode(count, length, from, to)

	if res == 1 {
		auth.Render(c, gin.H{
			"title": "Задание создано"}, "prof-task-creation-successful.html")
	} else {
		c.HTML(http.StatusUnauthorized, "prof-task-creation.html", gin.H{
			"ErrorTitle":   "Какая-то ошибка",
			"ErrorMessage": "Попробуй снова"})
	}
}
