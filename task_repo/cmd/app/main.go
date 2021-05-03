package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"gitlab.com/greenteam1/task_repo/pkg/config"
	"gitlab.com/greenteam1/task_repo/pkg/db"
	"gitlab.com/greenteam1/task_repo/pkg/repository/postgres"
	"gitlab.com/greenteam1/task_repo/pkg/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := new(config.Config)
	err := conf.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	pool, err := db.Connect(ctx, conf.DatabaseURL)
	if err != nil {
		log.Println(err)
	}
	defer pool.Close()

	taskRepository := postgres.NewTasksRepository(pool)
	userRepository := postgres.NewUserRepository(pool)
	solutionRepository := postgres.NewSolutionRepository(pool)
	srv := server.New(conf, taskRepository, userRepository, solutionRepository)

	go func() {
		if err := srv.Run(ctx); err != nil && err != http.ErrServerClosed {
			srv.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.WithFields(log.Fields{
			"err": err,
		}).Fatal("Server forced to shutdown")
	}
}
