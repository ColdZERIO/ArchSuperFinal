package services

import (
	"Go/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Repository interface {
	AddData(*models.Task) (int, error)
	GetData() ([]models.Task, error)
	GetDataByID(int) (models.Task, error)
	UpdateDataByID(models.Task) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddTask(r *http.Request) (int, error) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := s.repo.AddData(&task)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (s *Service) GetTasksList(w http.ResponseWriter) error {
	tasks, err := s.repo.GetData()
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(map[string]any{
		"tasks": tasks,
	})

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

	task, err := s.repo.GetDataByID(id)
	if err != nil {
		return err
	}

	// newTime, err := nextDate(time.Now(), task.Date, task.Repeat)
	// if err != nil {
	// 	return err
	// }

	//task.Date = newTime

	err = s.repo.UpdateDataByID(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) NextDate(now, date, repeat string) (int, error) {
	 timeNow, err := time.Parse(layout, now)
	 if err != nil {
		return 0, err
	 }

	 newTime, err := nextDate(timeNow, date, repeat)
	 if err != nil {
		return 0, err
	 }

	 return newTime, nil
}
