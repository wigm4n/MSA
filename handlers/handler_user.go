package handlers

import (
	"MSA/auth"
	"MSA/data"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func Registration(c *gin.Context) {
	email := c.PostForm("email")
	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")

	newUser := data.User{Email: email, FirstName: firstName, LastName: lastName}

	isAllCorrect, err := data.IsExist(email)
	if err != nil {
		log.Println(err)
		return
	}

	if isAllCorrect == true {
		newUser.Password, _ = data.GenerateNewPassword()
		if err := newUser.RegisterNewUser(); err == nil {
			token := GenerateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)

			//?????
			//data.DeleteSessionByUserID(newUser.ID)
			//newUser.CreateSession()

			auth.Render(c, gin.H{
				"title": "Личный кабинет"}, "prof-registration-successful.html")

		} else {
			c.HTML(http.StatusBadRequest, "prof-registration.html", gin.H{
				"ErrorTitle":   "Ошибка регистрации",
				"ErrorMessage": err.Error()})

		}
	} else {
		c.HTML(http.StatusBadRequest, "prof-registration.html", gin.H{
			"ErrorTitle":   "Ошибка регистрации",
			"ErrorMessage": "Указанный email уже зарегистрирован"})
	}
}

func ShowLoginPage(c *gin.Context) {
	auth.Render(c, gin.H{"title": "MSA"}, "prof-login.html")
}

func PerformLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if data.IsUserValid(email, password) {
		user, _ := data.GetUserByEmail(email)
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		//??????
		data.DeleteSessionByUserID(user.ID)
		user.CreateSession()

		auth.Render(c, gin.H{
			"title": "Личный кабинет"}, "prof-personal-area.html")

	} else {
		c.HTML(http.StatusUnauthorized, "prof-login.html", gin.H{
			"ErrorTitle":   "Ошибка авторизации",
			"ErrorMessage": "Неверные данные учетной записи"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	// передать ID текущего пользователя
	id := 1
	data.DeleteSessionByUserID(id)

	auth.Render(c,
		gin.H{"title": "Главная страница"}, "index.html")
}
