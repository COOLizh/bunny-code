package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// PGTaskRepoMock ...
type PGTaskRepoMock struct {
	mock.Mock
}

// GetTaskByID provides a mock function with given fields: id
func (_m *PGTaskRepoMock) GetTaskByID(_ context.Context, id int) (r0 models.Task, r1 error) {
	ret := _m.Called(id)

	if rf, ok := ret.Get(0).(func(int) models.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTasks ...
func (_m *PGTaskRepoMock) GetAllTasks(_ context.Context) (r0 []models.Task, r1 error) {
	ret := _m.Called()

	if rf, ok := ret.Get(0).(func() []models.Task); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]models.Task)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetTestCasesByTaskID provides a mock function with given fields: id
func (_m *PGTaskRepoMock) GetTestCasesByTaskID(_ context.Context, id int) (r0 []models.TestCase, r1 error) {
	ret := _m.Called(id)

	if rf, ok := ret.Get(0).(func(int) []models.TestCase); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).([]models.TestCase)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetMockTask create simple mock Task
func SetMockTask() (task models.Task) {
	task = models.Task{
		ID:          1,
		Name:        "name",
		Description: "desc",
		TimeLimit:   10,
		Memory:      10,
	}

	return
}

// SetMockTestCase create simple mock Task
func SetMockTestCase() (tc models.TestCase) {
	tc = models.TestCase{
		ID:       1,
		TaskID:   1,
		TestData: "data",
		Answer:   "answer",
	}

	return
}
