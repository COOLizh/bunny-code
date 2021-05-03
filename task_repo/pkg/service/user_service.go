// Package service ...
package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/repository"
	"gitlab.com/greenteam1/task_repo/pkg/token"
)

// UserService ...
type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

// NewUserService return new UserService
func NewUserService(ur repository.UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

// Register creates user if not exists
func (us *UserService) Register(ctx context.Context, user *models.User) (newUser models.User, err error) {
	err = user.Prepare()
	if err != nil {
		err = fmt.Errorf("invalid user data")
		return
	}

	newUser, err = us.UserRepository.Create(ctx, user)
	if err != nil {
		return
	}

	return
}

// Login ...
func (us *UserService) Login(ctx context.Context, user *models.User, jwt string) (res *models.Login, err error) {
	res = new(models.Login)
	userFromDB, err := us.UserRepository.GetByUserName(ctx, user.Username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(userFromDB.Password),
		[]byte(user.Password),
	); err != nil {
		return
	}

	userToken, err := token.Create(&userFromDB, jwt)
	if err != nil {
		return
	}

	res.Authorization = userToken
	return
}
