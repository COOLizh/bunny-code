// Package dockerapi provides functions to interact with docker engine API
package dockerapi

import (
	"os"

	"github.com/docker/docker/client"
)

// NewDockerClient returns new client for communicating with Docker's API
func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClient(
		os.Getenv("DOCKER_API_HOST"),
		os.Getenv("DOCKER_API_VER"),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
