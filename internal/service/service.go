package services

import (
	"Go/internal/models"
	"log"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type Repository interface {
	InsertSQL(*models.Task) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddTask(task *models.Task) error {
	err := s.repo.InsertSQL(task)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
