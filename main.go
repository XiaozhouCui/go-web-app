package main

import (
	"errors"
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

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page and 2 + 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	// handler function to handle http request
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listening to localhost:8080
	_ = http.ListenAndServe(portNumber, nil)
}
