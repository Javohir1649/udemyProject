package main

import (
	"fmt"
	"net/http"
	"udemyProject/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Homepage)
	http.HandleFunc("/about", handlers.Aboutpage)

	fmt.Printf("starting server on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
