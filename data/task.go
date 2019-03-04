package data

import "log"

type Task struct {
	ID       int
	Name     string
	TaskType int
}

type TaskExtended struct {
	ID       int
	Name     string `json:"name"`
	TaskType int    `json:"task_type"`
	Count    int    `json:"count"`
	Size     int    `json:"size"`
	Min      int    `json:"min"`
	Max      int    `json:"max"`
}

func (task *Task) CreateNewTask() {
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
