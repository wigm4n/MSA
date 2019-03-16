package json_responses

import (
	"MSA/data"
	"encoding/json"
	"log"
)

type ResponseStatus struct {
	Status string `json:"status"`
}

type ResponseSuccessAuth struct {
	Status string `json:"status"`
	Key    string `json:"key"`
}

type ResponseForums struct {
	Forums []data.Forum `json:"forums"`
}

type ResponseMessages struct {
	Messages []data.Message `json:"messages"`
}

func ReturnStatus(state bool) (response []byte, err error) {
	var responseStatus ResponseStatus
	if state {
		responseStatus.Status = "success"
	} else {
		responseStatus.Status = "error"
	}
	log.Println("Json:", responseStatus)
	return json.Marshal(responseStatus)
}

func ReturnAuthResponse(state bool, token string) (response []byte, err error) {
	if state {
		var responseSuccessAuth ResponseSuccessAuth
		responseSuccessAuth.Status = "success"
		responseSuccessAuth.Key = token
		return json.Marshal(responseSuccessAuth)
	} else {
		var responseStatus ResponseStatus
		responseStatus.Status = "error"
		return json.Marshal(responseStatus)
	}
}

func ReturnTasks(forums []data.Forum) (response []byte, err error) {
	if forums == nil || len(forums) != 0 {
		var responseForums ResponseForums
		responseForums.Forums = forums
		return json.Marshal(responseForums)
	} else {
		var responseStatus ResponseStatus
		responseStatus.Status = "error"
		return json.Marshal(responseStatus)
	}
}

func ReturnForums(forums []data.Forum) (response []byte, err error) {
	if forums == nil || len(forums) != 0 {
		var responseForums ResponseForums
		responseForums.Forums = forums
		return json.Marshal(responseForums)
	} else {
		var responseStatus ResponseStatus
		responseStatus.Status = "error"
		return json.Marshal(responseStatus)
	}
}

func ReturnMessages(messages []data.Message) (response []byte, err error) {
	if messages == nil || len(messages) != 0 {
		var responseMessages ResponseMessages
		responseMessages.Messages = messages
		return json.Marshal(responseMessages)
	} else {
		var responseStatus ResponseStatus
		responseStatus.Status = "error"
		return json.Marshal(responseStatus)
	}
}
