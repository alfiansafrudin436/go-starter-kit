package repository

import (
	"context"
	"database/sql"
	"go-starter-kit/config"

	"github.com/google/uuid"
)

// AuthRepository defines the interface for auth data operations
type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*GetUserByIDRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserResetToken(ctx context.Context, arg UpdateUserResetTokenParams) error
	GetUserByResetToken(ctx context.Context, token sql.NullString) (*GetUserByResetTokenRow, error)
}

// Repository implements AuthRepository
type Repository struct {
	query *Queries
}

// NewRepository initializes a new auth Repository using the global DB connection
func NewRepository() *Repository {
	db := config.Application.DB
	return &Repository{
		query: New(db),
	}
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return r.query.GetUserByEmail(ctx, email)
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*GetUserByIDRow, error) {
	return r.query.GetUserByID(ctx, id)
}

func (r *Repository) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	return r.query.CreateUser(ctx, arg)
}

func (r *Repository) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	return r.query.UpdateUserPassword(ctx, arg)
}

func (r *Repository) UpdateUserResetToken(ctx context.Context, arg UpdateUserResetTokenParams) error {
	return r.query.UpdateUserResetToken(ctx, arg)
}

func (r *Repository) GetUserByResetToken(ctx context.Context, token sql.NullString) (*GetUserByResetTokenRow, error) {
	return r.query.GetUserByResetToken(ctx, token)
}
