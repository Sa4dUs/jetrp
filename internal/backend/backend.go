package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, Golang!",
		Time:    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", handler)

	fmt.Printf("Backend server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
