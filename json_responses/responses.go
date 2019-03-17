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
	Forums []Output `json:"forums"`
}

type ResponseTasks struct {
	Forums []Output `json:"tasks"`
}

type ResponseMessages struct {
	Messages []data.Message `json:"messages"`
}

type ResponseGroups struct {
	Groups []data.Group `json:"groups"`
}

type Output struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"group_name"`
	Date      string `json:"date"`
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

func ReturnTasks(tasks []data.Forum) (response []byte, err error) {
	if tasks != nil || len(tasks) != 0 {
		var forumsOutput []Output
		for i := 0; i < len(tasks); i++ {
			one := Output{ID: tasks[i].ID, Name: tasks[i].Name, GroupName: tasks[i].GroupName, Date: tasks[i].Date.Format("02.01.2006")}
			forumsOutput = append(forumsOutput, one)
		}
		var responseForums ResponseTasks
		responseForums.Forums = forumsOutput
		return json.Marshal(responseForums)
	} else {
		var responseStatus ResponseStatus
		responseStatus.Status = "error"
		return json.Marshal(responseStatus)
	}
}

func ReturnForums(forums []data.Forum) (response []byte, err error) {
	if forums != nil || len(forums) != 0 {
		var forumsOutput []Output
		for i := 0; i < len(forums); i++ {
			one := Output{ID: forums[i].ID, Name: forums[i].Name, GroupName: forums[i].GroupName, Date: forums[i].Date.Format("02.01.2006")}
			forumsOutput = append(forumsOutput, one)
		}
		var responseForums ResponseForums
		responseForums.Forums = forumsOutput
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

func ReturnGroups(groups []data.Group) (response []byte, err error) {
	if groups == nil || len(groups) != 0 {
		var responseGroups ResponseGroups
		responseGroups.Groups = groups
		return json.Marshal(responseGroups)
	} else {
		var responseStatus ResponseStatus
		responseStatus.Status = "error"
		return json.Marshal(responseStatus)
	}
}
