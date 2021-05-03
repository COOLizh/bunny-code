// Package repository ...
package repository

import (
	"context"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// UserRepository ...
type UserRepository struct {
	User UserRepositoryInterface
}

// NewUserRepository returns new UserRepository
func NewUserRepository(db UserRepositoryInterface) *UserRepository {
	return &UserRepository{
		User: db,
	}
}

// UserRepositoryInterface represent the user's repository contract
type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.User) (models.User, error)
	GetByUserName(ctx context.Context, username string) (models.User, error)
}
