package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

const APIVersion = "1.38"

func main() {
	// Cretae Docker client from environment options.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion(APIVersion))
	if err != nil {
		log.Fatalf("Error instantiating client: %v\n", err)
	}

	ctx := context.Background()
	// Pull image.
	imagePull, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		log.Fatalf("Error pulling image: %v\n", err)
	}
	io.Copy(os.Stdout, imagePull)
	defer imagePull.Close()

	// Create container from image.
	cont := &container.Config{
		Image: "alpine:3.8",
		Cmd:   []string{"echo", "Hello, world!"},
	}
	host := &container.HostConfig{
		AutoRemove: true,
	}
	net := &network.NetworkingConfig{}
	resp, err := cli.ContainerCreate(ctx, cont, host, net, "go-client-container")
	if err != nil {
		log.Fatalf("Error creating container: %v\n", err)
	}

	// Start container.
	startOpts := types.ContainerStartOptions{}
	if err := cli.ContainerStart(ctx, resp.ID, startOpts); err != nil {
		log.Fatalf("Error starting container: %v\n", err)
	}

    // Everything after thins point emulates the process of running a container
    // in the foreground. To replicate the detach behaviour one simply starts
    // the container and lets it go.

	// Get copy logs to stdout and stderr.
	logOpts := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	}
	out, err := cli.ContainerLogs(ctx, resp.ID, logOpts)
	if err != nil {
		log.Fatalf("Error getting logs: %v\n", err)
	}
	// If no using TTY then stream is multiplexed and stream data will contain
	// extra bits. To avoid demultiplexing the stream use stdcopy.
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	defer out.Close()

	// Wait for container to be removed.
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionRemoved)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	case status := <-statusCh:
		log.Printf("Status - error: %v code: %v\n", status.Error, status.StatusCode)
	}

}
