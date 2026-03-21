package main

import (
	pkg "Go/pkg/db"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()

	port := getEnv("TODO_PORT", ":7540")
	dbFile := getEnv("TODO_DBFILE", "scheduler.db")

	log.Println(port)
	webDir := "./web"

	err = pkg.Init(dbFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	srv := chi.NewRouter()

	srv.Handle("/*", http.FileServer(http.Dir(webDir)))

	log.Printf("Server START http://localhost%s/\n", port)
	log.Fatal(http.ListenAndServe(port, srv))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
