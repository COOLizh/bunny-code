package models

// Solution describes user's solution from request
type Solution struct {
	Solution    []byte      `json:"solution"`
	MemoryLimit int64       `json:"memory_limit"`
	TimeLimit   int64       `json:"time_limit"`
	Language    string      `json:"language"`
	TestCases   []*TestCase `json:"test_cases"`
}

// TestCase describes test that should be performed on user's solution
type TestCase struct {
	TestData []byte `json:"test_data"`
	Answer   []byte `json:"answer"`
}

//AggregatedSolution struct is containing solution id's and solution request
type AggregatedSolution struct {
	ID       string
	Solution *Solution
}
