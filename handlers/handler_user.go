package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"log"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	email := r.FormValue("email")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")

	newUser := data.User{Email: email, FirstName: firstName, LastName: lastName}

	isAllCorrect, err := data.IsExist(email)
	if err != nil {
		log.Println(err)
		return
	}

	if isAllCorrect == true {
		newUser.Password, _ = data.GenerateNewPassword()
		if err := newUser.RegisterNewUser(); err == nil {
			response, _ := json_responses.ReturnStatus(true)
			w.Write(response)
		} else {
			response, _ := json_responses.ReturnStatus(false)
			w.Write(response)
		}
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func PerformLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	if data.IsUserValid(email, password) {
		response, _ := json_responses.ReturnStatus(true)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}
	email := r.FormValue("email")
	status := data.ResetPassword(email)
	response, _ := json_responses.ReturnStatus(status)
	w.Write(response)
}
