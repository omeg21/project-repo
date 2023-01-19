package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/omeg21/project-repo/pkg/config"
	"github.com/omeg21/project-repo/pkg/controller"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Use(SessionLoad)

	mux.Get("/", controller.Repo.Home)
	mux.Get("/about", controller.Repo.About)
	return mux
}
