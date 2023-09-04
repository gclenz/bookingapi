package user

import (
	"time"

	"github.com/google/uuid"
)

var (
	CustomerRole = "customer"
	StaffRole    = "staff"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Document  string    `json:"document"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
) *User {
	now := time.Now()
	return &User{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Document:  document,
		Role:      CustomerRole,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func BuildUser(
	id string,
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
	role string,
	createdAt time.Time,
	updateAt time.Time,
) *User {
	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Document:  document,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updateAt,
	}
}
