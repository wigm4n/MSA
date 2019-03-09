package data

import "log"

type Session struct {
	ID     int
	UserId int
	Token  string
}

func AddNewSession(userId int, token string) (err error) {
	isExist, err := IsTokenUserWithToken(userId)
	if err != nil {
		log.Println(err)
		return
	}
	if isExist {
		UpdateSession(userId, token)
	} else {
		CreateNewSession(userId, token)
	}
	return
}

func CreateNewSession(userId int, token string) (err error) {
	statement := "INSERT INTO sessions (user_id, token) VALUES ($1, $2)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Query(userId, token)
	return
}

//проверка, существует ли пользователь с любым токеном
func IsTokenUserWithToken(userId int) (exists bool, err error) {
	row, err := db.Query("SELECT * FROM sessions WHERE user_id = $1", userId)
	if err != nil {
		log.Println(err)
		return
	}
	exists = row.Next()
	return
}

//проверка, существует ли токен
func IsTokenExist(token string) (exists bool, err error) {
	row, err := db.Query("SELECT * FROM sessions WHERE token = $1", token)
	if err != nil {
		log.Println(err)
		return
	}
	exists = row.Next()
	return
}

//удаление сессии по id пользователя
func DeleteSessionByUserID(id int) (err error) {
	statement := "DELETE FROM sessions WHERE user_id = $1"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return
}

func UpdateSession(userId int, token string) (isSuccess bool) {
	statement := "UPDATE sessions SET token = $1 WHERE user_id = $2"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Query(token, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
