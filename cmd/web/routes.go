package main

import (
	"gowebapp/pkg/config"
	"gowebapp/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	// create a multiplexer (mux)
	mux := pat.New() // pat is a 3rd-party package

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
