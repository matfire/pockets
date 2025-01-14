package rpc

import (
	"context"

	"github.com/matfire/pockets/server/docker"
	"github.com/matfire/pockets/shared"
)

type PocketsService struct {
	shared.UnsafePocketsServer
}

func NewPocketsService() *PocketsService {
	return &PocketsService{}
}

func (s *PocketsService) GetContainers(ctx context.Context, req *shared.GetContainersRequest) (*shared.ContainerList, error) {

	list := docker.GetContainers()
	return &list, nil
}
