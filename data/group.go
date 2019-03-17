package data

import "log"

type Group struct {
	CreatorUserId int    `json:"creator_user_id"`
	Name          string `json:"name"`
}

//получение пользователя по email
func GetGroupNameByGroupId(id int) (name string, err error) {
	err = db.QueryRow("SELECT name FROM groups WHERE creator_user_id = $1", id).
		Scan(&name)
	if err != nil {
		log.Println("GetGroupNameByGroupId exception:", err)
		return
	}
	return
}
