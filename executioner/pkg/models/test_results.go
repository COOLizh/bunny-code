package models

import (
	"gitlab.com/greenteam1/executioner/pkg/pb"
)

// TestResults describes test results received from executor
type TestResults struct {
	PassedTestsCount int64         `json:"passed_test_count"`
	TestResults      []*TestResult `json:"results"`
}

// AddResult adds new element to resulting array
func (t *TestResults) AddResult(res string, time, mem int64) {
	t.TestResults = append(t.TestResults, &TestResult{
		Result:      res,
		TimeSpent:   time,
		MemorySpent: mem,
	})
}

// PrepareResponse converts result to *pb.StatusHandleResponse
func (t *TestResults) PrepareResponse() *pb.StatusHandleResponse {
	var results []*pb.StatusHandleResponse_TestsData_TestResult
	for _, testResult := range t.TestResults {
		if testResult == nil {
			break
		}
		result := &pb.StatusHandleResponse_TestsData_TestResult{
			Result:      testResult.Result,
			TimeSpent:   testResult.TimeSpent,
			MemorySpent: testResult.MemorySpent,
		}
		results = append(results, result)
	}

	resp := &pb.StatusHandleResponse{
		Ready: true,
		TestsData: &pb.StatusHandleResponse_TestsData{
			PassedTestsCount: t.PassedTestsCount,
			TestResults:      results,
		},
	}
	return resp
}

// TestResult describes result of testing user's solution on a test case
type TestResult struct {
	Result      string `json:"status"`
	TimeSpent   int64  `json:"time"`
	MemorySpent int64  `json:"memory"`
}
