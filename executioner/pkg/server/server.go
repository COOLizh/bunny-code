/*Package server comment*/
package server

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"gitlab.com/greenteam1/executioner/pkg/pb"
	"gitlab.com/greenteam1/executioner/pkg/queue"
	"gitlab.com/greenteam1/executioner/pkg/repository"
	"gitlab.com/greenteam1/executioner/pkg/repository/inmemory"

	"gitlab.com/greenteam1/executioner/pkg/executor"
	"gitlab.com/greenteam1/executioner/pkg/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	maxTasksInChannel = 100
)

// Server struct
type Server struct {
	GRPCServer      *grpc.Server
	Executor        executor.Executor
	SolutionResults repository.SolutionsRepository
	Queues          map[string]queue.Queue
}

// New func returns new server instance
func New(exec executor.Executor) *Server {
	return &Server{
		GRPCServer:      grpc.NewServer(),
		Executor:        exec,
		SolutionResults: inmemory.NewSolutionsRepository(),
		Queues:          make(map[string]queue.Queue),
	}
}

// Run register nested services and makes server ready to accept connections
func (s *Server) Run(host string, port string) error {
	ctx := context.Background()
	ctxt, cancel := context.WithCancel(ctx)
	defer cancel()

	ln, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	pb.RegisterCodeHandlerServer(s.GRPCServer, s)
	pb.RegisterStatusHandlerServer(s.GRPCServer, s)

	langs, err := s.Executor.GetLanguages()
	if err != nil {
		return err
	}
	s.StartWorkers(ctxt, langs)

	if err := s.GRPCServer.Serve(ln); err != nil {
		return err
	}

	return nil
}

// Shutdown stops server gracefully. It stops the server from accepting new connections and RPCs and blocks until all the pending RPCs are finished.
func (s *Server) Shutdown() {
	s.GRPCServer.GracefulStop()
}

// CodeHandle handles user's solution
func (s *Server) CodeHandle(ctx context.Context, req *pb.CodeHandleRequest) (*pb.CodeHandleResponse, error) {
	solution := convertPBReqToSolution(req)
	solutionID, err := generateSolutionID()
	if err != nil {
		return &pb.CodeHandleResponse{}, status.Error(codes.Internal, "error while creating ID")
	}

	aggregatedSolution := &models.AggregatedSolution{
		ID:       solutionID,
		Solution: solution,
	}

	err = s.Queues[solution.Language].Push(aggregatedSolution)
	if err != nil {
		return nil, err
	}

	err = s.SolutionResults.Add(ctx, solutionID, &pb.StatusHandleResponse{})
	if err != nil {
		return nil, err
	}

	return &pb.CodeHandleResponse{
		ID:         solutionID,
		JobCreated: true,
	}, nil
}

// StatusCheck returns status of task
func (s *Server) StatusCheck(ctx context.Context, req *pb.StatusHandleRequest) (*pb.StatusHandleResponse, error) {
	resp, err := s.SolutionResults.Pop(ctx, req.ID)
	if err != nil {
		return &pb.StatusHandleResponse{}, err
	}
	return resp, nil
}

// StartWorkers starts new workers from input slice - each in separate goroutine
func (s *Server) StartWorkers(ctx context.Context, keys []string) {
	for _, key := range keys {
		s.Queues[key] = queue.NewChannelQueue(maxTasksInChannel)
		log.Printf("Started worker %s with bufferized channels on %d elements...", key, maxTasksInChannel)
		go func(queue queue.Queue) {
			for {
				incomingSolution, err := queue.Pop()
				if err != nil {
					log.Println(err)
					break
				}
				result, err := s.Executor.TestUserSolution(&incomingSolution)
				if err != nil {
					log.Println(err)
					break
				}
				err = s.SolutionResults.Add(ctx, incomingSolution.ID, result.PrepareResponse())
				if err != nil {
					log.Println(err)
				}
			}
		}(s.Queues[key])
	}
}

func convertPBReqToSolution(req *pb.CodeHandleRequest) *models.Solution {
	var testCases []*models.TestCase
	for _, reqTestCase := range req.TestCases {
		testCase := &models.TestCase{
			TestData: reqTestCase.TestData,
			Answer:   reqTestCase.Answer,
		}
		testCases = append(testCases, testCase)
	}

	sol := &models.Solution{
		Solution:    req.Solution,
		MemoryLimit: req.MemoryLimit,
		TimeLimit:   req.TimeLimit,
		Language:    req.Language,
		TestCases:   testCases,
	}
	return sol
}

func generateSolutionID() (string, error) {
	fileNameID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return fileNameID.String(), err
}
