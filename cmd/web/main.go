package main

import (
	"fmt"
	"gowebapp/pkg/handlers"
	"net/http"
)

// package level variable
const portNumber = ":8080"

func main() {
	// handler function to handle http request
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listening to localhost:8080
	_ = http.ListenAndServe(portNumber, nil)
}
