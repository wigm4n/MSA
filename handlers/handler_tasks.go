package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"MSA/sampling"
	"encoding/json"
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

func GetTasksByProfessor(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	log.Println("Process is in GetTasksByProf method")
	token := r.FormValue("token")
	email, err := data.GetEmailByToken(token)
	if err != nil {
		log.Println(err, "Ошибка в email по токену")
	}

	email = "test"

	user, err := data.GetUserByEmail(email)
	if err != nil {
		log.Println(err, "Ошибка в получении пользователя из базы данных")
	}

	tasks, err := data.GetForumsByUserName(user.ID)
	if err != nil {
		log.Println(err, "Ошибка в получении заданий для пользователя "+email)
	}

	if !(tasks == nil) {
		response, _ := json_responses.ReturnTasks(tasks)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func GetTasksForStudents(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	log.Println("Process is in GetTasksForStudents method")

	groupId, _ := strconv.Atoi(r.FormValue("group_id"))
	optionNumber := r.FormValue("option_number")
	log.Println(optionNumber)

	groupId = 1

	tasks, err := data.GetForumsByGroup(groupId)
	if err != nil {
		log.Println(err, "Ошибка в получении заданий для пользователя ")
	}

	if !(tasks == nil) {
		response, _ := json_responses.ReturnTasks(tasks)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	log.Println("Process is in CreateTask method")

	var taskExtended data.TaskExtended

	taskExtended.Name = r.FormValue("task_name")
	taskExtended.TaskType, _ = strconv.Atoi(r.FormValue("task_type"))
	taskExtended.Token = r.FormValue("token")

	mockString := "[{" +
		"\"group_id\": 1," +
		"\"count\": 20," +
		"\"size\": 50," +
		"\"size_2\": 50," +
		"\"size_3\": 50," +
		"\"expected_value\": 1," +
		"\"std_deviation\": 20," +
		"\"decimal_places\": 1},{" +
		"\"group_id\": 2," +
		"\"count\": 240," +
		"\"size\": 5," +
		"\"size_2\": 5," +
		"\"size_3\": 5," +
		"\"expected_value\": 1.1," +
		"\"std_deviation\": 2.1," +
		"\"decimal_places\": 6}]"
	taskExtended.Name = "домашка номер 1"

	var fields []data.TaskFields
	//json.Unmarshal([]byte(r.FormValue("task_fields")), &fields)
	json.Unmarshal([]byte(mockString), &fields)
	taskExtended.TaskFieldsList = fields

	email, err := data.GetEmailByToken(taskExtended.Token)
	if err != nil {
		log.Println(err, "Ошибка в получении email по токену")
	}
	taskExtended.Email = email

	log.Println("Read task params from user:", taskExtended)

	//===== запись в базу данных
	//user, err := data.GetUserByEmail(taskExtended.Email)
	if err != nil {
		log.Println(err, "Ошибка в получении пользователя из базы данных")
	}

	var status bool
	for i := 0; i < len(taskExtended.TaskFieldsList); i++ {
		/*taskForDB := data.CreateNewTaskObject(taskExtended.Name, taskExtended.TaskFieldsList[i].GroupId,
			taskExtended.TaskFieldsList[i].Count)
		err = taskForDB.CreateNewTaskInDB(user.ID)
		if err != nil {
			log.Println(err, "Ошибка в создании задания в базе данных")
		}
		data.CreateWelcomeMessage(taskForDB.ID, taskExtended.Email, taskExtended.Name)
		if err != nil {
			log.Println(err, "Ошибка в создании первого собщения в базе данных")
		}*/

		//ГЕНЕРАЦИЯ ЗАДАНИЙ ВКЛЮЧЕНА
		status = TaskType(taskExtended.TaskType).TaskType(taskExtended.TaskFieldsList[i])
		if !status {
			log.Println("Ошибка в генерации данных")
			response, _ := json_responses.ReturnStatus(status)
			w.Write(response)
			return
		}
		log.Println("Task data generated:", status)

	}
	//=====

	response, err := json_responses.ReturnStatus(status)

	//ГЕНЕРАЦИЯ ЗАДАНИЙ ВЫКЛЮЧЕНА
	//response, err := json_responses.ReturnStatus(true)

	if err = r.ParseForm(); err != nil {
		log.Fatalln("Error in the formation of a response from the server... Error is:", err)
		return
	}
	w.Write(response)
}

func (task TaskType) TaskType(taskExtended data.TaskFields) bool {
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
