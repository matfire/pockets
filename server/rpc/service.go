package rpc

import (
	"context"

	"connectrpc.com/connect"

	"github.com/matfire/pockets/server/docker"
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
