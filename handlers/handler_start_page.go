package handlers

import (
	"MSA/auth"
	"github.com/gin-gonic/gin"
)

func ShowMainPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Главная страница"},
		"index.html")
}
