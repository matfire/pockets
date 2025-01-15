package docker

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/rpc"
	sharedv1 "github.com/matfire/pockets/shared/v1"
)

type Container struct {
	Name string `json:"name"`
}

type ListResponse struct {
	Containers []Container `json:"containers"`
}

func List(config *config.App) {
	fmt.Print("listing containers")
	client := rpc.GetRPCCLient(config)
	var data *sharedv1.GetContainersResponse
	getContainers := func() {
		res, err := client.GetContainers(context.Background(), connect.NewRequest(&sharedv1.GetContainersRequest{}))
		if err != nil {
			panic(err)
		}
		data = res.Msg
	}
	if err := spinner.New().Title("Fetching containers...").Action(getContainers).Run(); err != nil {
		fmt.Println(err)
	}
	t := table.New().Headers("Name")

	for _, container := range data.Containers {
		t.Row(container.Name)
	}
	fmt.Println(t)
}
