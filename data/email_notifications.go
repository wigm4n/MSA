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

func SendEmail(email string, newPassword string) (exists bool) {
	m := gomail.NewMessage()
	m.SetHeader("From", credentials.Email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Восстановление пароля")
	m.SetBody("text/html", "Доброго времени дня, "+email+"!"+
		"<p>Вы запросили восстановление пароля."+
		"<p>Теперь ваш новый пароль:<b>  "+newPassword+"</b>"+
		"<p>Всего хорошего,<br>Ваш сервис восстановления паролей MSA.")

	d := gomail.NewDialer("smtp.gmail.com", 587, credentials.Email, credentials.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return false
	}
	return true
}
