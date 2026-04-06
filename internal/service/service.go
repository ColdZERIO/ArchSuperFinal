package services

import (
	"Go/internal/models"
	"encoding/json"
	"log"
	"net/http"
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
	SelectSQL(models.Task) ([]models.Task, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddTask(r *http.Request) error {
	var task models.Task

	decod := json.NewDecoder(r.Body)
	err := decod.Decode(&task)
	if err != nil {
		log.Println(err)
		return err
	}
	defer r.Body.Close()

	err = s.repo.InsertSQL(&task)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) TasksList() error {
	var task models.Task

	tasks, err := s.repo.SelectSQL(task)
	if err != nil {
		return err
	}

	encod := json.NewEncoder()
}
