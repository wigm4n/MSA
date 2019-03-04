package data

import (
	"log"
	"time"
)

type Message struct {
	ID       int
	UserName string `json:"username"`
	Text     string `json:"text"`
}

func (message *Message) CreateNewMessage(taskId int) {
	statement := "INSERT INTO messages (task_id, user_id, text, date) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(taskId, message.UserName, message.Text, time.Now()).Scan(&message.ID)
	return
}
