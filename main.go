package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handler function to handle http request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// print in browser (http response)
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		// print in terminal
		fmt.Println(fmt.Sprintf("Number of bytes writter: %d", n))
	})

	// listening to localhost:8080
	_ = http.ListenAndServe(":8080", nil)
}
