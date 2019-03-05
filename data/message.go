package data

import (
	"fmt"
	"log"
	"time"
)

type Message struct {
	ID       int
	UserName string `json:"username"`
	Text     string `json:"text"`
}

type Forums struct {
	Forums []Task `json:"tasks"`
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

func GetForumsByUserName(ownerName string) (forums []Task) {
	rows, err := db.Query("SELECT id, name, date FROM tasks WHERE creator_user_id = $1", ownerName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Name, &task.Date)
		forums = append(forums, task)
	}
	return
}

func GetMessagesByForum(forumId int) (messages []Message) {
	rows, err := db.Query("SELECT user_id, text FROM messages WHERE task_id = $1 order by date asc", forumId)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var message Message
		err = rows.Scan(&message.UserName, &message.Text)
		messages = append(messages, message)
	}
	return
}
