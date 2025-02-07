package main

import (
	"fmt"
	"log"
	"net/http"
	"numbers-classifier-api/handlers"
)

func main() {
	// Define the endpoint handler
	http.HandleFunc("/api/classify-number", handlers.ClassifyNumberHandler)

	fmt.Println("Server running on 0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
