package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var todos = []string{}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			log.Println("GET /todos")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todos)
		case http.MethodPost:
			log.Println("POST /todos")
			p, _ := io.ReadAll(r.Body)
			todos = append(todos, string(p))
			w.WriteHeader(http.StatusCreated)
		}
	})

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
