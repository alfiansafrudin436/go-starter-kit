package repository

import (
	"context"
	"go-starter-kit/config"

	"github.com/google/uuid"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIDRow, error)
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error
	DeactivateUser(ctx context.Context, id uuid.UUID) error
}

// Repository implements UserRepository
type Repository struct {
	query *Queries
}

// NewRepository initializes a new user Repository using the global DB connection
func NewRepository() *Repository {
	db := config.Application.DB
	return &Repository{
		query: New(db),
	}
}

func (r *Repository) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	return r.query.GetAllUsers(ctx)
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIDRow, error) {
	return r.query.GetUserByID(ctx, id)
}

func (r *Repository) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error {
	return r.query.UpdateUserName(ctx, arg)
}

func (r *Repository) DeactivateUser(ctx context.Context, id uuid.UUID) error {
	return r.query.DeactivateUser(ctx, id)
}
