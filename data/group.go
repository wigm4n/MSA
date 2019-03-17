package data

import (
	"fmt"
	"log"
)

type Group struct {
	ID            int    `json:"id"`
	CreatorUserId int    `json:"creator_user_id"`
	Name          string `json:"name"`
}

//получение имени группы по email
func GetGroupNameByGroupId(id int) (name string, err error) {
	err = db.QueryRow("SELECT name FROM groups WHERE id = $1", id).
		Scan(&name)
	if err != nil {
		log.Println("GetGroupNameByGroupId exception:", err)
		return
	}
	return
}

func GetGroupsByUserId(userId int) (groups []Group, err error) {
	rows, err := db.Query("SELECT id, name FROM groups WHERE creator_user_id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var group Group
		err = rows.Scan(&group.ID, &group.Name)
		if err != nil {
			log.Println("GetGroupsByUserId exception, err:", err)
		}
		groups = append(groups, group)
	}
	return
}

func (group *Group) CreateNewGroup() (err error) {
	statement := "INSERT INTO groups (creator_user_id, name) VALUES ($1, $2) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(group.CreatorUserId, group.Name).Scan(&group.ID)
	return
}

func DeleteGroupById(id int) (err error) {
	statement := "DELETE FROM groups WHERE id = $1"
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
