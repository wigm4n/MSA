package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetForums(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in GetForums handler")

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

	forums, err := data.GetForumsByUserName(user.ID)
	if err != nil {
		log.Println(err, "Ошибка в получении форумов для пользователя "+email)
	}

	if !(forums == nil) {
		response, _ := json_responses.ReturnForums(forums)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func GetForum(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in GetForum handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var forumIdBody data.ForumIdBody
	json.Unmarshal(bodyBytes, &forumIdBody)

	messages, err := data.GetMessagesByForum(forumIdBody.Id)
	if err != nil {
		log.Println(err, "Ошибка в получении сообщений для форума")
	}

	if !(messages == nil) {
		response, _ := json_responses.ReturnMessages(messages)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in SendMessage handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var message data.Message
	json.Unmarshal(bodyBytes, &message)

	err = message.CreateNewMessage()
	if err != nil {
		log.Println(err, "Ошибка в отправки сообщения")
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	} else {
		response, _ := json_responses.MessageId(message.Id)
		w.Write(response)
	}
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in DeleteMessage handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var messageIdBody data.GroupIdBody
	json.Unmarshal(bodyBytes, &messageIdBody)

	if err := data.DeleteMessage(messageIdBody.Id); err == nil {
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
		log.Println("Невозможно удалить сообщение по id:", err)
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}
