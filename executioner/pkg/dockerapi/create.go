package dockerapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"gitlab.com/greenteam1/executioner/pkg/models"
)

// StartContainers starts all needed containers
func StartContainers(cli *client.Client, config map[string]*models.BuildConfig) (map[string]string, error) {
	var err error

	containersIDs := make(map[string]string)

	for key := range config {
		var containerID string
		containerID, err = StartContainer(cli, config[key])
		if err != nil {
			return nil, err
		}
		log.Println(containerID)

		containersIDs[key] = containerID
	}
	return containersIDs, nil
}

// StartContainer starts container from image
func StartContainer(cli *client.Client, config *models.BuildConfig) (string, error) {
	ctx := context.Background()
	dockerFilePath := config.DockerFilePath
	err := BuildImage(ctx, cli, dockerFilePath, config.ContainerName)
	if err != nil {
		return "", nil
	}
	containerID, err := CreateContainer(ctx, cli, config.ContainerName)
	if err != nil {
		return containerID, err
	}
	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return containerID, err
	}

	return containerID, nil
}

// CreateContainer creates container with specified language
func CreateContainer(ctx context.Context, cli *client.Client, imageName string) (string, error) {
	timestamp := fmt.Sprint(time.Now().Unix())
	containerName := imageName + timestamp
	containerConfig := container.Config{
		Image:        imageName,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin:    true,
		Tty:          true,
	}
	containerCreateResponse, err := cli.ContainerCreate(
		ctx,
		&containerConfig,
		nil,
		nil,
		containerName,
	)
	if err != nil {
		return "", err
	}

	return containerCreateResponse.ID, nil
}
