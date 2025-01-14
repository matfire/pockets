package rpc

import (
	"net/http"

	"github.com/matfire/pockets/cli/config"
	"github.com/matfire/pockets/shared/v1/sharedv1connect"
)

func GetRPCCLient(config *config.App) sharedv1connect.PocketsServiceClient {
	return sharedv1connect.NewPocketsServiceClient(
		http.DefaultClient,
		config.Endpoint,
	)
}
