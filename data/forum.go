package data

import (
	"fmt"
	"log"
	"time"
)

type Message struct {
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
	stmt.Query(taskId, message.UserName, message.Text, time.Now())
	return
}

func CreateWelcomeMessage(taskId int, userId string, taskName string) (err error) {
	statement := "INSERT INTO messages (task_id, user_id, text, date) VALUES ($1, $2, $3, $4)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	stmt.Query(taskId, userId, "Этот форум создан для обсуждения задания \""+taskName+"\".", time.Now())
	return
}

func GetForumsByUserName(ownerNameId int) (forums []Forum, err error) {
	rows, err := db.Query("SELECT id, name, group_id, date FROM tasks WHERE creator_user_id = $1", ownerNameId)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var task Forum
		var groupId int
		err = rows.Scan(&task.ID, &task.Name, &groupId, &task.Date)
		groupName, err := GetGroupNameByGroupId(groupId)
		if err != nil {
			log.Println("GetForumsByUserName exception, err:", err)
			return
		}
		task.GroupName = groupName
		forums = append(forums, task)
	}
	return
}

func GetMessagesByForum(forumId int) (messages []Message, err error) {
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
	if err != nil {
		log.Println("GetMessagesByForum exception, err:", err)
		return
	}
	return
}
