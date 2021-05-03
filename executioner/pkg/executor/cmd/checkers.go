package cmd

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"gitlab.com/greenteam1/executioner/pkg/models"
)

// checkWrongAnswerError checks solution for a wrong answer
func (e *Executor) checkWrongAnswerError(testResult *models.TestResult, cmd *exec.Cmd, answer []byte) (*models.TestResult, error) {
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(out, answer) {
		testResult.Result = wrongAnswer
		return testResult, nil
	}
	return testResult, nil
}

// checkTimeLimitError checks solution for a time limit exceeded
func (e *Executor) checkTimeLimitError(testResult *models.TestResult, cmd *exec.Cmd, timeLimit int64) *models.TestResult {
	testResult.TimeSpent = cmd.ProcessState.SystemTime().Milliseconds()

	if testResult.TimeSpent > timeLimit {
		testResult.Result = timeLimitError
		return testResult
	}
	return testResult
}

// checkCompilationError returns binary file path if user solution was build, else returns error
func (e *Executor) checkCompilationError(sol *models.Solution) (string, error) {
	//step 1: creating a folder, if one does not exist, where user solutions will be located
	if _, err := os.Stat(os.Getenv("SOLUTION_FOLDER_PATH")); os.IsNotExist(err) {
		err = os.Mkdir(os.Getenv("SOLUTION_FOLDER_PATH"), os.ModePerm)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}

	//generating unique id
	fileNameID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	fileName := fileNameID.String()

	//step 2: creating a file with the extension of the desired programming language and fill it with the code
	filePath, binaryFilePath, err := e.saveCodeFile(fileName, sol.Language, string(sol.Solution))
	if err != nil {
		return "", err
	}

	//step 3: build user solution
	err = e.buildUserSolution(filePath, binaryFilePath, sol.Language)
	if err != nil {
		return compilationError, err
	}

	return binaryFilePath, nil
}
