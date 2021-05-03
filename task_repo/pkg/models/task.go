package models

// Task struct
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TimeLimit   int64  `json:"time_limit"`
	Memory      int64  `json:"memory"`
}
