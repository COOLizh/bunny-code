package dockerapi

import (
	"context"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ExecuteSolutionInContainer executes user solution in container and sends testData to container's Stdin.
// It returns data, received from container's stdout ant execution time in milliseconds
func ExecuteSolutionInContainer(ctx context.Context, cli *client.Client, containerID string, testData string, cmd []string) ([]byte, int64, error) {
	containerCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          cmd,
	})

	if err != nil {
		return nil, 0, err
	}

	containerAttachResp, err := cli.ContainerExecAttach(ctx, containerCreateResp.ID, types.ExecConfig{
		Tty: true,
	})

	if err != nil {
		return nil, 0, err
	}

	start := time.Now()
	testDataReader := strings.NewReader(testData + "\n")
	_, err = testDataReader.WriteTo(containerAttachResp.Conn)

	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if err := containerAttachResp.CloseWrite(); err != nil {
			log.Println(err)
		}
	}()

	respReader := containerAttachResp.Reader

	var answerBytesLine []byte
	for {
		line, _, err := respReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, 0, err
		}
		answerBytesLine = line
	}
	timeSpent := time.Since(start).Milliseconds()

	return answerBytesLine, timeSpent, nil
}

// BuildSolutionInContainer builds user solution in container
func BuildSolutionInContainer(ctx context.Context, cli *client.Client, containerID string, cmd []string) ([]byte, error) {
	containerCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		AttachStderr: true,
		Cmd:          cmd,
	})

	if err != nil {
		return nil, err
	}
	containerAttachResp, err := cli.ContainerExecAttach(ctx, containerCreateResp.ID, types.ExecConfig{})

	if err != nil {
		return nil, err
	}
	respReader := containerAttachResp.Reader

	var answerBytesLine []byte
	for {
		line, _, err := respReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return answerBytesLine, err
		}
		answerBytesLine = line
	}
	return answerBytesLine, nil
}

// DeleteSolutionFromContainer removes user solution from container
func DeleteSolutionFromContainer(ctx context.Context, cli *client.Client, containerID, fileName string) error {
	cmd := strings.Fields(fmt.Sprintf("rm -fr %s", fileName))

	containerCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd: cmd,
	})
	if err != nil {
		return err
	}

	err = cli.ContainerExecStart(ctx, containerCreateResp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}

	return nil
}

// DeleteBuildFromContainer removes built user solution from container
func DeleteBuildFromContainer(ctx context.Context, cli *client.Client, containerID, binPath string, solName string) error {
	cmd := strings.Fields(fmt.Sprintf("rm -f %s", filepath.Join(binPath, solName)))

	containerCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd: cmd,
	})
	if err != nil {
		return err
	}

	err = cli.ContainerExecStart(ctx, containerCreateResp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}

	return nil
}
