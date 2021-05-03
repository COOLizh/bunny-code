/*Package docker provides dockerized realization of executor*/
package docker

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"gitlab.com/greenteam1/executioner/pkg/dockerapi"
	"gitlab.com/greenteam1/executioner/pkg/executor"
	"gitlab.com/greenteam1/executioner/pkg/models"
)

// test statuses
const (
	testPassed        = "OK"
	wrongAnswer       = "WA"
	compilationError  = "CE"
	timeLimitExceeded = "TL"
)

// Executor struct includes build config, docker client, containers IDs, and test user solution methods
type Executor struct {
	Config        map[string]*models.BuildConfig
	DockerCli     *client.Client
	ContainersIDs map[string]string
}

// ExecutorOption is s the signature of option functions
type ExecutorOption func(*Executor) error

// NewExecutor creates new instance of Executor
func NewExecutor(opts ...ExecutorOption) (*Executor, error) {
	exec := new(Executor)
	var err error
	for _, opt := range opts {
		err = opt(exec)
		if err != nil {
			return nil, err
		}
	}
	return exec, nil
}

//Config is an option function for setting up config
func Config(buildConfigPath string) ExecutorOption {
	return func(e *Executor) error {
		config := executor.NewBuildConfig()
		err := config.Read(buildConfigPath)
		if err != nil {
			return err
		}
		e.Config = config.Config
		return nil
	}
}

//Cli is an option function for serring up DockerCli
func Cli() ExecutorOption {
	return func(e *Executor) error {
		dockerCli, err := dockerapi.NewDockerClient()
		if err != nil {
			return err
		}
		e.DockerCli = dockerCli
		return nil
	}
}

//ContainersIDs is an optionfunction forstting up containers ids
func ContainersIDs() ExecutorOption {
	return func(e *Executor) error {
		ContainersIDs, err := dockerapi.StartContainers(e.DockerCli, e.Config)
		if err != nil {
			return err
		}
		e.ContainersIDs = ContainersIDs
		return nil
	}
}

// TestUserSolution provides checking the user's solution
func (e *Executor) TestUserSolution(sol *models.AggregatedSolution) (*models.TestResults, error) {
	ctx := context.Background()

	solPath, err := e.storeUserSolution(sol.ID, sol.Solution)
	if err != nil {
		return &models.TestResults{}, err
	}
	defer func() {
		removeErr := os.RemoveAll(solPath)
		if removeErr != nil {
			log.Println(removeErr)
		}
	}()

	solArchiveReader, err := e.solArchive(solPath, sol.ID, sol.Solution.Language)
	if err != nil {
		return nil, err
	}
	defer func() {
		removeErr := os.Remove(solArchiveReader.Name())
		if removeErr != nil {
			log.Println(removeErr)
		}
	}()

	containerID := e.ContainersIDs[sol.Solution.Language]

	err = e.DockerCli.CopyToContainer(
		ctx,
		containerID,
		e.Config[sol.Solution.Language].SrcPath,
		solArchiveReader,
		types.CopyToContainerOptions{AllowOverwriteDirWithFile: true},
	)

	if err != nil {
		return nil, err
	}

	resp, err := e.runTestCasesInContainer(
		ctx,
		e.DockerCli,
		containerID,
		sol.Solution,
		sol.ID,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetLanguages returns a list of supported languages, defined in config. It will throw an error, if no one language would be declared
func (e *Executor) GetLanguages() ([]string, error) {
	lAmount := len(e.Config)
	if lAmount == 0 {
		return nil, fmt.Errorf("no languages specified in config")
	}

	langs := make([]string, 0, lAmount)
	for k := range e.Config {
		langs = append(langs, k)
	}

	return langs, nil
}
