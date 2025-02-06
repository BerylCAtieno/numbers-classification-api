package main

import (
	"fmt"
	"net/http"
	"numbers-classifier-api/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/api/classify-number", handlers.ClassifyNumberHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server running on port 8080...")
	server.ListenAndServe()
}
