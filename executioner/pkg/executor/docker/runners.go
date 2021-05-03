package docker

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/client"
	"golang.org/x/tools/txtar"

	"gitlab.com/greenteam1/executioner/pkg/dockerapi"
	"gitlab.com/greenteam1/executioner/pkg/models"
)

type file struct {
	Name string
	Data []byte
}

// runTestCasesInContainer builds solution and executes it with given test data
func (e *Executor) runTestCasesInContainer(ctx context.Context, cli *client.Client,
	containerID string, sol *models.Solution, solutionID string) (*models.TestResults, error) {
	resp := &models.TestResults{
		TestResults: make([]*models.TestResult, 0, len(sol.TestCases)),
	}

	binPath := e.pathForExecSolutionInContainer(sol.Language, solutionID)
	cmd := e.commandForBuild(sol.Language, solutionID, binPath)
	buildResult, err := dockerapi.BuildSolutionInContainer(ctx, cli, containerID, cmd)
	if err != nil {
		return nil, err
	}
	if buildResult != nil {
		err = dockerapi.DeleteSolutionFromContainer(ctx, cli, containerID, solutionID)
		if err != nil {
			log.Println(err)
		}
		resp.AddResult(compilationError, 0, 0)
		return resp, nil
	}

Loop:
	for _, v := range sol.TestCases {
		var answer []byte
		var timeSpent int64
		if answer, timeSpent, err = dockerapi.ExecuteSolutionInContainer(
			ctx, cli, containerID, string(v.TestData), binPath,
		); err != nil {
			return nil, err
		}

		switch {
		case !bytes.Equal(answer, v.Answer):
			resp.AddResult(wrongAnswer, timeSpent, 0)
			break Loop
		case timeSpent > sol.TimeLimit:
			resp.AddResult(timeLimitExceeded, timeSpent, 0)
			break Loop
		default:
			resp.AddResult(testPassed, timeSpent, 0)
			resp.PassedTestsCount++
		}
	}
	err = dockerapi.DeleteSolutionFromContainer(ctx, cli, containerID, solutionID)
	if err != nil {
		log.Println(err)
	}
	err = dockerapi.DeleteBuildFromContainer(ctx, cli, containerID, e.Config[sol.Language].BinPath, solutionID)
	if err != nil {
		log.Println(err)
	}

	return resp, nil
}

// storeUserSolution saves user's solution in file and returns path to it
func (e *Executor) storeUserSolution(id string, sol *models.Solution) (string, error) {
	solutionsDir := os.Getenv("SOLUTION_FOLDER_PATH")
	err := e.createDir(solutionsDir)
	if err != nil {
		return "", err
	}

	solPath := filepath.Join(solutionsDir, id)

	err = e.saveCodeFile(
		id,
		solPath,
		sol.Language,
		sol.Solution,
	)
	if err != nil {
		return "", err
	}

	return solPath, nil
}

// saveCodeFile create user solution file and return file path and biniray file path of user solution
func (e *Executor) saveCodeFile(solName, dirName, language string, code []byte) error {
	if _, ok := e.Config[language]; !ok {
		return fmt.Errorf("ERROR: %s is not supported language by Executor", language)
	}

	err := e.createDir(dirName)
	if err != nil {
		return err
	}

	fileExt := e.Config[language].FileExtension
	fileSet := lookupFiles(code, solName)
	for _, v := range fileSet {
		var filePathToSave string
		switch filepath.Ext(v.Name) {
		case "":
			filePathToSave, err = e.concatDirFileExt(dirName, v.Name, fileExt)
			if err != nil {
				return err
			}
		default:
			filePathToSave = filepath.Join(dirName, v.Name)
		}

		err = e.createDir(filepath.Dir(filePathToSave))
		if err != nil {
			return err
		}

		f, err := os.OpenFile(filePathToSave, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer func() {
			closeErr := f.Close()
			if closeErr != nil {
				log.Println(closeErr)
			}
		}()
		if _, err := f.Write(v.Data); err != nil {
			return err
		}
	}

	return nil
}

func lookupFiles(input []byte, id string) []file {
	a := txtar.Parse(input)

	fileSet := make([]file, 0, 1+len(a.Files))
	fileSet = append(fileSet, file{
		Name: id,
		Data: changeDirName(a.Comment, []byte(id)),
	})
	for _, v := range a.Files {
		fileSet = append(fileSet, file{
			Name: v.Name,
			Data: v.Data,
		})
	}

	return fileSet
}

func changeDirName(data []byte, name []byte) []byte {
	tpl := []byte(os.Getenv("IMPORT_TPL"))
	return bytes.Replace(data, tpl, name, -1)
}

func (e *Executor) solArchive(solPath, id, lang string) (*os.File, error) {
	filesToTar := make([]string, 0)

	if err := filepath.Walk(solPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				filesToTar = append(filesToTar, path)
			}
			return nil
		}); err != nil {
		return nil, err
	}

	tarName := fmt.Sprintf("%s_%s", lang, id)
	if err := dockerapi.CreateTar(filesToTar, tarName); err != nil {
		return nil, err
	}

	return os.Open(filepath.Clean(tarName))
}

func (e *Executor) pathForExecSolutionInContainer(lang string, solutionID string) []string {
	var path []string
	path = append(path, fmt.Sprintf("%s%s", e.Config[lang].BinPath, solutionID))
	return path
}

func (e *Executor) commandForBuild(lang, solID string, binPath []string) []string {
	cmd := strings.Fields(e.Config[lang].BuildCommands)
	cmd = append(cmd, binPath...)
	cmd = append(cmd, filepath.Join(
		solID,
		fmt.Sprintf("%s%s", solID, e.Config[lang].FileExtension),
	))
	return cmd
}

func (e *Executor) createDir(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return os.Mkdir(name, os.ModePerm)
	}
	return nil
}

func (e *Executor) concatDirFileExt(dir, file, ext string) (string, error) {
	var f strings.Builder
	_, err := fmt.Fprintf(&f, "%s%s", file, ext)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, f.String()), nil
}
