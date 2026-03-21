package services

import (
	"time"
)

type Todo struct {
	id      int
	date    time.Time
	title   string
	comment string
	repeat  string
}

type Repository interface {
	SelectDB()
	InsertDB()
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
