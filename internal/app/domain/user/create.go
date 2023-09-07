package user

import (
	"context"
	"time"
)

type CreateUser struct {
	repository Repository
}

func NewCreateUser(repository Repository) *CreateUser {
	return &CreateUser{
		repository: repository,
	}
}

func (cu *CreateUser) Execute(
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
	dateOfBirth time.Time,
	password string,
	ctx context.Context,
) (*User, error) {
	user := NewUser(firstName, lastName, email, phone, document, dateOfBirth, password)
	err := cu.repository.Create(user, ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
