package main

import (
	"MSA/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlers.ShowMainPage)

	fmt.Println("testing hooks")

	router.Run()
}
