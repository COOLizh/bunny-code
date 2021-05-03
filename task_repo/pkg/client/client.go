/*Package client sends a request to the executor and receives a response to the user's decision*/
package client

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"gitlab.com/greenteam1/task_repo/pkg/models"
	"gitlab.com/greenteam1/task_repo/pkg/pb"
)

// SendSolution sends code to executor and gets the result of testing
func SendSolution(ctx context.Context, execAddr, language string, userSolution []byte,
	timeLimit, memoryLimit int64, testCases []models.TestCase) (models.SolutionSendResponse, error) {
	reqTestCases := make([]*pb.CodeHandleRequest_TestCase, len(testCases))
	for i := 0; i < len(testCases); i++ {
		reqTestCases[i] = &pb.CodeHandleRequest_TestCase{
			TestData: []byte(testCases[i].TestData),
			Answer:   []byte(testCases[i].Answer),
		}
	}
	req := &pb.CodeHandleRequest{
		Solution:    userSolution,
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   reqTestCases,
	}
	resp, err := sendSolutionToExecutor(ctx, req, execAddr)
	if err != nil {
		return models.SolutionSendResponse{}, err
	}

	return models.SolutionSendResponse{ID: resp.ID}, err
}

// GetSolutionResult requests result from executor by solution id
func GetSolutionResult(ctx context.Context, execAddr, uuid string) (models.SolutionResult, bool, error) {
	conn, err := getClientConnection(execAddr)
	if err != nil {
		return models.SolutionResult{}, false, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	resp, err := getStatusCheckResponse(ctx, conn, uuid)
	if err != nil {
		return models.SolutionResult{}, false, err
	}
	if !resp.Ready {
		return models.SolutionResult{}, false, err
	}

	return prepareTestResults(uuid, resp), true, err
}

// prepareTestResults converts data from response to models.SolutionResult type
func prepareTestResults(uuid string, resp *pb.StatusHandleResponse) models.SolutionResult {
	testsResult := make([]*models.TestResult, 0, len(resp.TestsData.TestResults))
	for _, v := range resp.TestsData.TestResults {
		testsResult = append(testsResult, &models.TestResult{
			Status: v.Result,
			Time:   v.TimeSpent,
		})
	}

	return models.SolutionResult{
		ID:               uuid,
		PassedTestsCount: resp.TestsData.PassedTestsCount,
		Results:          testsResult,
	}
}

// getClientConnection returns client connection
func getClientConnection(target string) (*grpc.ClientConn, error) {
	return grpc.Dial(target, grpc.WithInsecure())
}

/* getCodeHandleResponse sends a request to the executor to test the task
The response is the task ID and whether the job was created for the task */
func getCodeHandleResponse(ctx context.Context, conn *grpc.ClientConn,
	req *pb.CodeHandleRequest) (*pb.CodeHandleResponse, error) {
	codeHandleClient := pb.NewCodeHandlerClient(conn)
	return codeHandleClient.CodeHandle(ctx, req)
}

/* getStatusCheckResponse sends a request to the executor to check the complete testing of the task. If it is tested, it returns the test results */
func getStatusCheckResponse(ctx context.Context, conn *grpc.ClientConn,
	uuid string) (*pb.StatusHandleResponse, error) {
	statusCheckClient := pb.NewStatusHandlerClient(conn)
	statusRequest := &pb.StatusHandleRequest{
		ID: uuid,
	}
	return statusCheckClient.StatusCheck(ctx, statusRequest)
}

// sendSolutionToExecutor returns the result of the user's solution
func sendSolutionToExecutor(ctx context.Context, req *pb.CodeHandleRequest,
	execAddr string) (*pb.CodeHandleResponse, error) {
	conn, err := getClientConnection(execAddr)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	return getCodeHandleResponse(ctx, conn, req)
}
