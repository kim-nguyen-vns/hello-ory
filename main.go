package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler responds to HTTP requests with "Hello, World!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s, User: %s\n", r.URL.Path, r.Header.Get("X-User-ID"))
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register handler function for the root URL path
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
