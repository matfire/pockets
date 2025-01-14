package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/matfire/pockets/shared"
)

func GetContainers() shared.ContainerList {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// TODO filter by network and/or by labels
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("label", "pockets")),
	})
	if err != nil {
		panic(err)
	}
	res := shared.ContainerList{Containers: []*shared.ContainerObject{}}
	for _, c := range containers {
		res.Containers = append(res.Containers, &shared.ContainerObject{Id: c.ID, Status: c.Status, Name: c.Names[0]})
	}
	return res

}
