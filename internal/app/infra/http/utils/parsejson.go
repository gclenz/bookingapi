package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
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

func ParseJSON(target any, w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(target)
	if err != nil {
		slog.Error(fmt.Sprintf("decoding error: %v", err.Error()))
		w.Header().Add("Content-Type", "application/json")
		var mr *MalformedRequest
		switch {
		case strings.Contains(err.Error(), "json: unknown field"):
			mr = &MalformedRequest{Status: http.StatusUnprocessableEntity, Message: `{"message": "the application could not process the received entity"}`}
		default:
			mr = &MalformedRequest{Status: http.StatusInternalServerError, Message: `{"message": "something went wrong"}`}
		}
		if errors.As(err, &mr) {
			http.Error(w, mr.Message, mr.Status)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			panic(http.ErrAbortHandler)
		}
		http.Error(w, `{"message": "Something went wrong."}`, http.StatusInternalServerError)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		panic(http.ErrAbortHandler)
	}
}
