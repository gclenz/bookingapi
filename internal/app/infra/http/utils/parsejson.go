package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	Status  int
	Message string
}

func (mr *MalformedRequest) Error() string {
	return mr.Message
}

func ParseJSON(target any, w http.ResponseWriter, r *http.Request) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(target)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "json: unknown field"):
			return &MalformedRequest{Status: http.StatusUnprocessableEntity, Message: `{"message": "the application could not process the received entity"}`}

		default:
			return &MalformedRequest{Status: http.StatusInternalServerError, Message: `{"message": "something went wrong"}`}
		}
	}

	return nil
}
