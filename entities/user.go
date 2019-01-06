package entities

type User struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (user *User) RegisterNewUser() (err error) {
	// запрос в базу
	return
}

func IsUserValid(email string, password string) (exists bool) {
	// запрос в базу
	if email == "test@hse.ru" && password == "1234" {
		return true
	} else {
		return false
	}
}

func GetUserByEmail(email string) (user User, err error) {
	// запрос в базу
	user.ID = 1
	user.Email = "test@hse.ru"
	return
}
