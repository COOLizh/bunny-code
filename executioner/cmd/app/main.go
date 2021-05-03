/*Package main comment*/
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/greenteam1/executioner/pkg/executor/docker"
	"gitlab.com/greenteam1/executioner/pkg/server"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found, looking for OS environment")
	}
}

func main() {
	port := os.Getenv("EXECUTIONER_PORT")
	host := os.Getenv("EXECUTIONER_HOST")
	buildConfigPath := os.Getenv("BUILD_CONFIG_PATH")

	log.Printf("Starting server on port %s", port)

	exec, err := docker.NewExecutor(
		docker.Config(buildConfigPath),
		docker.Cli(),
		docker.ContainersIDs(),
	)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(exec)

	go func() {
		if err := srv.Run(host, port); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	srv.Shutdown()
}
