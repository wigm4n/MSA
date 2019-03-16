package data

import (
	"TS/data"
	"log"
)

type User struct {
	ID        int
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

//регистрация нового пользователя
func (user *User) RegisterNewUser() (err error) {
	statement := "INSERT INTO users (email, firstName, lastName, password) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Email, user.FirstName, user.LastName, Encrypt(user.Password)).Scan(&user.ID)
	return
}

//проверка, существует ли пользователь
func IsExist(email string) (exists bool, err error) {
	row, err := db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		log.Println(err)
		return
	}
	exists = !row.Next()
	return
}

//проверка аутентификационных данных пользователя
func IsUserValid(email string, password string) (exists bool) {
	log.Println("in IsUserValid method")
	user, err := GetUserByEmail(email)
	if err != nil {
		log.Println("in IsUserValid err log:", err)
		return false
	}
	if user.Password == Encrypt(password) {
		return true
	}
	return false
}

//получение пользователя по email
func GetUserByEmail(email string) (user User, err error) {
	log.Println("in GetUserByEmail method")
	err = db.QueryRow("SELECT id, email, firstname, lastname, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		log.Println("in GetUserByEmail exception:", err)
		return
	}
	return
}

func GetEmailByToken(token string) (email string, err error) {
	var session data.Session
	err = db.QueryRow("SELECT id, user_id FROM sessions WHERE token = $1", token).
		Scan(&session.Id, &session.UserId)
	if err != nil {
		log.Println("GetEmailByToken exception, cannot get session:", err)
		return
	}
	err = db.QueryRow("SELECT email FROM users WHERE id = $1", session.UserId).
		Scan(&email)
	if err != nil {
		log.Println("GetUserByEmail exception, cannot get email:", err)
		return
	}
	return
}

func UpdatePassword(email string, password string) (isSuccess bool) {
	statement := "UPDATE users SET password = $1 WHERE email = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Query(password, email)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
