package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := getEnv("TODO_PORT", "")
	webDir := "./web"

	srv := chi.NewRouter()

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	log.Println("Server START http://localhost:7540/")
	log.Fatal(http.ListenAndServe(port, srv))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
