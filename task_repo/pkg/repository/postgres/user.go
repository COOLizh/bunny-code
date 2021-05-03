// Package postgres ...
package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// UserRepositoryPostgresql ...
type UserRepositoryPostgresql struct {
	db *pgxpool.Pool
}

// NewUserRepository ...
func NewUserRepository(db *pgxpool.Pool) *UserRepositoryPostgresql {
	return &UserRepositoryPostgresql{
		db: db,
	}
}

// Create ...
func (r *UserRepositoryPostgresql) Create(ctx context.Context, user *models.User) (newUser models.User, err error) {
	query := "INSERT INTO USERS(login, password) VALUES($1, $2) RETURNING id"

	var id int

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, query, user.Username, user.Password).Scan(&id)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_login_key\" (SQLSTATE 23505)" {
			return newUser, fmt.Errorf("user with username %q already exists", user.Username)
		}

		return
	}

	newUser = models.User{
		ID:       id,
		Username: user.Username,
		Password: user.Password,
	}

	return
}

// GetByUserName ...
func (r *UserRepositoryPostgresql) GetByUserName(ctx context.Context, username string) (user models.User, err error) {
	query := "SELECT id, login, password FROM users WHERE login = $1"

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("user with username %q doesn't exist", username)
		}

		return
	}

	return
}
