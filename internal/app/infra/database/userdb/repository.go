package userdb

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gclenz/tinybookingapi/internal/app/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &UserRepository{
		db: db,
	}
}

// Create implements user.Repository.
func (ur *UserRepository) Create(user *user.User, ctx context.Context) error {
	_, err := ur.db.ExecContext(
		ctx,
		InsertUserQuery,
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Document,
		&user.DateOfBirth,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		slog.Error("UserRepository(Create) error:", err)
		return err
	}

	return nil
}

// FindByID implements user.Repository.
func (ur *UserRepository) FindByID(userID string, ctx context.Context) (*user.User, error) {
	row := ur.db.QueryRowContext(
		ctx,
		SelectUserByIDQuery,
		userID,
	)
	err := row.Err()
	if err != nil {
		slog.Error("UserRepository(FindByID) error:", err)
		return nil, err
	}

	var user user.User
	err = row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Document,
		&user.DateOfBirth,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		slog.Error("UserRepository(FindByID) error:", err)
		return nil, err
	}

	return &user, nil
}
