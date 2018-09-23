package main

import (
	"MSA/handlers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlers.ShowMainPage)

	router.Run()
}
