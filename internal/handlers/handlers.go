package handlers

import "time"

type Service interface {
	NextDate(time.Time, string, string) (string, error)
}

type Handler struct {
	service Service
}

func NewService(service Service) *Handler {
	return &Handler{service: service}
}