package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrCreateUser   = errors.New(
		"something went wrong while trying to create a new user. Please, try again. If the error persists, contact the development team",
	)
	ErrDuplicatedKey = errors.New(
		"there is already an user with some of the provided data",
	)
	ErrMissingData = errors.New(
		"you should provide the following data to create an user: first name, last name, email, phone, document, and date of birth",
	)
)
