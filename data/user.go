package data

import (
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
	user, err := GetUserByEmail(email)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if user.Password == Encrypt(password) {
		return true
	}
	return false
}

func ResetPassword(email string) (exists bool) {
	// TODO: сделать сервис email и восстановку пароля
	return true
}

//получение пользователя по email
func GetUserByEmail(email string) (user User, err error) {
	err = db.QueryRow("SELECT id, email, firstname, lastname, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
