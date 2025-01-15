package docker

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/matfire/pockets/server/utils"
	sharedv1 "github.com/matfire/pockets/shared/v1"

	getport "github.com/jsumners/go-getport"
)

func CreateContainer(data *sharedv1.CreateContainerRequest, app *utils.App) (*sharedv1.CreateContainerResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	images, err := cli.ImageList(context.Background(), image.ListOptions{
		All: true,
	})
	imgIdx := slices.IndexFunc(images, func(el image.Summary) bool {
		repoIdx := slices.IndexFunc(el.RepoTags, func(t string) bool {
			return t == fmt.Sprintf("pockets:%s", data.Version)
		})
		return repoIdx != -1
	})
	if imgIdx == -1 {
		return nil, errors.New("could not find requested image")
	}
	volume, err := cli.VolumeCreate(context.Background(), volume.CreateOptions{})
	port, err := getport.GetTcpPortForAddress("0.0.0.0")
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	res, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: fmt.Sprintf("pockets:%s", data.Version),
		ExposedPorts: nat.PortSet{
			"8080": struct{}{},
		},
		Labels: map[string]string{
			"pockets": "",
		},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: volume.Name,
				Target: "/pb/pb_data",
			},
		},
		PortBindings: nat.PortMap{
			// this should be the same as the exposed port
			"8080": []nat.PortBinding{
				{
					// wildcart locahost because docker shenanigans
					HostIP: "0.0.0.0",
					//TODO this should be randomly generated
					HostPort: fmt.Sprintf("%d", port.Port),
				},
			},
		},
	}, nil, nil, data.Name)

	err = cli.ContainerStart(context.Background(), res.ID, container.StartOptions{})
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	log.Info("Waiting for pocketbase to start properly")
	time.Sleep(5 * time.Second)
	adminCmd, err := cli.ContainerExecCreate(context.Background(), res.ID, container.ExecOptions{
		Cmd:          []string{"/pb/pocketbase", "superuser", "upsert", app.AdminUser, app.AdminPassword},
		AttachStdout: true,
		AttachStderr: true,
	})
	//CHECK AND BLOCK FOR COMMAND STATUS
	if err != nil {
		return nil, err
	}
	err = cli.ContainerExecStart(context.Background(), adminCmd.ID, container.ExecStartOptions{
		Detach: false,
		Tty:    false,
	})
	if err != nil {
		return nil, err
	}
	wg.Add(1)
	log.Info("Waiting for command to run")
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			log.Info("Checking for exec status")
			status, err := cli.ContainerExecInspect(context.Background(), adminCmd.ID)
			if err != nil {
				log.Error(err)
			}
			if !status.Running {
				ticker.Stop()
				log.Infof("Got status code %d", status.ExitCode)
				wg.Done()
			}
		}

	}()
	wg.Wait()
	containerInfo, err := cli.ContainerInspect(context.Background(), res.ID)
	if err != nil {
		return nil, err
	}
	return &sharedv1.CreateContainerResponse{
		Container: &sharedv1.Container{
			Id:     containerInfo.ID,
			Name:   containerInfo.Name,
			Status: containerInfo.State.Status,
		},
	}, nil
}

func GetContainers() *sharedv1.GetContainersResponse {
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
	res := sharedv1.GetContainersResponse{Containers: []*sharedv1.Container{}}
	for _, c := range containers {
		res.Containers = append(res.Containers, &sharedv1.Container{Id: c.ID, Status: c.Status, Name: c.Names[0]})
	}
	return &res

}

func StartContainer(data *sharedv1.StartContainerRequest) (*sharedv1.StartContainerResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	err = cli.ContainerStart(context.Background(), data.Id, container.StartOptions{})
	if err != nil {
		return nil, err
	}
	return &sharedv1.StartContainerResponse{
		Status: true,
	}, nil
}

func StopContainer(data *sharedv1.StopContainerRequest) (*sharedv1.StopContainerResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	err = cli.ContainerStop(context.Background(), data.Id, container.StopOptions{})
	if err != nil {
		return nil, err
	}
	return &sharedv1.StopContainerResponse{
		Status: false,
	}, nil
}

func DeleteContainer(data *sharedv1.DeleteContainerRequest) (*sharedv1.DeleteContainerResponse, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	err = cli.ContainerRemove(context.Background(), data.Id, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		return nil, err
	}
	return &sharedv1.DeleteContainerResponse{
		Status: true,
	}, nil
}
