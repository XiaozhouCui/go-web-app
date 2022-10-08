package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// package level variable
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {
	// handler function to handle http request
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// listening to localhost:8080
	_ = http.ListenAndServe(portNumber, nil)
}
