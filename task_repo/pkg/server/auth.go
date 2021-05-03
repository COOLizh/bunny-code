package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// Register ...
func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Register user",
	}

	user := new(models.User)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		s.errJSON(w, marshallErr, http.StatusUnprocessableEntity)
		return
	}

	newUser, err := s.UserService.Register(ctx, user)
	if err != nil {
		s.logError(userAlreadyExists, err, eventAttributes)
		s.errJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response := newUser.PrepareResponse()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		s.errJSON(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Register a new User with Username: %s", response.Username)
}

// Login ...
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Login user",
	}

	user := new(models.User)
	err := user.PopulateFromRequest(r.Body)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		s.errJSON(w, marshallErr, http.StatusUnprocessableEntity)
		return
	}

	if !user.IsValid() {
		s.logError(loginErr, err, eventAttributes)
		s.errJSON(w, loginErr, http.StatusUnprocessableEntity)
		return
	}

	response, err := s.UserService.Login(ctx, user, s.conf.JwtSalt)
	if err != nil {
		s.logError(loginErr, err, eventAttributes)
		s.errJSON(w, loginErr, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		s.logError(marshallErr, err, eventAttributes)
		s.errJSON(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Login User with Username: %s", user.Username)
}
