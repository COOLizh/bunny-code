/*Package cmd provides cmd realization of executor*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/greenteam1/executioner/pkg/executor"
	"gitlab.com/greenteam1/executioner/pkg/models"
)

//Executor the cmd executor struct which contains config
type Executor struct {
	Config map[string]*models.BuildConfig
}

//ExecutorOption is s the signature of option functions
type ExecutorOption func(*Executor) error

// test statuses
const (
	testPassed       = "OK"
	wrongAnswer      = "WA"
	timeLimitError   = "TL"
	compilationError = "CE"
)

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

// TestUserSolution provides checking the user's solution
func (e *Executor) TestUserSolution(sol *models.AggregatedSolution) (*models.TestResults, error) {
	// step 1: check user solution for compilation error
	binaryFilePath, err := e.checkCompilationError(sol.Solution)
	if err != nil {
		if binaryFilePath == compilationError {
			var testResult *models.TestResult = new(models.TestResult)
			testResult.Result = compilationError
			return &models.TestResults{
				TestResults: []*models.TestResult{testResult},
			}, nil
		}
		return nil, err
	}

	// step 2: testing user solution
	resp, err := e.runTestCases(binaryFilePath, sol.Solution)
	if err != nil {
		resp = nil
	}

	// step 3: delete unnecessary binary file
	removeErr := os.Remove(binaryFilePath)
	if removeErr != nil {
		log.Println(removeErr)
	}

	return resp, err
}
