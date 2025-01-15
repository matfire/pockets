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

func setupImageCheck(t *testing.T) func(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Could not connect to docker: %v", err)
	}
	data, err := cli.ImageList(context.Background(), image.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("reference", "pockets:v0.24.3")),
	})
	if len(data) > 0 {
		t.Log("Found image; deleting")
		cli.ImageRemove(context.Background(), data[0].ID, image.RemoveOptions{
			Force: true,
		})
	}
	return func(t *testing.T) {
		t.Log("teardown")
	}
}

func setupImageCreate(t *testing.T) func(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Could not connect to docker: %v", err)
	}
	data, err := cli.ImageList(context.Background(), image.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("reference", "pockets:v0.24.3")),
	})
	if len(data) > 0 {
		t.Log("Found image; deleting")
		cli.ImageRemove(context.Background(), data[0].ID, image.RemoveOptions{
			Force: true,
		})
	}
	return func(t *testing.T) {}
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

func TestImageCheck(t *testing.T) {
	teardown := setupImageCheck(t)
	defer teardown(t)
	server := rpc.PocketsServer{}
	version := "v0.24.3"
	response, err := server.CheckImage(context.Background(), connect.NewRequest(&sharedv1.CheckImageRequest{Version: version}))
	if err != nil {
		t.Fatalf("CheckImage failed: %v", err)
	}
	if response.Msg.Exists {
		t.Fatalf("CheckImage failed: image exists")
	}

}

func TestImageCreation(t *testing.T) {
	teardown := setupImageCreate(t)
	defer teardown(t)
	server := rpc.PocketsServer{}
	version := "v0.24.3"
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
