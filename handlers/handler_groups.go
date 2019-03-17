package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"log"
	"net/http"
	"strconv"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	token := r.FormValue("token")

	email, err := data.GetEmailByToken(token)
	if err != nil {
		log.Println("Невозможно получить email по токену:", err)
		return
	}
	user, err := data.GetUserByEmail(email)
	if err != nil {
		log.Println("Невозможно получить пользователя по email:", err)
		return
	}

	groups, err := data.GetGroupsByUserId(user.ID)
	if err != nil {
		log.Println(err, "Ошибка в получении форумов для пользователя "+email)
	}

	if !(groups == nil) {
		response, _ := json_responses.ReturnGroups(groups)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func AddGroup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	token := r.FormValue("token")
	groupName := r.FormValue("group_name")

	email, err := data.GetEmailByToken(token)
	if err != nil {
		log.Println("Невозможно получить email по токену:", err)
		return
	}
	user, err := data.GetUserByEmail(email)
	if err != nil {
		log.Println("Невозможно получить пользователя по email:", err)
		return
	}

	newGroup := data.Group{CreatorUserId: user.ID, Name: groupName}
	if err := newGroup.CreateNewGroup(); err == nil {
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
		log.Println("Невозможно создать новую группу:", err)
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))

	if err := data.DeleteGroupById(id); err == nil {
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
		log.Println("Невозможно удалить группу по id:", err)
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}
