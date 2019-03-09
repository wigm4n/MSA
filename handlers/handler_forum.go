package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"MSA/testing"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetForums(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	email := r.FormValue("email")
	email = "chyps97@gmail.com"

	if testing.IsTestModeOn() {
		var forums []data.Forum
		forums = append(forums, data.Forum{ID: 1, Name: "Домашнее задание №1", Date: time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC)},
			data.Forum{ID: 2, Name: "Домашнее задание №2", Date: time.Date(2019, 11, 18, 20, 34, 58, 651387237, time.UTC)},
			data.Forum{ID: 3, Name: "Домашнее задание №3", Date: time.Date(2019, 11, 19, 20, 34, 58, 651387237, time.UTC)})
		//если хочешь получить список форумов в ответе
		response, _ := json_responses.ReturnForums(forums)
		//если хочешь получить ошибку в ответе
		//response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	} else {
		user, err := data.GetUserByEmail(email)
		if err != nil {
			log.Println(err, "Ошибка в получении пользователя из базы данных")
		}

		forums, err := data.GetForumsByUserName(user.ID)
		if err != nil {
			log.Println(err, "Ошибка в получении форумов для пользователя "+email)
		}

		if !(forums == nil) {
			response, _ := json.Marshal(forums)
			w.Write(response)
		} else {
			response, _ := json.Marshal(false)
			w.Write(response)
		}
	}
}

func GetForum(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	if testing.IsTestModeOn() {
		var messages []data.Message
		messages = append(messages, data.Message{UserName: "Жукова", Text: "Как идет ваше дз?"},
			data.Message{UserName: "Аноним", Text: "Так себе, тетя"},
			data.Message{UserName: "Ололо3228", Text: "Плюсую к анониму"})
		//если хочешь получить список сообщений в ответе
		response, _ := json_responses.ReturnMessages(messages)
		//если хочешь получить ошибку в ответе
		//response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	} else {
		id, _ := strconv.Atoi(r.FormValue("id"))
		id = 1

		messages := data.GetMessagesByForum(id)

		response, _ := json.Marshal(messages)
		w.Write(response)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	if testing.IsTestModeOn() {
		//поменять параметр на false, если хочешь вернуть ошибку
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
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
}
