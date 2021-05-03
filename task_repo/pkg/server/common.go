package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

func (s *Server) getTaskID(r *http.Request) (int, error) {
	uuid, ok := mux.Vars(r)["task_id"]
	if !ok {
		return 0, errors.New("can't get task id from request")
	}
	return strconv.Atoi(uuid)
}

func (s *Server) getSolID(r *http.Request) (string, error) {
	uuid, ok := mux.Vars(r)["sol_id"]
	if !ok {
		return "", errors.New("can't get solution id from request")
	}
	return uuid, nil
}

func (s *Server) getUserID(r *http.Request) (int, error) {
	userIDInterface := r.Context().Value(models.ContextTokenID)
	if userIDInterface == nil {
		return 0, errors.New("can't get user id from request")
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		return 0, errors.New("can't process user id from request")
	}
	return userID, nil
}

func (s *Server) logError(
	nameError string,
	err error,
	eventAttributes logrus.Fields,
) {
	eventAttributes["err"] = err
	s.Logger.WithFields(eventAttributes).Error(nameError)
}

func (s *Server) errJSON(w http.ResponseWriter, errText string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	body := struct {
		Error string `json:"error"`
	}{Error: errText}

	if err := json.NewEncoder(w).Encode(&body); err != nil {
		eventAttributes := logrus.Fields{
			"action": "marshal response",
		}
		s.logError(marshallErr, err, eventAttributes)
	}
}
