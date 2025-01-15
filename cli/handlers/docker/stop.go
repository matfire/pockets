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

func Stop(config *config.App, name string) {
	var err error
	client := rpc.GetRPCCLient(config)
	stopContainer := func() {
		_, errr := client.StopContainer(context.Background(), connect.NewRequest(&sharedv1.StopContainerRequest{Id: name}))
		err = errr
	}
	if spinnerErr := spinner.New().Title("Stopping Container...").Action(stopContainer).Run(); spinnerErr != nil {
		fmt.Println(spinnerErr)
	}
	if err != nil {
		fmt.Printf("create request failed with error %v", err)
		return
	}

}
