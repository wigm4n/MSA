package data

import (
	"log"
	"time"
)

type Task struct {
	ID       int
	Name     string `json:"name"`
	TaskType int
	Date     time.Time `json:"date"`
}

type Forum struct {
	ID   int
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type TaskExtended struct {
	ID            int
	Email         string  `json:"creator_email"`
	Name          string  `json:"task_name"`
	TaskType      int     `json:"task_type"`
	Count         int     `json:"count"`
	Size          int     `json:"size"`
	Size2         int     `json:"size2"`
	Size3         int     `json:"size3"`
	ExpectedValue float64 `json:"expected_value"`
	StdDeviation  float64 `json:"std_deviation"`
	DecimalPlaces int     `json:"decimal_places"`
	Alpha         float64 `json:"alpha"`
}

func CreateNewTaskObject(name string, taskType int) (task Task) {
	task.Name = name
	task.TaskType = taskType
	return
}

func (task *Task) CreateNewTaskInDB(userId int) (err error) {
	statement := "INSERT INTO tasks (creator_user_id, name, task_type, date) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(userId, task.Name, task.TaskType, time.Now()).Scan(&task.ID)
	return
}
