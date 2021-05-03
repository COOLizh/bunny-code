package server

import "net/http"

const (
	reqTimeout          int    = 5
	loginErr            string = "Invalid user data"
	dbErr               string = "DB error"
	marshallErr         string = "Marshall data error"
	userAlreadyExists   string = "User with this username already exists"
	noTaskFound         string = "Task with this id doesn't exist"
	convErr             string = "Convertation error"
	sendSolutionErr     string = "Send a solution error"
	getSolutionResErr   string = "Get solution result error"
	getHistoryErr       string = "Get task history error"
	invalidByteSequence string = "Invalid byte sequence for encoding 'UTF 8'"
)

var resultStatusMap = map[string]int{
	"OK": http.StatusOK,
	"WA": http.StatusOK,
	"TL": http.StatusRequestTimeout,
	"CE": http.StatusBadRequest,
}
