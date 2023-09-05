package user

import "context"

type Repository interface {
	Create(user *User, ctx context.Context) error
	FindByID(userID string, ctx context.Context) (*User, error)
}
