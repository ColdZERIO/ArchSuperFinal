package services

import (
	"Go/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
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
	SelectByID(int) (models.Task, error)
	UpdateSQL(models.Task) error
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

func (s *Service) TasksList(w http.ResponseWriter) error {
	var task models.Task

	tasks, err := s.repo.SelectSQL(task)
	if err != nil {
		return err
	}

	encod := json.NewEncoder(w)
	err = encod.Encode(tasks)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Service) UpdateTask(r *http.Request) error {
	queryID := r.URL.Query().Get("id")
	if queryID == "" {
		return errors.New("id is empty")
	}

	id, err := strconv.Atoi(queryID)
	if err != nil {
		return errors.New("invalid id format")
	}

	task, err := s.repo.SelectByID(id)
	if err != nil {
		return err
	}

	err = s.repo.UpdateSQL(task)
	if err != nil {
		return err
	}

	return nil
}
