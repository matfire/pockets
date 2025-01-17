package docker

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/charmbracelet/huh/spinner"
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/rpc"
	sharedv1 "github.com/matfire/pockets/shared/v1"
)

func Start(config *config.App, name string) {
	var err error
	client := rpc.GetRPCCLient(config)
	startContainer := func() {
		_, errr := client.StartContainer(context.Background(), connect.NewRequest(&sharedv1.StartContainerRequest{
			Id: name,
		}))
		err = errr
	}
	if spinnerErr := spinner.New().Title("Starting Container...").Action(startContainer).Run(); spinnerErr != nil {
		fmt.Println(spinnerErr)
	}
	if err != nil {
		fmt.Printf("create request failed with error %v", err)
		return
	}

}
