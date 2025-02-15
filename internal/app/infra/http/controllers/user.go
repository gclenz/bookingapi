package controllers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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

type AuthenticateUserRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type RequestAuthenticationCodeRequest struct {
	Email string `json:"email"`
}

type UserController struct {
	createUser                *user.Create
	authenticate              *user.Authenticate
	requestAuthenticationCode *user.RequestAuthenticationCode
	jwtHandler                utils.IJWTHandler
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var u CreateUserRequest

	utils.ParseJSON(&u, w, r)

	ctx := r.Context()
	_, err := uc.createUser.Execute(u.FirstName, u.LastName, u.Email, u.Phone, u.Document, u.DateOfBirth, ctx)
	if err != nil {
		slog.Error("UserController(CreateUser)", "error", err)
		switch {
		case errors.Is(err, user.ErrMissingData),
			errors.Is(err, user.ErrCreateUser),
			errors.Is(err, user.ErrDuplicatedKey):
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"message": "%s."}`, err)))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Something went wrong. The document, email, and phone are valid for just one account."}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var u AuthenticateUserRequest
	utils.ParseJSON(&u, w, r)

	ctx := r.Context()
	token, err := uc.authenticate.Execute(u.Email, u.Code, ctx)
	if err != nil {
		slog.Error("UserController(AuthenticateUser)", "error", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   os.Getenv("GO_ENV") == "production",
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Name:     "access_token",
		Value:    fmt.Sprintf("Bearer %s", token),
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) RequestAuthenticationCode(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	var rac RequestAuthenticationCodeRequest
	utils.ParseJSON(&rac, w, r)

	ctx := r.Context()
	err := uc.requestAuthenticationCode.Execute(rac.Email, ctx)
	if err != nil {
		slog.Error("UserController(RequestAuthenticationCode)", "error", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewUserController(
	createUser *user.Create,
	requestAuthenticationCode *user.RequestAuthenticationCode,
	authenticate *user.Authenticate,
	jwtHandler utils.IJWTHandler,
) *UserController {
	return &UserController{
		createUser:                createUser,
		authenticate:              authenticate,
		requestAuthenticationCode: requestAuthenticationCode,
		jwtHandler:                jwtHandler,
	}
}
