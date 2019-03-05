package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetForums(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	email := r.FormValue("email")
	//email = "test3@hse.ru"

	forums := data.GetForumsByUserName(email)

	response, _ := json.Marshal(forums)
	w.Write(response)
}

func GetForum(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	//id = 4

	messages := data.GetMessagesByForum(id)

	response, _ := json.Marshal(messages)
	w.Write(response)
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
