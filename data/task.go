package data

import (
	"log"
	"time"
)

type Task struct {
	ID            int
	Name          string    `json:"name"`
	GroupId       int       `json:"group_id"`
	Count         int       `json:"count"`
	PathToArchive string    `json:"path_to_prof"`
	PathToTasks   string    `json:"path_to_tasks"`
	Date          time.Time `json:"date"`
}

type Forum struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	GroupName string    `json:"group_name"`
	Date      time.Time `json:"date"`
}

type TasksStudentsBody struct {
	GroupId int `json:"group_id"`
}

type TaskIdBody struct {
	TaskId int `json:"task_id"`
}

type OptionNumber struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OptionNumberBody struct {
	Options []OptionNumber `json:"options"`
}

type DownloadTaskStudent struct {
	Id           int `json:"task_id"`
	OptionNumber int `json:"id"`
}

type TaskFront struct {
	Name           string       `json:"task_name"`
	TaskType       int          `json:"task_type"`
	Token          string       `json:"token"`
	TaskFieldsList []TaskFields `json:"task_fields"`
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

func CreateNewTaskObject(name string, groupId int, count int, pathToArchive string, pathToTasks string) (task Task) {
	task.Name = name
	task.GroupId = groupId
	task.Count = count
	task.PathToArchive = pathToArchive
	task.PathToTasks = pathToTasks
	return
}

func (task *Task) CreateNewTaskInDB(userId int) (err error) {
	statement := "INSERT INTO tasks (creator_user_id, name, group_id, option_count, path_to_archive, path_to_tasks," +
		" date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(userId, task.Name, task.GroupId, task.Count, task.PathToArchive, task.PathToTasks,
		time.Now()).Scan(&task.ID)
	return
}

//проверка, существует ли пользователь
func IsTaskExistByGroupId(id int) (exists bool, err error) {
	row, err := db.Query("SELECT * FROM tasks WHERE group_id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}
	exists = !row.Next()
	return
}

func GetOptionCountByTaskId(id int) (count int, err error) {
	err = db.QueryRow("SELECT option_count FROM tasks WHERE id = $1", id).Scan(&count)
	if err != nil {
		log.Println("in GetOptionCountByTaskId exception:", err)
		return
	}
	return
}

func GetProfPathByTaskId(id int) (path string, err error) {
	err = db.QueryRow("SELECT path_to_archive FROM tasks WHERE id = $1", id).Scan(&path)
	if err != nil {
		log.Println("in GetProfPathByTaskId exception:", err)
		return
	}
	return
}

func GetTasksPathByTaskId(id int) (path string, err error) {
	err = db.QueryRow("SELECT path_to_tasks FROM tasks WHERE id = $1", id).Scan(&path)
	if err != nil {
		log.Println("in GetTasksPathByTaskId exception:", err)
		return
	}
	return
}

func DeleteTaskById(id int) (err error) {
	statement := "DELETE FROM tasks WHERE id = $1"
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
