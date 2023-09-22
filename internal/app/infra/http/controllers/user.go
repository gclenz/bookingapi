package controllers

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
	"github.com/gclenz/tinybookingapi/internal/app/infra/http/utils"
)

type CreateUserRequest struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Document    string    `json:"document"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type UserController struct {
	createUser *user.CreateUser
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var u CreateUserRequest
	err := utils.ParseJSON(&u, w, r)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Message, mr.Status)
			return
		}
		slog.Error("CreateUser error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Something went wrong.`))
		return
	}

	ctx := r.Context()
	_, err = uc.createUser.Execute(u.FirstName, u.LastName, u.Email, u.Phone, u.Document, u.DateOfBirth, ctx)
	if err != nil {
		slog.Error("CreateUser error:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Something went wrong. The document, email, and phone are valid for just one account."}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateUserController(createUser *user.CreateUser) *UserController {
	return &UserController{
		createUser: createUser,
	}
}
