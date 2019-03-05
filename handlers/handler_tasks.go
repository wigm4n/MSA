package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"log"
	"net/http"
	"strconv"
)

type TaskType int

const (
	Task1 TaskType = 1
	Task2 TaskType = 2
	Task3 TaskType = 3
	Task4 TaskType = 4
	Task5 TaskType = 5
	Task6 TaskType = 6
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	log.Println("Process is in CreateTask method")

	var taskExtended data.TaskExtended

	taskExtended.Name = r.FormValue("TaskName")
	taskExtended.TaskType, _ = strconv.Atoi(r.FormValue("task_type"))
	taskExtended.Count, _ = strconv.Atoi(r.FormValue("count"))
	taskExtended.Size, _ = strconv.Atoi(r.FormValue("size"))
	taskExtended.Size2, _ = strconv.Atoi(r.FormValue("size_2"))
	taskExtended.Size3, _ = strconv.Atoi(r.FormValue("size_3"))
	taskExtended.Min, _ = strconv.Atoi(r.FormValue("min"))
	taskExtended.Max, _ = strconv.Atoi(r.FormValue("max"))
	taskExtended.Alpha, _ = strconv.ParseFloat(r.FormValue("alpha"), 64)

	log.Println("Read task params from user:", taskExtended)

	//===== запись в базу данных
	taskForDB := data.CreateNewTaskObject(taskExtended.Name, taskExtended.TaskType)
	taskForDB.CreateNewTaskInDB()
	//=====

	status := TaskType(taskExtended.TaskType).TaskType(taskExtended)
	log.Println("Task data generated:", status)

	response, err := json_responses.ReturnStatus(status)

	if err = r.ParseForm(); err != nil {
		log.Fatalln("Error in the formation of a response from the server... Error is:", err)
		return
	}
	w.Write(response)
}

func (task TaskType) TaskType(taskExtended data.TaskExtended) bool {
	switch task {
	case Task1:
		return true
	case Task2:
		return true
	case Task3:
		return true
	case Task4:
		return true
	case Task5:
		return true
	case Task6:
		return true
	default:
		return false
	}
}
