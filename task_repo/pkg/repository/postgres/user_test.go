package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/models/mocks"
)

func TestNewUserRepository(t *testing.T) {
	t.Run("NewUserRepository", func(t *testing.T) {
		r := NewUserRepository(&pgxpool.Pool{})
		assert.IsType(t, &UserRepositoryPostgresql{}, r)
	})
}

func TestCreate(t *testing.T) {
	mockPGUserRepo := mocks.PGUserRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockUser := mocks.SetMockUser()
		mockPGUserRepo.On(
			"Create",
			mock.Anything,
			mock.Anything,
		).Return(mockUser, nil).Once()

		user, err := mockPGUserRepo.Create(
			context.TODO(),
			&models.User{Username: "name", Password: "pass"},
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

		user, err := mockPGUserRepo.Create(
			context.TODO(),
			&models.User{Username: "name", Password: "pass"},
		)

		assert.Error(t, err)
		assert.Exactly(t, models.User{}, user)
	})
}

func TestGetByUserName(t *testing.T) {
	mockPGUserRepo := mocks.PGUserRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockUser := mocks.SetMockUser()
		mockPGUserRepo.On(
			"GetByUserName",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockUser, nil).Once()

		user, err := mockPGUserRepo.GetByUserName(
			context.TODO(),
			"name",
		)

		assert.NoError(t, err)
		assert.Exactly(t, mockUser, user)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPGUserRepo.On(
			"GetByUserName",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(models.User{}, errors.New("Unexpexted Error")).Once()

		user, err := mockPGUserRepo.GetByUserName(
			context.TODO(),
			"name",
		)

		assert.Error(t, err)
		assert.Exactly(t, models.User{}, user)
	})
}
