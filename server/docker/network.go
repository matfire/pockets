package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func CreateNetwork(name string) {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	res, err := cli.NetworkCreate(context.Background(), name, network.CreateOptions{})
	fmt.Printf("%v", res.Warning)
}
