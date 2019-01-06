package handlers

import (
	"MSA/auth"
	"MSA/entities"
	"github.com/gin-gonic/gin"
)

func ShowMainPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Главная страница"},
		"main-page.html")
}

func ShowPersonalAreaPage(c *gin.Context) {
	session, _ := entities.GetCurrentSession()
	user, _ := entities.GetUserById(session.UserID)

	auth.Render(c, gin.H{
		"title": "Личный кабинет", "payload": user}, "personal-area.html")
}
