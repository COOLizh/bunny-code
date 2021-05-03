package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// GetAllTasks ...
func (s *Server) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Get testCase by taskID",
	}

	tasks, err := s.TasksService.TaskRepository.GetAllTasks(ctx)
	if err != nil {
		s.logError(dbErr, err, eventAttributes)
		http.Error(w, dbErr, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&tasks)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		http.Error(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Info("Get all tasks")
}

// GetTaskByID by id task
func (s *Server) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Get task",
	}

	taskID, err := s.getTaskID(r)
	if err != nil {
		s.logError(convErr, err, eventAttributes)
		http.Error(w, convErr, http.StatusInternalServerError)
		return
	}

	eventAttributes["taskID"] = taskID

	task, err := s.TasksService.GetTaskByID(ctx, taskID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			http.Error(w, noTaskFound, http.StatusNotFound)
			return
		}
		s.logError(dbErr, err, eventAttributes)
		http.Error(w, dbErr, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&task)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		http.Error(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Get task with ID: %d", taskID)
}

// GetTestCasesByTaskID by id task
func (s *Server) GetTestCasesByTaskID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Get testCase by taskID",
	}

	taskID, err := s.getTaskID(r)
	if err != nil {
		s.logError(convErr, err, eventAttributes)
		http.Error(w, convErr, http.StatusInternalServerError)
		return
	}

	testCases, err := s.TasksService.GetTestCasesByTaskID(ctx, taskID)

	eventAttributes["taskID"] = taskID

	if err != nil {
		s.logError(dbErr, err, eventAttributes)
		http.Error(w, dbErr, http.StatusInternalServerError)
		return
	}

	if testCases == nil {
		http.Error(w, noTaskFound, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&testCases)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		http.Error(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Get testcases by task with ID: %d", taskID)
}
