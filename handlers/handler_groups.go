package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in GetGroups handler")

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

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in GetAllGroups handler")

	groups, err := data.GetAllGroups()
	if err != nil {
		log.Println(err, "Ошибка в получении всех групп")
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
	log.Println("Processing in GetGroups handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var groupAddBody data.GroupAddBody
	json.Unmarshal(bodyBytes, &groupAddBody)

	email, err := data.GetEmailByToken(groupAddBody.Token)
	if err != nil {
		log.Println("Невозможно получить email по токену:", err)
		return
	}
	user, err := data.GetUserByEmail(email)
	if err != nil {
		log.Println("Невозможно получить пользователя по email:", err)
		return
	}

	newGroup := data.Group{CreatorUserId: user.ID, Name: groupAddBody.GroupName}
	if err := newGroup.CreateNewGroup(); err == nil {
		response, _ := json_responses.ReturnId(newGroup.ID)
		w.Write(response)
	} else {
		log.Println("Невозможно создать новую группу:", err)
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in GetGroups handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var groupIdBody data.GroupIdBody
	json.Unmarshal(bodyBytes, &groupIdBody)

	isExist, err := data.IsTaskExistByGroupId(groupIdBody.Id)
	if err != nil {
		log.Println("Невозможно проверить наличие задания, привязанного к id группы:", err)
		return
	}
	if isExist {
		if err := data.DeleteGroupById(groupIdBody.Id); err == nil {
			response, _ := json_responses.ReturnStatus(true)
			w.Write(response)
		} else {
			log.Println("Невозможно удалить группу по id:", err)
			response, _ := json_responses.ReturnStatus(false)
			w.Write(response)
		}
	} else {
		response, _ := json_responses.ReturnDeleteGroupUnavailable()
		log.Println("Невозможно удалить группу по id, потому что к группе привязано задание:", err)
		w.Write(response)
	}
}
