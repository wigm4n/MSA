package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"MSA/sampling"
	"MSA/testing"
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

	if testing.IsTestModeOn() {
		//поменять параметр на false, если хочешь вернуть ошибку
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
		var taskExtended data.TaskExtended

		taskExtended.Email = r.FormValue("creator_email")
		taskExtended.Name = r.FormValue("task_name")
		taskExtended.TaskType, _ = strconv.Atoi(r.FormValue("task_type"))
		taskExtended.Count, _ = strconv.Atoi(r.FormValue("count"))
		taskExtended.Alpha, _ = strconv.ParseFloat(r.FormValue("alpha"), 64)
		taskExtended.Size, _ = strconv.Atoi(r.FormValue("size"))
		taskExtended.Size2, _ = strconv.Atoi(r.FormValue("size_2"))
		taskExtended.Size3, _ = strconv.Atoi(r.FormValue("size_3"))
		taskExtended.ExpectedValue, _ = strconv.ParseFloat(r.FormValue("expected_value"), 64)
		taskExtended.StdDeviation, _ = strconv.ParseFloat(r.FormValue("std_deviation"), 64)
		taskExtended.DecimalPlaces, _ = strconv.Atoi(r.FormValue("decimal_places"))

		log.Println("Read task params from user:", taskExtended)

		taskExtended.Email = "chyps97@gmail.com"
		taskExtended.Name = "Домашнее задание №1"
		taskExtended.TaskType = 1
		taskExtended.Count = 10
		taskExtended.Alpha = 0.05
		taskExtended.Size = 10
		//taskExtended.Size2 = 10
		//taskExtended.Size3 = 10
		taskExtended.ExpectedValue = 12
		taskExtended.StdDeviation = 0.7
		taskExtended.DecimalPlaces = 2

		//===== запись в базу данных
		taskForDB := data.CreateNewTaskObject(taskExtended.Name, taskExtended.TaskType)
		user, err := data.GetUserByEmail(taskExtended.Email)
		if err != nil {
			log.Println(err, "Ошибка в получении пользователя из базы данных")
		}
		err = taskForDB.CreateNewTaskInDB(user.ID)
		if err != nil {
			log.Println(err, "Ошибка в создании задания в базе данных")
		}
		data.CreateWelcomeMessage(taskForDB.ID, taskExtended.Email, taskExtended.Name)
		if err != nil {
			log.Println(err, "Ошибка в создании первого собщения в базе данных")
		}
		//=====

		//status := TaskType(taskExtended.TaskType).TaskType(taskExtended)
		//log.Println("Task data generated:", status)

		//response, err := json_responses.ReturnStatus(status)
		response, err := json_responses.ReturnStatus(true)

		if err = r.ParseForm(); err != nil {
			log.Fatalln("Error in the formation of a response from the server... Error is:", err)
			return
		}
		w.Write(response)
	}
}

func (task TaskType) TaskType(taskExtended data.TaskExtended) bool {
	switch task {
	case Task1:
		return sampling.ReturnTask1(taskExtended)
	case Task2:
		return sampling.ReturnTask2(taskExtended)
	case Task3:
		return sampling.ReturnTask3(taskExtended)
	case Task4:
		return sampling.ReturnTask4(taskExtended)
	case Task5:
		return sampling.ReturnTask5(taskExtended)
	case Task6:
		return sampling.ReturnTask6(taskExtended)
	default:
		return false
	}
}
