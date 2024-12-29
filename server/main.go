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
)

func main() {
	r := chi.NewRouter()
	log.Info("Creating network...")
	docker.CreateNetwork("pockets")
	log.Info("Network Created! Or maybe it already existed")
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprint("hello world")))
	})
	r.Mount("/v1", v1.GetRouter())
	log.Info("listening on port 3000")
	http.ListenAndServe(":3000", r)
}
