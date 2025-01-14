package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/matfire/pockets/server/docker"
	v1 "github.com/matfire/pockets/server/routers/v1"
	"github.com/matfire/pockets/server/rpc"
	"github.com/matfire/pockets/server/utils"
	"github.com/matfire/pockets/shared"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	r := chi.NewRouter()
	log.Info("Creating network...")
	docker.CreateNetwork("pockets")
	log.Info("Network Created! Or maybe it already existed")
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprint("hello world")))
	})
	port := viper.GetString("PORT")
	adminUser := viper.GetString("ADMIN_USER")
	adminPassword := viper.GetString("ADMIN_PASSWORD")
	app := utils.App{
		AdminUser:     adminUser,
		AdminPassword: adminPassword,
	}
	r.Mount("/v1", v1.GetRouter(&app))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	shared.RegisterPocketsServer(grpcServer, rpc.NewPocketsService())
	grpcServer.Serve(lis)
	log.Info(fmt.Sprintf("listening on port %s", port))
}
