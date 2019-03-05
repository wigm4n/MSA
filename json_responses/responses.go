package json_responses

import (
	"encoding/json"
)

type ResponseStatus struct {
	Status string
}

func ReturnStatus(state bool) (response []byte, err error) {
	var responseStatus ResponseStatus
	if state {
		responseStatus.Status = "success"
	} else {
		responseStatus.Status = "error"
	}
	return json.Marshal(responseStatus)
}
