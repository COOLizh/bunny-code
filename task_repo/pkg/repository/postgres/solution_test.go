package postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"gitlab.com/greenteam1/task_repo/pkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/greenteam1/task_repo/pkg/models/mocks"
)

func TestNewSolutionRepository(t *testing.T) {
	t.Run("NewSolutionRepository", func(t *testing.T) {
		r := NewSolutionRepository(&pgxpool.Pool{})
		assert.IsType(t, &SolutionRepositoryPostgresql{}, r)
	})
}

func TestStoreSolution(t *testing.T) {
	mockPGSolRepo := mocks.PGSolutionRepoMock{}

	t.Run("success", func(t *testing.T) {
		tm := time.Now()
		mockPGSolRepo.On(
			"StoreSolution",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
			mock.Anything,
		).Return(tm, nil).Once()

		resp, err := mockPGSolRepo.StoreSolution(
			context.Background(),
			"1234",
			"lang",
			1,
			1,
			[]byte{},
		)

		assert.NoError(t, err)
		assert.Exactly(t, tm, resp)
	})

	t.Run("error-failed", func(t *testing.T) {
		tm := time.Now()
		mockPGSolRepo.On(
			"StoreSolution",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
			mock.Anything,
		).Return(tm, errors.New("error")).Once()

		resp, err := mockPGSolRepo.StoreSolution(
			context.Background(),
			"1234",
			"lang",
			1,
			1,
			[]byte{},
		)

		assert.Error(t, err)
		assert.Exactly(t, tm, resp)
	})
}

func TestStoreResult(t *testing.T) {
	mockPGSolRepo := mocks.PGSolutionRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockRes := mocks.SetMockResult()
		mockPGSolRepo.On(
			"StoreResult",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.Anything,
		).Return(nil).Once()

		err := mockPGSolRepo.StoreResult(
			context.Background(),
			"1234",
			mockRes,
		)

		assert.NoError(t, err)
	})
}

func TestGetResultByID(t *testing.T) {
	mockPGSolRepo := mocks.PGSolutionRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockRes := mocks.SetMockResult()
		mockPGSolRepo.On(
			"GetResultByID",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockRes, nil).Once()

		res, err := mockPGSolRepo.GetResultByID(
			context.Background(),
			"1234",
		)

		assert.NoError(t, err)
		assert.Exactly(t, mockRes, res)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPGSolRepo.On(
			"GetResultByID",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(models.SolutionResult{}, errors.New("error")).Once()

		resp, err := mockPGSolRepo.GetResultByID(
			context.Background(),
			"1234",
		)

		assert.Error(t, err)
		assert.Exactly(t, models.SolutionResult{}, resp)
	})
}

func TestGetResultAndTestCaseCountByID(t *testing.T) {
	mockPGSolRepo := mocks.PGSolutionRepoMock{}

	t.Run("success", func(t *testing.T) {
		mockRes := mocks.SetMockResult()
		var mockCount int64 = 3
		mockPGSolRepo.On(
			"GetResultAndTestCaseCountByID",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockRes, mockCount, nil).Once()

		res, count, err := mockPGSolRepo.GetResultAndTestCaseCountByID(
			context.Background(),
			"1234",
		)

		assert.NoError(t, err)
		assert.Exactly(t, mockRes, res)
		assert.Exactly(t, mockCount, count)
	})

	t.Run("success", func(t *testing.T) {
		mockRes := mocks.SetMockResult()
		var mockCount int64 = 3
		mockPGSolRepo.On(
			"GetResultAndTestCaseCountByID",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(mockRes, mockCount, nil).Once()

		res, count, err := mockPGSolRepo.GetResultAndTestCaseCountByID(
			context.Background(),
			"1234",
		)

		assert.NoError(t, err)
		assert.Exactly(t, mockRes, res)
		assert.Exactly(t, mockCount, count)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockPGSolRepo.On(
			"GetResultAndTestCaseCountByID",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(models.SolutionResult{}, int64(0), errors.New("error")).Once()

		resp, count, err := mockPGSolRepo.GetResultAndTestCaseCountByID(
			context.Background(),
			"1234",
		)

		assert.Error(t, err)
		assert.Exactly(t, models.SolutionResult{}, resp)
		assert.Exactly(t, int64(0), count)
	})
}
