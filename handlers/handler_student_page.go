package handlers

import (
	"MSA/auth"
	"github.com/gin-gonic/gin"
)

func ShowStudentPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title": "Страница студента"},
		"student-main-page.html")
}
