package user

import (
	"context"
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/infra/http/utils"
)

type Authenticate struct {
	repository Repository
	jwtHandler utils.IJWTHandler
}

func NewAuthenticate(repository Repository, jwtHandler utils.IJWTHandler) *Authenticate {
	return &Authenticate{
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (au *Authenticate) Execute(email string, code string, ctx context.Context) (string, error) {
	user, err := au.repository.FindByEmail(email, ctx)
	if err != nil {
		return "", ErrUserNotFound
	}
	if user.Code != code {
		return "", ErrInvalidCode
	}
	if time.Now().After(user.CodeExpiration) {
		return "", ErrExpiredCode
	}
	token, err := au.jwtHandler.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	err = au.repository.UpdateCode(user.Email, "------", time.Now(), ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}
