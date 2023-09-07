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
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Document    string    `json:"document"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Role        string    `json:"role"`
	Password    string    `json:"password,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewUser(
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
	dateOfBirth time.Time,
	password string,
) *User {
	now := time.Now()
	return &User{
		ID:          uuid.NewString(),
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Phone:       phone,
		Document:    document,
		DateOfBirth: dateOfBirth,
		Role:        CustomerRole,
		Password:    password,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func BuildUser(
	id string,
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
	dateOfBirth time.Time,
	role string,
	password string,
	createdAt time.Time,
	updateAt time.Time,
) *User {
	return &User{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Phone:       phone,
		Document:    document,
		DateOfBirth: dateOfBirth,
		Role:        role,
		Password:    password,
		CreatedAt:   createdAt,
		UpdatedAt:   updateAt,
	}
}
