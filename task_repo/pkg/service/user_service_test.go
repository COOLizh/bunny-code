package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/models/mocks"
)

func TestRegister(t *testing.T) {
	mockPGUserRepo := mocks.PGUserRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockUser := mocks.SetMockUser()
		mockPGUserRepo.On(
			"Create",
			mock.Anything,
			mock.Anything,
		).Return(mockUser, nil).Once()

		userServiceMock := NewUserService(&mockPGUserRepo)

		ctx := context.Background()
		user, err := userServiceMock.Register(
			ctx,
			&models.User{Username: "name", Password: "password12"},
		)

		assert.NoError(t, err)
		assert.Exactly(t, mockUser, user)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPGUserRepo.On(
			"Create",
			mock.Anything,
			mock.Anything,
		).Return(models.User{}, errors.New("Unexpexted Error")).Once()

		userServiceMock := NewUserService(&mockPGUserRepo)

		user, err := userServiceMock.Register(
			context.TODO(),
			&models.User{Username: "name", Password: "password12"},
		)

		assert.Error(t, err)
		assert.Exactly(t, models.User{}, user)
	})
}

func TestLogin(t *testing.T) {
	mockPGUserRepo := mocks.PGUserRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockUser := mocks.SetMockUser()
		mockPGUserRepo.On(
			"GetByUserName",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockUser, nil).Once()

		userServiceMock := NewUserService(&mockPGUserRepo)

		ctx := context.Background()
		// ctx = context.WithValue(ctx, models.ContextJWTKey, "verysecret")
		login, err := userServiceMock.Login(
			ctx,
			&models.User{Username: "tester6", Password: "qwerty"},
			"verysecret",
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, login.Authorization)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPGUserRepo.On(
			"GetByUserName",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(models.User{}, nil).Once()

		userServiceMock := NewUserService(&mockPGUserRepo)

		login, err := userServiceMock.Login(
			context.TODO(),
			&models.User{Username: "name", Password: "password12"},
			"verysecret2",
		)

		assert.Error(t, err)
		assert.Empty(t, login.Authorization)
	})
}
