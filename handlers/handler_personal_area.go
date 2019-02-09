package handlers

import (
	"MSA/auth"
	"github.com/gin-gonic/gin"
)

func ShowPersonalAreaPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Главная страница"},
		"prof-personal-area.html")
}
