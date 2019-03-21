package handlers

import (
	"MSA/data"
	"MSA/json_responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in Registration handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var registrationBody data.RegistrationBody
	json.Unmarshal(bodyBytes, &registrationBody)

	if err := r.ParseForm(); err != nil {
		log.Fatalln("ParseForm() err:", err)
		return
	}

	newUser := data.User{Email: registrationBody.Email, FirstName: registrationBody.FirstName,
		LastName: registrationBody.LastName, Patronymic: registrationBody.Patronymic}

	isAllCorrect, err := data.IsExist(registrationBody.Email)
	if err != nil {
		log.Println(err)
		return
	}

	if isAllCorrect == true {
		newUser.Password = data.GenerateNewPassword()
		if err := newUser.RegisterNewUser(); err == nil {
			res := data.CreateNewUserEmail(newUser)
			var response []byte
			if res {
				response, _ = json_responses.ReturnStatus(true)
			} else {
				response, _ = json_responses.ReturnStatus(false)
			}
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
	log.Println("Processing in PerformLogin handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var authBody data.AuthBody
	json.Unmarshal(bodyBytes, &authBody)

	token := data.GenerateSessionToken()

	if data.IsUserValid(authBody.Email, authBody.Password) {
		user, err := data.GetUserByEmail(authBody.Email)
		if err != nil {
			response, _ := json_responses.ReturnAuthResponse(false, token)
			w.Write(response)
		}
		data.AddNewSession(user.ID, token)
		response, _ := json_responses.ReturnAuthResponse(true, token)
		w.Write(response)
	} else {
		response, _ := json_responses.ReturnAuthResponse(false, token)
		w.Write(response)
	}
}

func ResetPasswordToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in ResetPasswordToken handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var changePasswordBody data.ChangePasswordBody
	json.Unmarshal(bodyBytes, &changePasswordBody)

	email, _ := data.GetEmailByToken(changePasswordBody.Token)

	isExist, err := data.IsExist(email)
	if err != nil {
		log.Println(err)
		return
	}
	if !isExist {
		resp := data.UpdatePassword(email, data.Encrypt(changePasswordBody.Password))
		if resp {
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

func ResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in ResetPasswordEmail handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var emailBody data.EmailBody
	json.Unmarshal(bodyBytes, &emailBody)

	isExist, err := data.IsExist(emailBody.Email)
	if err != nil {
		log.Println(err)
		return
	}
	if !isExist {
		newPassword := data.GenerateNewPassword()
		resp := data.UpdatePassword(emailBody.Email, data.Encrypt(newPassword))
		if resp {
			response, _ := json_responses.ReturnStatus(data.ResetPasswordEmail(emailBody.Email, newPassword))
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

func CheckSession(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing in CheckSession handler")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tokenBody data.TokenBody
	json.Unmarshal(bodyBytes, &tokenBody)

	exist, err := data.IsTokenExist(tokenBody.Token)
	if err != nil {
		log.Println(err)
		return
	}

	email, err := data.GetEmailByToken(tokenBody.Token)
	if err != nil {
		log.Println("Невозможно получить email по токену")
	}

	user, err := data.GetUserByEmail(email)
	if err != nil {
		log.Println("Невозможно получить пользователя по email")
	}

	if exist {
		log.Println("User authorized!")
		response, _ := json_responses.ReturnFio(strings.Title(user.LastName) + " " +
			string([]rune(user.FirstName)[0]) + "." + string([]rune(user.Patronymic)[0]) + ".")
		w.Write(response)
	} else {
		log.Println("User unauthorized!")
		response, _ := json_responses.ReturnStatus(false)
		w.Write(response)
	}
}
