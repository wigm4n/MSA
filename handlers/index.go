package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "Main page")
}
