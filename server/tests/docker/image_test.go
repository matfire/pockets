package docker

import (
	"context"
	"fmt"
	"testing"

	"connectrpc.com/connect"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/matfire/pockets/server/rpc"
	sharedv1 "github.com/matfire/pockets/shared/v1"
)

func setupImageCreate(version string, t *testing.T) func(t *testing.T) {
	DeleteImage(version, t)
	return func(t *testing.T) {}
}

func DeleteImage(version string, t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Could not connect to docker: %v", err)
	}
	data, err := cli.ImageList(context.Background(), image.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("reference", fmt.Sprintf("pockets:%s", version))),
	})
	if len(data) == 0 {
		t.Logf("Could not find image for version %s", version)
	}
	cli.ImageRemove(context.Background(), data[0].ID, image.RemoveOptions{
		Force: true,
	})
}

func checkImage(version string, t *testing.T) bool {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Could not connect to docker: %v", err)
	}
	data, err := cli.ImageList(context.Background(), image.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("reference", fmt.Sprintf("pockets:%s", version))),
	})
	return len(data) > 0
}

func TestImageCreation(t *testing.T) {

	version := "v0.24.3"
	teardown := setupImageCreate(version, t)
	defer teardown(t)
	server := rpc.PocketsServer{}
	ctx := context.Background()
	_, err := server.CreateImage(ctx, connect.NewRequest(&sharedv1.CreateImageRequest{Version: version}))
	if err != nil {
		t.Fatalf("CreateImage failed: %v", err)
	}
	exists := checkImage(version, t)
	if exists == false {
		t.Fatalf("Image not created")
	}
}
