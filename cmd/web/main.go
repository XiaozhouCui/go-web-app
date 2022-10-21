package main

import (
	"fmt"
	"gowebapp/pkg/config"
	"gowebapp/pkg/handlers"
	"gowebapp/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

// package level variable
const portNumber = ":8080"

func main() {
	// create a global app config variable
	var app config.AppConfig

	// initialise session
	session := scs.New()
	session.Lifetime = 24 * time.Hour

	// create a map (tc) to cache all templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// store the template cache to the app config
	app.TemplateCache = tc
	app.UseCache = false

	// declare repository variable
	repo := handlers.NewRepo(&app)
	// pass the repo back to the handlers
	handlers.NewHandlers(repo)
	// give render package access to the app config
	render.NewTemplates(&app) // reference to app using pointer

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// create a server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
