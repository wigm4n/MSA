package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"MSA/sampling"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
	log.Println("Processing in GetTasksByProfessor handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tokenBody data.TokenBody
	json.Unmarshal(bodyBytes, &tokenBody)

	email, err := data.GetEmailByToken(tokenBody.Token)
	if err != nil {
		log.Println(err, "Ошибка в email по токену")
	}

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
	log.Println("Processing in GetTasksForStudents handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tasksStudentsBody data.TasksStudentsBody
	json.Unmarshal(bodyBytes, &tasksStudentsBody)

	tasks, err := data.GetForumsByGroup(tasksStudentsBody.GroupId)
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

func GetOptionNumberForTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in GetTasksForStudents handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tasksStudentsBody data.TaskIdBody
	json.Unmarshal(bodyBytes, &tasksStudentsBody)

	optionCount, err := data.GetOptionCountByTaskId(tasksStudentsBody.TaskId)
	if err != nil {
		log.Println(err, "Ошибка в получении заданий для пользователя ")
	}

	if optionCount > 0 {
		var options []data.OptionNumber
		for i := 0; i < optionCount; i++ {
			var option data.OptionNumber
			option.Id = i + 1
			option.Name = "Вариант №" + strconv.Itoa(i+1)
			options = append(options, option)
		}
		response, _ := json_responses.ReturnOptions(options)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func DownloadTaskForProf(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in DownloadTaskForProf handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var taskIdBody data.GroupTokenProfBody
	json.Unmarshal(bodyBytes, &taskIdBody)

	//MOCK
	//taskIdBody.Id = 6
	//taskIdBody.Token = "XTPSAIOLSTFOPOPM"

	isExist, _ := data.IsTokenExist(taskIdBody.Token)
	if err != nil {
		log.Println(err)
		return
	}

	if isExist {
		path, err := data.GetProfPathByTaskId(taskIdBody.Id)
		if err != nil {
			log.Println(err, "Ошибка в получении пути для задания")
		}

		dataq, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
		w.Header().Set("Content-Disposition", "attachment; filename=Архив с заданиями.zip")

		http.ServeContent(w, r, "Архив с заданиями.zip", time.Now(), bytes.NewReader(dataq))
	} else {
		log.Println("Невозможно скачать задание с указанным токеном")
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func DownloadTaskForStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in DownloadTaskForStudent handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var downloadTask data.DownloadTaskStudent
	json.Unmarshal(bodyBytes, &downloadTask)

	//MOCK
	//downloadTask.Id = 3
	//downloadTask.OptionNumber = 4

	path, err := data.GetTasksPathByTaskId(downloadTask.Id)
	if err != nil {
		log.Println(err, "Ошибка в получении пути для задания")
	}

	dataq, err := ioutil.ReadFile(path + "/Task-" + strconv.Itoa(downloadTask.OptionNumber) + ".xlsx")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	w.Header().Set("Content-Disposition", "attachment; filename=Вариант"+strconv.Itoa(downloadTask.OptionNumber)+".xlsx")

	http.ServeContent(w, r, "Вариант"+strconv.Itoa(downloadTask.OptionNumber)+".xlsx", time.Now(), bytes.NewReader(dataq))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in DeleteTask handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var taskIdBody data.GroupIdBody
	json.Unmarshal(bodyBytes, &taskIdBody)

	err = data.DeleteAllMessagesByTaskId(taskIdBody.Id)
	if err != nil {
		log.Println(err, "Ошибка в удалении форума по id задания")
	}

	if err := data.DeleteTaskById(taskIdBody.Id); err == nil {
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
		log.Println("Невозможно удалить задание по id:", err)
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in CreateTask handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var taskFront data.TaskFront
	json.Unmarshal(bodyBytes, &taskFront)
	var taskExtended data.TaskExtended
	taskExtended.TaskFieldsList = taskFront.TaskFieldsList
	taskExtended.Name = taskFront.Name
	taskExtended.TaskType = taskFront.TaskType
	taskExtended.Token = taskFront.Token

	email, err := data.GetEmailByToken(taskExtended.Token)
	if err != nil {
		log.Println(err, "Ошибка в получении email по токену")
	}
	taskExtended.Email = email

	user, err := data.GetUserByEmail(taskExtended.Email)
	if err != nil {
		log.Println(err, "Ошибка в получении пользователя из базы данных")
	}

	var status bool
	for i := 0; i < len(taskExtended.TaskFieldsList); i++ {
		groupName, err := data.GetGroupNameByGroupId(taskExtended.TaskFieldsList[i].GroupId)
		if err != nil {
			log.Println(err, "Ошибка в получении имени группы по id", err)
		}
		var pathToArchive, pathToTasks string
		status, pathToArchive, pathToTasks = TaskType(taskExtended.TaskType).TaskType(taskExtended.TaskFieldsList[i],
			i, groupName)
		if !status {
			log.Println("Ошибка в генерации данных")
			response, _ := json_responses.ReturnStatus(status)
			w.Write(response)
			return
		}
		log.Println("Task data generated:", status)

		taskForDB := data.CreateNewTaskObject(taskExtended.Name, taskExtended.TaskFieldsList[i].GroupId,
			taskExtended.TaskFieldsList[i].Count, pathToArchive, pathToTasks)
		err = taskForDB.CreateNewTaskInDB(user.ID)
		if err != nil {
			log.Println(err, "Ошибка в создании задания в базе данных")
		}
		data.CreateWelcomeMessage(taskForDB.ID, taskExtended.Email, taskExtended.Name)
		if err != nil {
			log.Println(err, "Ошибка в создании первого собщения в базе данных")
		}
	}

	response, err := json_responses.ReturnStatus(status)

	if err = r.ParseForm(); err != nil {
		log.Fatalln("Error in the formation of a response from the server... Error is:", err)
		return
	}
	w.Write(response)
}

func (task TaskType) TaskType(taskExtended data.TaskFields, i int, groupName string) (bool, string, string) {
	switch task {
	case Task1:
		return sampling.ReturnTask1(taskExtended, i, groupName)
	case Task2:
		return sampling.ReturnTask2(taskExtended, i, groupName)
	case Task3:
		return sampling.ReturnTask3(taskExtended, i, groupName)
	case Task4:
		return sampling.ReturnTask4(taskExtended, i, groupName)
	case Task5:
		return sampling.ReturnTask5(taskExtended, i, groupName)
	case Task6:
		return sampling.ReturnTask6(taskExtended, i, groupName)
	default:
		return false, "", ""
	}
}
