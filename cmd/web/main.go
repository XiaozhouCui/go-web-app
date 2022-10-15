package main

import (
	"fmt"
	"gowebapp/pkg/config"
	"gowebapp/pkg/handlers"
	"gowebapp/pkg/render"
	"log"
	"net/http"
)

// package level variable
const portNumber = ":8080"

func main() {
	var app config.AppConfig

	// create a map (tc) to cache templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	// handler function to handle http request
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listening to localhost:8080
	_ = http.ListenAndServe(portNumber, nil)
}
