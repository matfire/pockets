package docker

import (
	"context"
	"encoding/json"
	"fmt"
	getport "github.com/jsumners/go-getport"
	"io"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/matfire/pockets/server/routers/v1/types"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("got %d containers", len(containers))
	data, err := json.Marshal(types.StatusResponse{Message: "ok"})
	if err != nil {
		panic("could not marshal empty json")
	}
	w.Write(data)
}

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	var body types.ContainerCreateBody
	//TODO add validation so that container name is valid (no spaces, probably more rules...)

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &body)
	if err != nil {
		panic(err)
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	port, err := getport.GetTcpPortForAddress("0.0.0.0")
	if err != nil {
		fmt.Printf("%v", err)
		panic("could not get port")
	}
	res, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: "pockets:0.23",
		ExposedPorts: nat.PortSet{
			"8080": struct{}{},
		},
		Labels: map[string]string{
			"pockets": "",
		},
	}, &container.HostConfig{
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
	}, nil, nil, body.Name)

	cli.ContainerStart(context.Background(), res.ID, container.StartOptions{})
	data, err := json.Marshal(res)
	w.WriteHeader(201)
	w.Write(data)
}
