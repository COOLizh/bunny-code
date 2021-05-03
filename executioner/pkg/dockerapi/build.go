package dockerapi

import (
	"archive/tar"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	tarNameBase = "buildContext.tar"
)

// BuildImage builds image from docker file
func BuildImage(ctx context.Context, cli *client.Client, dockerfilePath, lang string) error {
	filePathDockerfile, err := filepath.Abs(dockerfilePath)
	if err != nil {
		return err
	}
	filesToTar := []string{filePathDockerfile}
	err = CreateTar(filesToTar, tarNameBase)
	if err != nil {
		return err
	}
	// #nosec G304
	imageBuildContext, err := os.Open(tarNameBase)
	if err != nil {
		return err
	}

	imageBuildOptions := types.ImageBuildOptions{
		Dockerfile: filePathDockerfile,
		Tags:       []string{lang},
	}

	imageBuildResponse, err := cli.ImageBuild(ctx, imageBuildContext, imageBuildOptions)
	if err != nil {
		return err
	}
	defer func() {
		if err = imageBuildResponse.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	defer func() {
		if err = os.Remove(tarNameBase); err != nil {
			log.Println(err)
		}
	}()

	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

// CreateTar creates tar from file
func CreateTar(files []string, tarName string) error {
	outputFile, err := os.Create(tarName)
	if err != nil {
		return err
	}
	defer func() {
		if err = outputFile.Close(); err != nil {
			log.Println(err)
		}
	}()

	tarWriter := tar.NewWriter(outputFile)
	defer func() {
		if err = tarWriter.Close(); err != nil {
			log.Println(err)
		}
	}()

	for _, file := range files {
		err = AddToTar(tarWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddToTar adds files to tar
func AddToTar(tarWriter *tar.Writer, fileName string) error {
	// #nosec G304
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(fileStat, fileStat.Name())
	if err != nil {
		return err
	}

	header.Name = fileName

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}
