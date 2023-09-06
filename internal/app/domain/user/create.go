package user

import "context"

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
	ctx context.Context,
) (*User, error) {
	user := NewUser(firstName, lastName, email, phone, document)
	err := cu.repository.Create(user, ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
