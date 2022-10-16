package main

import (
	"gowebapp/pkg/config"
	"gowebapp/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// create a multiplexer (mux)
	mux := chi.NewRouter()

	// use Recoverer middleware
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
