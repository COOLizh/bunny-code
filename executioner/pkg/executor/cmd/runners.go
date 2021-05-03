package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gitlab.com/greenteam1/executioner/pkg/models"
)

// saveCodeFile create user solution file and return file path and biniray file path of user solution
func (e *Executor) saveCodeFile(fileName, language, code string) (string, string, error) {
	if _, ok := e.Config[language]; !ok {
		return "", "", fmt.Errorf("ERROR: %s is not supported language by executor", language)
	}
	fileExtension := e.Config[language].FileExtension
	binaryFilePath := filepath.Join(os.Getenv("SOLUTION_FOLDER_PATH"), fileName)
	filePath := binaryFilePath + fileExtension
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", "", err
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}()
	if _, err := f.WriteString(code); err != nil {
		return "", "", err
	}
	return filePath, binaryFilePath, nil
}

// buildUserSolution returns error if user solution build failed
func (e *Executor) buildUserSolution(filePath, binaryFilePath, language string) error {
	buildCommands := strings.Split(e.Config[language].BuildCommands, " ")
	// #nosec G204
	buildCommands = append(buildCommands, binaryFilePath)
	buildCommands = append(buildCommands, filePath)
	cmd := exec.Command(buildCommands[0], buildCommands[1:]...) //nolint:gosec

	// build
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	// remove unnecessary file
	removeErr := os.Remove(filePath)
	if removeErr != nil {
		log.Println(removeErr)
	}

	return err
}

// inputTestCase returns error if test case can not be enter
func (e *Executor) inputTestCase(cmd *exec.Cmd, testData []byte) error {
	var err error
	stdin, err := cmd.StdinPipe()
	go func() {
		var stdinError error
		defer func() {
			stdinError = stdin.Close()
			if stdinError != nil {
				log.Println(stdinError)
			}
		}()
		_, stdinError = io.WriteString(stdin, string(testData))
		if stdinError != nil {
			log.Println(stdinError)
		}
	}()
	return err
}

// runTestCase runs test cases and returns result of testing user solution
func (e *Executor) runTestCases(binaryFilePath string, sol *models.Solution) (*models.TestResults, error) {
	resp := &models.TestResults{
		TestResults: make([]*models.TestResult, len(sol.TestCases)),
	}

	for i := 0; i < len(sol.TestCases); i++ {
		var testResult *models.TestResult = new(models.TestResult)

		// step 1: run user solution
		cmd := e.runUserSolution(binaryFilePath)
		var err error

		// step 2: input test values
		err = e.inputTestCase(cmd, sol.TestCases[i].TestData)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// step 3: check Wrong Answer error
		testResult, err = e.checkWrongAnswerError(testResult, cmd, sol.TestCases[i].Answer)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if testResult.Result == wrongAnswer {
			resp.TestResults[i] = testResult
			return resp, err
		}

		// step 4: check Time Limit error
		testResult = e.checkTimeLimitError(testResult, cmd, sol.TimeLimit)
		if testResult.Result == timeLimitError {
			resp.TestResults[i] = testResult
			return resp, nil
		}

		// step 5: check Memory Limit error
		testResult.MemorySpent = 1

		// if test passed
		testResult.Result = testPassed
		resp.TestResults[i] = testResult
		resp.PassedTestsCount++
	}
	return resp, nil
}

// runUserSolution runs user solution
func (e *Executor) runUserSolution(binaryFilePath string) *exec.Cmd {
	// #nosec G204
	return exec.Command(binaryFilePath)
}
