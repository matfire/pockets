package rpc

import (
	"context"

	"connectrpc.com/connect"

	"github.com/matfire/pockets/server/docker"
	"github.com/matfire/pockets/server/utils"
	sharedv1 "github.com/matfire/pockets/shared/v1"
	"github.com/matfire/pockets/shared/v1/sharedv1connect"
)

type PocketsServer struct {
	sharedv1connect.UnimplementedPocketsServiceHandler
}

func (s *PocketsServer) GetContainers(ctx context.Context, req *connect.Request[sharedv1.GetContainersRequest]) (*connect.Response[sharedv1.GetContainersResponse], error) {
	data := docker.GetContainers()
	res := connect.NewResponse(data)
	return res, nil
}

func (s *PocketsServer) CreateContainer(ctx context.Context, req *connect.Request[sharedv1.CreateContainerRequest]) (*connect.Response[sharedv1.CreateContainerResponse], error) {
	data, err := docker.CreateContainer(req.Msg, utils.GetConfig())
	res := connect.NewResponse(data)
	return res, err
}
