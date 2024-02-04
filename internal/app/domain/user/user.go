package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	CustomerRole = "customer"
	StaffRole    = "staff"
)

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Document       string    `json:"document"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	Role           string    `json:"role"`
	Code           string    `json:"code"`
	CodeExpiration time.Time `json:"codeExpiration"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func NewUser(
	firstName string,
	lastName string,
	email string,
	phone string,
	document string,
	dateOfBirth time.Time,
	code string,
) *User {
	now := time.Now()
	expiration := now.Add(time.Minute * 30)
	return &User{
		ID:             uuid.NewString(),
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Phone:          phone,
		Document:       document,
		DateOfBirth:    dateOfBirth,
		Role:           CustomerRole,
		Code:           code,
		CodeExpiration: expiration,
		CreatedAt:      now,
		UpdatedAt:      now,
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
	code string,
	codeExpiration time.Time,
	createdAt time.Time,
	updateAt time.Time,
) *User {
	return &User{
		ID:             id,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Phone:          phone,
		Document:       document,
		DateOfBirth:    dateOfBirth,
		Role:           role,
		Code:           code,
		CodeExpiration: codeExpiration,
		CreatedAt:      createdAt,
		UpdatedAt:      updateAt,
	}
}
