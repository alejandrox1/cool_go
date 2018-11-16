package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const APIVersion = "1.38"

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion(APIVersion))
	if err != nil {
		log.Fatalf("Error initiating client: %v\n", err)
	}

	ctx := context.Background()
	listOpts := types.ContainerListOptions{
		All: true,
	}
	containers, err := cli.ContainerList(ctx, listOpts)
	if err != nil {
		log.Fatalf("Error listing containers: %v\n", err)
	}

	for _, container := range containers {
		fmt.Printf("%+v\n", container)
	}
}
