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

type TaskExtended struct {
	ID       int
	Name     string  `json:"name"`
	TaskType int     `json:"task_type"`
	Count    int     `json:"count"`
	Size     int     `json:"size"`
	Size2    int     `json:"size2"`
	Size3    int     `json:"size3"`
	Min      int     `json:"min"`
	Max      int     `json:"max"`
	Alpha    float64 `json:"alpha"`
}

func CreateNewTaskObject(name string, taskType int) (task Task) {
	task.Name = name
	task.TaskType = taskType
	return
}

func (task *Task) CreateNewTaskInDB() {
	statement := "INSERT INTO tasks (name, task_type) VALUES ($1, $2) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(task.Name, task.TaskType).Scan(&task.ID)
	return
}
