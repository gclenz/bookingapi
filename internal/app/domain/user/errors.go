package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrCreateUser   = errors.New(
		"something went wrong while trying to create a new user. Please, try again. If the error persists, contact the development team",
	)
)
