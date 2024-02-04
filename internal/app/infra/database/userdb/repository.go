package userdb

import (
	"context"
	"database/sql"
	"log/slog"
	"strings"
	"time"

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
func (ur *UserRepository) Create(usr *user.User, ctx context.Context) error {
	_, err := ur.db.ExecContext(
		ctx,
		InsertUserQuery,
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.Phone,
		&usr.Document,
		&usr.DateOfBirth,
		&usr.Role,
		&usr.CreatedAt,
		&usr.UpdatedAt,
	)

	if err != nil {
		slog.Error("UserRepository(Create) error:", err)
		if strings.Contains(err.Error(), "duplicate key") {
			return user.ErrDuplicatedKey
		}
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
		&user.Code,
		&user.CodeExpiration,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		slog.Error("UserRepository(FindByID) error:", err)
		return nil, err
	}

	return &user, nil
}

// FindByEmail implements user.Repository.
func (ur *UserRepository) FindByEmail(email string, ctx context.Context) (*user.User, error) {
	row := ur.db.QueryRowContext(
		ctx,
		SelectUserByEmailQuery,
		email,
	)
	err := row.Err()
	if err != nil {
		slog.Error("UserRepository(FindByEmail) error:", err)
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
		&user.Code,
		&user.CodeExpiration,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		slog.Error("UserRepository(FindByEmail) error:", err)
		return nil, err
	}

	return &user, nil
}

// UpdateCode implements user.Repository.
func (ur *UserRepository) UpdateCode(email string, code string, codeExpiration time.Time, ctx context.Context) error {
	_, err := ur.db.ExecContext(
		ctx,
		UpdateUserCodeQuery,
		code,
		codeExpiration,
		email,
	)

	if err != nil {
		slog.Error("UserRepository(UpdateCode) error:", err)
		return err
	}

	return nil
}
