package main

import (
	"fmt"
	"net/http"
)

// package level variable
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// print in browser (http response)
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page and 2 + 2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	// handler function to handle http request
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// listening to localhost:8080
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
