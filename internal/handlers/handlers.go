package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"Go/internal/models"
)

type Service interface {
	AddTask(r *http.Request) (int, error)
	GetTasksList(w http.ResponseWriter) error
	GetTask(r *http.Request) (*models.Task, error)
	UpdateTask(r *http.Request) error
	NextDate(now, date, repeat string) (int, error)
	TaskDone(id string) error
	DeleteTask(id string) error
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
		id, err := h.service.AddTask(r)
		if err != nil {
			responseJson(err, w, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(map[string]any{
			"id": id,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case http.MethodGet:
		if r.URL.Path == "/api/task" {
			id := r.URL.Query().Get("id")
			if id == "" {
				responseJson(errors.New("Не указан идентификатор"), w, http.StatusBadRequest)
				return
			}

			task, err := h.service.GetTask(r)
			if err != nil {
				responseJson(err, w, http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
			return
		}

		err := h.service.GetTasksList(w)
		if err != nil {
			responseJson(err, w, http.StatusInternalServerError)
			return
		}
		return

	case http.MethodPut:
		err := h.service.UpdateTask(r)
		if err != nil {
			responseJson(err, w, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {
	now := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	newDate, err := h.service.NextDate(now, date, repeat)
	if err != nil {
		responseJson(err, w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) TaskDone(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error": "метод не поддерживается",
		})
		return
	}

	taskID := r.URL.Query().Get("id")
	err := h.service.TaskDone(taskID)
	if err != nil {
		responseJson(err, w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{})
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")

	err := h.service.DeleteTask(taskID)
	if err != nil {
		responseJson(err, w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{})
}
