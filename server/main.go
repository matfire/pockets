package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/matfire/pockets/server/docker"
	"github.com/matfire/pockets/server/rpc"
	"github.com/matfire/pockets/shared/v1/sharedv1connect"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log.Info("Setting up config")
	viper.SetConfigName("pockets")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("PKTS")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("ADMIN_USER", "admin@example.com")
	viper.SetDefault("ADMIN_PASSWORD", "test1234")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("config file not found, using env & default values")
		} else {
			panic(err)
		}
	}
	log.Info("Creating network...")
	docker.CreateNetwork("pockets")
	log.Info("Network Created! Or maybe it already existed")
	port := viper.GetInt("PORT")
	pocketServer := rpc.PocketsServer{}
	path, handler := sharedv1connect.NewPocketsServiceHandler(&pocketServer)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	log.Info(fmt.Sprintf("listening on port %d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), h2c.NewHandler(mux, &http2.Server{}))
}
