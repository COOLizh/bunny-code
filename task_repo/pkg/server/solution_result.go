package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

// SendSolution ...
func (s *Server) SendSolution(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Send solution to executor",
	}

	taskID, err := s.getTaskID(r)
	if err != nil {
		s.logError(convErr, err, eventAttributes)
		http.Error(w, convErr, http.StatusInternalServerError)
		return
	}

	solutionRequest := new(models.SolutionSendRequest)

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(solutionRequest)

	if err != nil {
		s.logError(invalidByteSequence, err, eventAttributes)
		http.Error(w, invalidByteSequence, http.StatusBadRequest)
		return
	}

	sendSolResp, err := s.TasksService.SendSolution(ctx, solutionRequest, s.conf.ExecutionerPort, taskID)
	if err != nil {
		s.logError(sendSolutionErr, err, eventAttributes)
		http.Error(w, sendSolutionErr, http.StatusBadRequest)
		return
	}

	userID, err := s.getUserID(r)
	if err != nil {
		s.logError(sendSolutionErr, err, eventAttributes)
		http.Error(w, sendSolutionErr, http.StatusBadRequest)
		return
	}

	if _, err = s.SolutionService.StoreSolution(ctx, sendSolResp.ID, userID, taskID,
		solutionRequest.Language, solutionRequest.Code); err != nil {
		s.logError(sendSolutionErr, err, eventAttributes)
		http.Error(w, sendSolutionErr, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&sendSolResp)
	if err != nil {
		eventAttributes["err"] = err
		s.Logger.WithFields(eventAttributes).Error(marshallErr)
		http.Error(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Send a solution to executor by task with ID: %d", taskID)
}

// GetSolutionResult responds with solution result, if it is ready, or with message 'pending', if not
func (s *Server) GetSolutionResult(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Request solution result from executor",
	}
	w.Header().Set("Content-Type", "application/json")

	uuid, err := s.getSolID(r)
	if err != nil {
		s.logError(convErr, err, eventAttributes)
		http.Error(w, convErr, http.StatusInternalServerError)
		return
	}

	res, ready, err := s.SolutionService.GetSolutionResult(ctx, s.conf.ExecutionerPort, uuid)
	if err != nil {
		s.logError(getSolutionResErr, err, eventAttributes)
		st, ok := status.FromError(err)
		if !ok {
			http.Error(w, getSolutionResErr, http.StatusInternalServerError)
			return
		}
		respondWithErrFromStatus(st, w)
		return
	}
	if !ready {
		w.WriteHeader(http.StatusOK)
		msg := models.SolutionStatusResponseMessage{Message: "pending"}
		err = json.NewEncoder(w).Encode(&msg)
		if err != nil {
			eventAttributes["err"] = err
			s.Logger.WithFields(eventAttributes).Error(marshallErr)
			http.Error(w, marshallErr, http.StatusBadRequest)
		}
		s.Logger.Infof("Respond with \"pending\" by solution with ID: %s", uuid)
		return
	}

	w.WriteHeader(getSendSolutionHTTPStatus(&res))
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		eventAttributes["err"] = err
		s.Logger.WithFields(eventAttributes).Error(marshallErr)
		http.Error(w, marshallErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Respond with solution by ID: %s", uuid)
}

// GetTaskHistory ...
func (s *Server) GetTaskHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(reqTimeout))
	defer cancel()

	eventAttributes := logrus.Fields{
		"action": "Get task history",
	}

	userID, err := s.getUserID(r)
	if err != nil {
		s.logError(getHistoryErr, err, eventAttributes)
		s.errJSON(w, getHistoryErr, http.StatusBadRequest)
		return
	}

	taskID, err := s.getTaskID(r)
	if err != nil {
		s.logError(convErr, err, eventAttributes)
		s.errJSON(w, getHistoryErr, http.StatusInternalServerError)
		return
	}

	history, err := s.SolutionService.GetSolutionsHistory(ctx, userID, taskID)
	if err != nil || len(history) < 1 {
		s.logError(sendSolutionErr, err, eventAttributes)
		s.errJSON(w, getHistoryErr, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&history)
	if err != nil {
		eventAttributes["err"] = err
		s.Logger.WithFields(eventAttributes).Error(marshallErr)
		s.errJSON(w, getHistoryErr, http.StatusBadRequest)
		return
	}

	s.Logger.Infof("Respond with history by task with ID: %d", taskID)
}

func respondWithErrFromStatus(st *status.Status, w http.ResponseWriter) {
	switch st.Code() {
	case codes.Unavailable:
		http.Error(w, getSolutionResErr, http.StatusInternalServerError)
		return
	case codes.NotFound:
		http.Error(w, st.Message(), http.StatusNotFound)
		return
	default:
		http.Error(w, st.Code().String(), http.StatusInternalServerError)
		return
	}
}

func getSendSolutionHTTPStatus(solutionResult *models.SolutionResult) (httpStatus int) {
	if solutionResult.PassedTestsCount == solutionResult.TestsCount {
		httpStatus = resultStatusMap[solutionResult.Results[0].Status]
	} else {
		httpStatus = resultStatusMap[solutionResult.Results[solutionResult.PassedTestsCount].Status]
	}

	return
}
