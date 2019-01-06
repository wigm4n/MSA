package handlers

import (
	"MSA/auth"
	"MSA/entities"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	auth.Render(c, gin.H{
		"title": "Регистрация"}, "register.html")
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// проверить, существует ли уже пользователь в базе
	isAllCorrect := true

	newUser := entities.User{}
	// присвоить новый id, исходя из базы
	count := 1

	if isAllCorrect == true {
		newUser = entities.User{ID: count, Email: email, Password: password}
		if err := newUser.RegisterNewUser(); err == nil {
			token := GenerateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)

			entities.DeleteSessionByUserID(newUser.ID)
			newUser.CreateSession()

			auth.Render(c, gin.H{
				"title": "Личный кабинет"}, "login-successful.html")

		} else {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"ErrorTitle":   "Ошибка регистрации",
				"ErrorMessage": err.Error()})

		}
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Ошибка регистрации",
			"ErrorMessage": "Указанный email уже зарегистрирован"})
	}
}

func ShowLoginPage(c *gin.Context) {
	auth.Render(c, gin.H{"title": "MSA"}, "main-page.html")
}

func PerformLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if entities.IsUserValid(email, password) {
		user, _ := entities.GetUserByEmail(email)
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		entities.DeleteSessionByUserID(user.ID)
		user.CreateSession()

		auth.Render(c, gin.H{
			"title": "Личный кабинет"}, "personal-area.html")

	} else {
		c.HTML(http.StatusUnauthorized, "main-page.html", gin.H{
			"ErrorTitle":   "Ошибка авторизации",
			"ErrorMessage": "Неверные данные учетной записи"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	// передать ID текущего пользователя
	id := 1
	entities.DeleteSessionByUserID(id)

	auth.Render(c,
		gin.H{"title": "Главная страница"}, "main-page.html")
}
