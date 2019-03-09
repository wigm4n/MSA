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
		newUser.Password = data.GenerateNewPassword()
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

	email = "chyps97@gmail.com"
	password = "1234"

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
	email = "chyps97@gmail.com"

	isExist, err := data.IsExist(email)
	if err != nil {
		log.Println(err)
		return
	}
	if !isExist {
		newPassword := data.GenerateNewPassword()
		resp := data.UpdatePassword(email, data.Encrypt(newPassword))
		if resp {
			response, _ := json_responses.ReturnStatus(data.SendEmail(email, newPassword))
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
