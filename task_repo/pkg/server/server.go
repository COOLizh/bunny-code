// Package server ...
package server

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"gitlab.com/greenteam1/task_repo/pkg/config"
	"gitlab.com/greenteam1/task_repo/pkg/handlers"
	"gitlab.com/greenteam1/task_repo/pkg/middleware"
	"gitlab.com/greenteam1/task_repo/pkg/repository"
	"gitlab.com/greenteam1/task_repo/pkg/service"
)

// Server ...
type Server struct {
	conf            *config.Config
	httpSrv         *http.Server
	router          *mux.Router
	Logger          *logrus.Logger
	UserService     *service.UserService
	TasksService    *service.TaskService
	SolutionService *service.SolutionService
}

// New ...
func New(cfg *config.Config, taskRepository repository.TaskRepositoryInterface, userRepository repository.UserRepositoryInterface, solutionRepository repository.SolutionRepositoryInterface) *Server {
	taskSVC := service.NewTaskService(taskRepository)
	userSVC := service.NewUserService(userRepository)
	solutionSVC := service.NewSolutionService(solutionRepository)

	return &Server{
		conf:            cfg,
		Logger:          logrus.New(),
		UserService:     userSVC,
		TasksService:    taskSVC,
		SolutionService: solutionSVC,
	}
}

// SetupRouter ...
func (s *Server) SetupRouter() {
	router := mux.NewRouter()
	router.Use(middleware.SetCORSHeaders(s.conf))

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	subrouter.Methods(http.MethodOptions)
	subrouter.HandleFunc("/", handlers.HealthCheck).Methods(http.MethodGet)
	subrouter.HandleFunc("/registration", s.Register).Methods(http.MethodPost)
	subrouter.HandleFunc("/login", s.Login).Methods(http.MethodPost)
	subrouter.HandleFunc("/tasks", s.GetAllTasks).Methods(http.MethodGet)
	subrouter.HandleFunc("/task/{task_id}", s.GetTaskByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/testcase/{task_id}", s.GetTestCasesByTaskID).Methods(http.MethodGet)

	authRouter := subrouter.PathPrefix("/solution").Subrouter()
	authRouter.HandleFunc("/task/{task_id}", s.SendSolution).Methods(http.MethodPost)
	authRouter.HandleFunc("/{sol_id}", s.GetSolutionResult).Methods(http.MethodGet)
	authRouter.HandleFunc("/task/{task_id}/history", s.GetTaskHistory).Methods(http.MethodGet)
	authRouter.HandleFunc("/auth", handlers.AuthCheck).Methods(http.MethodGet)
	authRouter.Use(middleware.AuthValidateTokenFunc(s.conf))

	s.router = router
}

// Run ...
func (s *Server) Run(_ context.Context) error {
	s.SetupRouter()

	file, err := os.OpenFile(s.conf.FileLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		s.Logger.SetOutput(mw)
	} else {
		s.Logger.Error("Failed to log to file, using default stderr")
	}

	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	Formatter.ForceColors = true
	logrus.SetFormatter(Formatter)

	s.Logger.WithFields(logrus.Fields{
		"port": s.conf.APIPort,
	}).Info("Service starting")

	s.httpSrv = &http.Server{
		Addr:    s.conf.APIPort,
		Handler: s.router,
	}

	return s.httpSrv.ListenAndServe()
}

// Shutdown ...
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}
