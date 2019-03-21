package data

import (
	"log"
)

type User struct {
	ID         int
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
	Password   string `json:"password"`
}

type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordBody struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type RegistrationBody struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
}

type FioResponse struct {
	Fio    string `json:"fio"`
	Status string `json:"status"`
}

//регистрация нового пользователя
func (user *User) RegisterNewUser() (err error) {
	statement := "INSERT INTO users (email, firstName, lastName, patronymic, password) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Email, user.FirstName, user.LastName, user.Patronymic, Encrypt(user.Password)).Scan(&user.ID)
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
	err = db.QueryRow("SELECT id, email, firstname, lastname, patronymic, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Patronymic, &user.Password)
	if err != nil {
		log.Println("in GetUserByEmail exception:", err)
		return
	}
	return
}

func GetEmailByToken(token string) (email string, err error) {
	var userId int
	err = db.QueryRow("SELECT user_id FROM sessions WHERE token = $1", token).
		Scan(&userId)
	if err != nil {
		log.Println("GetEmailByToken exception, cannot get session:", err)
		return
	}
	err = db.QueryRow("SELECT email FROM users WHERE id = $1", userId).
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
