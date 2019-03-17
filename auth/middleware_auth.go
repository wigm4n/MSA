package auth

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

func GenerateSessionToken() (password string) {
	bytes := make([]byte, 16)
	for i := 0; i < 16; i++ {
		nBig, _ := rand.Int(rand.Reader, big.NewInt(25))
		bytes[i] = byte(65 + nBig.Int64())
	}
	password = string(bytes)
	return
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Респонс-формат JSON
		c.JSON(http.StatusOK, data["payload"])
	default:
		// Респонт-формат HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
