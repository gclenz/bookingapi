package user

type CreateUser struct {
	repository Repository
}

func (cu *CreateUser) Execute(firstName string, lastName string, email string, phone string, document string) (*User, error) {
	user := NewUser(firstName, lastName, email, phone, document)
	err := cu.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
