package data

type Session struct {
	UserID int
}

func (user User) CreateSession() (err error) {
	//Создать сессию в базе
	return
}

func DeleteSessionByUserID(id int) (err error) {
	//"logout" для текущего пользователя
	return
}

func GetCurrentSession() (session Session, err error) {
	// получаем текущую сессию
	return
}

func GetUserById(id int) (user User, err error) {
	// получаем текущего пользователя через сессию
	return
}
