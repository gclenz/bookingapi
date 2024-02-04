package user

import (
	"context"
	"time"
)

type Repository interface {
	Create(user *User, ctx context.Context) error
	FindByID(userID string, ctx context.Context) (*User, error)
	FindByEmail(email string, ctx context.Context) (*User, error)
	UpdateCode(email string, code string, codeExpiration time.Time, ctx context.Context) error
}
