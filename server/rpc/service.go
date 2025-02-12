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

func (s *PocketsServer) CheckImage(ctx context.Context, req *connect.Request[sharedv1.CheckImageRequest]) (*connect.Response[sharedv1.CheckImageResponse], error) {
	data, err := docker.CheckImage(req.Msg)
	res := connect.NewResponse(data)
	return res, err
}

func (s *PocketsServer) CreateImage(ctx context.Context, req *connect.Request[sharedv1.CreateImageRequest]) (*connect.Response[sharedv1.CreateImageResponse], error) {
	data, err := docker.CreateImage(req.Msg)
	res := connect.NewResponse(data)
	return res, err
}

func (s *PocketsServer) StartContainer(ctx context.Context, req *connect.Request[sharedv1.StartContainerRequest]) (*connect.Response[sharedv1.StartContainerResponse], error) {
	data, err := docker.StartContainer(req.Msg)
	res := connect.NewResponse(data)
	return res, err
}
func (s *PocketsServer) StopContainer(ctx context.Context, req *connect.Request[sharedv1.StopContainerRequest]) (*connect.Response[sharedv1.StopContainerResponse], error) {
	data, err := docker.StopContainer(req.Msg)
	res := connect.NewResponse(data)
	return res, err
}
func (s *PocketsServer) DeleteContainer(ctx context.Context, req *connect.Request[sharedv1.DeleteContainerRequest]) (*connect.Response[sharedv1.DeleteContainerResponse], error) {
	data, err := docker.DeleteContainer(req.Msg)
	res := connect.NewResponse(data)
	return res, err
}
