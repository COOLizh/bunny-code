// Package mocks ...
package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// PGUserRepoMock ...
type PGUserRepoMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx and user
func (_m *PGUserRepoMock) Create(ctx context.Context, user *models.User) (r0 models.User, r1 error) {
	ret := _m.Called(ctx, user)

	if rf, ok := ret.Get(0).(func(context.Context, *models.User) models.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserName provides a mock function with given fields: ctx and username
func (_m *PGUserRepoMock) GetByUserName(ctx context.Context, username string) (r0 models.User, r1 error) {
	ret := _m.Called(ctx, username)

	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetMockUser create simple mock Task
func SetMockUser() (user models.User) {
	user = models.User{
		ID:       1,
		Username: "Tester6",
		Password: "$2a$10$BVAZ4K33zfW7.Py6P/cIT.nybt9VFPLcsnEMUGRAXG.lXUZCjnylC",
	}

	return user
}
