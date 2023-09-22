package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
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

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var u CreateUserRequest
	err := dec.Decode(&u)
	if err != nil {
		slog.Error("CreateUser error:", err)
		w.WriteHeader(http.StatusInternalServerError)
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
