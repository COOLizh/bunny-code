package models

import (
	"time"
)

// SolutionHistoryItem ...
type SolutionHistoryItem struct {
	ID        string              `json:"id"`
	Solution  SolutionSendRequest `json:"solution"`
	Result    SolutionResult      `json:"result"`
	CreatedAt time.Time           `json:"created_at"`
}

// SolutionSendRequest contains fields that server gets from POST-Solution request
type SolutionSendRequest struct {
	Code     []byte `json:"solution"`
	Language string `json:"language"`
}

// SolutionSendResponse contains fields with which the server responds to the POST-Solution request if it's ok
type SolutionSendResponse struct {
	ID string `json:"id"`
}

// SolutionResult contains needful information about user solution result
type SolutionResult struct {
	ID               string        `json:"id"`
	PassedTestsCount int64         `json:"passed_tests_count"`
	TestsCount       int64         `json:"tests_count"`
	Results          []*TestResult `json:"results"`
}

// TestResult contains information about test result
type TestResult struct {
	Status string `json:"status"`
	Time   int64  `json:"time"`
}

// SolutionStatusResponseMessage contains status message
type SolutionStatusResponseMessage struct {
	Message string `json:"message"`
}
