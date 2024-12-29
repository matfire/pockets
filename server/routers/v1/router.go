package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/matfire/pockets/server/routers/v1/handlers/docker"
)

func GetRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/status", docker.GetStatus)
	r.Post("/create", docker.CreateContainer)
	return r
}
