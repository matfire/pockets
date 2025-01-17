package docker

import (
	"context"
	"strings"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/matfire/pockets/server/rpc"
	sharedv1 "github.com/matfire/pockets/shared/v1"
)

func setupContainerCreate(name string, version string, t *testing.T) func(t *testing.T) {
	client := rpc.PocketsServer{}

	_, err := client.CreateImage(context.Background(), connect.NewRequest(&sharedv1.CreateImageRequest{
		Version: version,
	}))
	if err != nil {
		t.Fatalf("ContainerCreate setup failed: %v", err)
	}
	t.Logf("Got name %s", name)

	return func(t *testing.T) {
		DeleteImage(version, t)
	}
}

func TestContainerCreate(t *testing.T) {
	client := rpc.PocketsServer{}
	name, _ := uuid.NewV6()
	version := "v0.24.3"
	teardown := setupContainerCreate(name.String(), version, t)
	defer teardown(t)
	res, err := client.CreateContainer(context.Background(), connect.NewRequest(&sharedv1.CreateContainerRequest{
		Name:    name.String(),
		Version: version,
	}))
	if err != nil {
		t.Fatalf("ContainerCreate failed: %v", err)
	}
	if strings.Contains(res.Msg.Container.Name, name.String()) == false {
		t.Fatalf("ContainerCreate failed: container name is different")
	}
}
