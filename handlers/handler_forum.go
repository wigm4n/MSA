package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"log"
	"net/http"
	"strconv"
)

func GetForums(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	token := r.FormValue("token")
	email, err := data.GetEmailByToken(token)
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
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	log.Println(id)

	messages, err := data.GetMessagesByForum(id)
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
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	userName := r.FormValue("username")
	text := r.FormValue("text")
	id, _ := strconv.Atoi(r.FormValue("id"))

	//userName = "TESTING_USER"
	//text = "LOCAL_TEST_TEXT_IN_4_TASK"
	//id = 4

	message := data.Message{UserName: userName, Text: text}
	message.CreateNewMessage(id)

	response, _ := json_responses.ReturnStatus(true)
	w.Write(response)
}
