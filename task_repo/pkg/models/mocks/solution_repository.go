package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// PGSolutionRepoMock ...
type PGSolutionRepoMock struct {
	mock.Mock
}

// StoreSolution ...
func (m *PGSolutionRepoMock) StoreSolution(ctx context.Context, uuid, lang string, userID, taskID int, sol []byte) (r0 time.Time, r1 error) {
	ret := m.Called(ctx, uuid, lang, userID, taskID, sol)

	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int, []byte) time.Time); ok {
		r0 = rf(ctx, uuid, lang, userID, taskID, sol)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int, []byte) error); ok {
		r1 = rf(ctx, uuid, lang, userID, taskID, sol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreResult ...
func (m *PGSolutionRepoMock) StoreResult(ctx context.Context, uuid string, result models.SolutionResult) (r0 error) {
	ret := m.Called(ctx, uuid, result)

	if rf, ok := ret.Get(0).(func(context.Context, string, models.SolutionResult) error); ok {
		r0 = rf(ctx, uuid, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetResultByID ...
func (m *PGSolutionRepoMock) GetResultByID(ctx context.Context, uuid string) (models.SolutionResult, error) {
	var r0 models.SolutionResult
	var r1 error

	ret := m.Called(ctx, uuid)

	if rf, ok := ret.Get(0).(func(context.Context, string) models.SolutionResult); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(models.SolutionResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResultAndTestCaseCountByID ...
func (m *PGSolutionRepoMock) GetResultAndTestCaseCountByID(ctx context.Context, uuid string) (r0 models.SolutionResult, r1 int64, r2 error) {
	ret := m.Called(ctx, uuid)

	if rf, ok := ret.Get(0).(func(context.Context, string) models.SolutionResult); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(models.SolutionResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) int64); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, uuid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetMockResult ...
func SetMockResult() models.SolutionResult {
	return models.SolutionResult{
		ID:               "259ce065-d143-47e4-8487-9ce81f14002f",
		PassedTestsCount: 2,
		TestsCount:       2,
		Results: []*models.TestResult{
			{Status: "OK", Time: 456},
			{Status: "OK", Time: 876},
		},
	}
}
