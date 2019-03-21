package data

import (
	"encoding/json"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"os"
)

type emailCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var credentials emailCredentials

func init() {
	file, err := ioutil.ReadFile("./properties_email.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &credentials)
	return
}

func ResetPasswordEmail(email string, newPassword string) (exists bool) {
	m := gomail.NewMessage()
	m.SetHeader("From", credentials.Email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Восстановление пароля")
	m.SetBody("text/html", "Доброго времени дня, "+email+"!"+
		"<p>Вы запросили восстановление пароля."+
		"<p>Теперь ваш новый пароль:<b>  "+newPassword+"</b>"+
		"<p>С уважением,<br>Сервис восстановления паролей портала pthams.hse.ru.")

	d := gomail.NewDialer("smtp.gmail.com", 587, credentials.Email, credentials.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return false
	}
	return true
}

func CreateNewUserEmail(user User) (exists bool) {
	m := gomail.NewMessage()
	m.SetHeader("From", credentials.Email)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Регистрация на сервисе pthams.hse.ru")
	m.SetBody("text/html", "Доброго времени дня, "+user.LastName+" "+user.FirstName+
		" "+user.Patronymic+"!"+
		"<p>Вы зарегистрированны на портале "+"<a href=\"http://pthams.hse.ru\">pthams.hse.ru</a>.</p>"+
		"<p>Ваш логин:<b>  "+user.Email+"</b>"+
		"<br>Ваш пароль:<b>  "+user.Password+"</b>"+
		"<p><p>Пароль возможно сменить в настройках личного кабинета на портале."+
		"<p>С уважением,<br>Сервис регистрации портала pthams.hse.ru.")

	d := gomail.NewDialer("smtp.gmail.com", 587, credentials.Email, credentials.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return false
	}
	return true
}
