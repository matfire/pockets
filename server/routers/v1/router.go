package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/matfire/pockets/server/routers/v1/handlers/docker"
	"github.com/matfire/pockets/server/utils"
)

func GetRouter(app *utils.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/status", docker.GetStatus)
	r.Post("/create", docker.CreateContainer)
	r.Post("/image/new", docker.CreateImage)
	r.Get("/image/check/{version}", docker.CheckImage)
	r.Put("/start/{containerId}", docker.StartContainer)
	r.Put("/stop/{containerId}", docker.StopContainer)
	r.Delete("/delete/{containerId}", docker.DeleteContainer)
	return r
}
