package docker

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/charmbracelet/huh/spinner"
	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/cli/rpc"
	sharedv1 "github.com/matfire/pockets/shared/v1"
)

type CreateRequestBody struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ImageCreateRequestBody struct {
	Version string `json:"version"`
}

type ImageCheckResponse struct {
	Exists bool `json:"exists"`
}

func Create(config *config.App, name string, version string) {
	imageExists := false
	client := rpc.GetRPCCLient(config)
	checkImage := func() bool {
		res, err := client.CheckImage(context.Background(), connect.NewRequest(&sharedv1.CheckImageRequest{
			Version: version,
		}))
		if err != nil {
			panic(err)
		}
		imageExists = res.Msg.Exists
		return res.Msg.Exists
	}

	createImage := func() {
		_, err := client.CreateImage(context.Background(), connect.NewRequest(&sharedv1.CreateImageRequest{
			Version: version,
		}))
		if err != nil {
			panic(err)
		}
	}

	//TODO customize retries & interval
	waitForImage := func() {
		iterations := 1
		for {
			if iterations > 5 {
				panic("image is taking too long")
			}

			exists := checkImage()
			if exists {
				break
			}
			iterations++
			time.Sleep(5 * time.Second)
		}
	}

	createContainer := func() {
		_, err := client.CreateContainer(context.Background(), connect.NewRequest(&sharedv1.CreateContainerRequest{
			Name:    name,
			Version: version,
		}))
		if err != nil {
			fmt.Printf("create request failed with error %v", err)
			return
		}
	}
	if err := spinner.New().Title("Checking for image existence...").Action(func() {
		checkImage()
	}).Run(); err != nil {
		fmt.Println(err)
	}
	if imageExists {
		if err := spinner.New().Title("Creating Container...").Action(createContainer).Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		if err := spinner.New().Title("Creating Image...").Action(createImage).Run(); err != nil {
			fmt.Println(err)
		}
		if err := spinner.New().Title("Waiting For Image...").Action(waitForImage).Run(); err != nil {
			fmt.Println(err)
		}
		if err := spinner.New().Title("Creating Container...").Action(createContainer).Run(); err != nil {
			fmt.Println(err)
		}
	}

}
