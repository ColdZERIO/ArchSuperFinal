package handlers

import (
	"log"
	"net/http"
	"time"
)

type Service interface {
	NextDate(time.Time, string, string) (string, error)
	NextDateAfter(time.Time, time.Time, string) (time.Time, error)
	AddTask(r *http.Request) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := h.service.AddTask(r)
		if err != nil {
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
}
