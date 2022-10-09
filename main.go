package main

import (
	"fmt"
	"net/http"
)

// package level variable
const portNumber = ":8080"

func main() {
	// handler function to handle http request
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listening to localhost:8080
	_ = http.ListenAndServe(portNumber, nil)
}
