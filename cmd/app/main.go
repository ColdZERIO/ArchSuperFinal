package main

import (
	"Go/internal/handlers"
	"Go/internal/repository"
	services "Go/internal/service"
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
	if err != nil {
		log.Fatal(err)
	}

	port := getEnv("TODO_PORT", ":7540")
	dbFile := getEnv("TODO_DBFILE", "scheduler.db")

	webDir := "./web"

	db, err := pkg.Init(dbFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	svc := services.NewService(repo)
	hand := handlers.NewHandler(svc)

	srv := chi.NewRouter()
	srv.Handle("/*", http.FileServer(http.Dir(webDir)))

	log.Printf("Server START http://localhost%s/\n", port)

	srv.Post("/api/task", hand.TaskHandler)
	srv.Get("/api/task", hand.TaskHandler)

	log.Fatal(http.ListenAndServe(port, srv))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
