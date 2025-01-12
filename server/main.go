package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/matfire/pockets/server/docker"
	v1 "github.com/matfire/pockets/server/routers/v1"
	"github.com/matfire/pockets/server/utils"
	"github.com/spf13/viper"
)

func main() {
	log.Info("Setting up config")
	viper.SetConfigName("pockets")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("PKTS")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("ADMIN_USER", "admin")
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
	port := viper.GetInt("PORT")
	adminUser := viper.GetString("ADMIN_USER")
	adminPassword := viper.GetString("ADMIN_PASSWORD")
	app := utils.App{
		AdminUser:     adminUser,
		AdminPassword: adminPassword,
	}
	r.Mount("/v1", v1.GetRouter(&app))
	log.Info(fmt.Sprintf("listening on port %d", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
