package models

// TestCase structSS
type TestCase struct {
	ID       int    `json:"id"`
	TaskID   int    `json:"task_id"`
	TestData string `json:"test_data"`
	Answer   string `json:"answer"`
}
