package data

import (
	"log"
	"time"
)

type Task struct {
	ID      int
	Name    string    `json:"name"`
	GroupId int       `json:"group_id"`
	Count   int       `json:"count"`
	Date    time.Time `json:"date"`
}

type Forum struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	GroupName string    `json:"group_name"`
	Date      time.Time `json:"date"`
}

type TaskExtended struct {
	ID             int
	Email          string
	Name           string       `json:"task_name"`
	TaskType       int          `json:"task_type"`
	Token          string       `json:"token"`
	TaskFieldsList []TaskFields `json:"task_fields"`
}

type TaskFields struct {
	GroupId       int     `json:"group_id"`
	Count         int     `json:"count"`
	Size          int     `json:"size"`
	Size2         int     `json:"size_2"`
	Size3         int     `json:"size_3"`
	ExpectedValue float64 `json:"expected_value"`
	StdDeviation  float64 `json:"std_deviation"`
	DecimalPlaces int     `json:"decimal_places"`
}

func CreateNewTaskObject(name string, groupId int, count int) (task Task) {
	task.Name = name
	task.GroupId = groupId
	task.Count = count
	return
}

func (task *Task) CreateNewTaskInDB(userId int) (err error) {
	statement := "INSERT INTO tasks (creator_user_id, name, group_id, option_count, date) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(userId, task.Name, task.GroupId, task.Count, time.Now()).Scan(&task.ID)
	return
}
