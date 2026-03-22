package handlers

import (
	"Go/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Service interface {
	NextDate(time.Time, string, string) (string, error)
	NextDateAfter(time.Time, time.Time, string) (time.Time, error)
	AddTask(*models.Task) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "invalid json body", http.StatusBadRequest)
			return
		}

		err = h.service.AddTask(&task)
		if err != nil {
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
}
