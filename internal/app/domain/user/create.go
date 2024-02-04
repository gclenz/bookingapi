package user

import (
	"context"
	"time"
)

type Create struct {
	repository Repository
}

func NewCreate(repository Repository) *Create {
	return &Create{
		repository: repository,
	}
}

func (cu *Create) Execute(
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
	dateOfBirth time.Time,
	ctx context.Context,
) (*User, error) {
	if firstName == "" || lastName == "" || email == "" || phone == "" || document == "" {
		return nil, ErrMissingData
	}
	code := "------"
	user := NewUser(firstName, lastName, email, phone, document, dateOfBirth, code)
	err := cu.repository.Create(user, ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
