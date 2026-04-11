package services

import (
	"Go/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"
)

type Repository interface {
	AddData(task *models.Task) (int, error)
	GetData() ([]models.Task, error)
	UpdateData(task *models.Task) error
	GetDataByID(id string) (*models.Task, error)
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

	if task.Title == "" {
		return 0, errors.New("Не указан заголовок задачи")
	}

	if task.Date == "" {
		task.Date = time.Now().Format(layout)
	}

	if err = repeatCheck(task.Repeat); err != nil {
		return 0, err
	}

	taskDate, err := time.Parse(layout, task.Date)
	if err != nil {
		return 0, err
	}

	if time.Now().After(taskDate) {
		task.Date = time.Now().Format(layout)
	}

	id, err := s.repo.AddData(&task)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (s *Service) GetTasksList(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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

func (s *Service) GetTask(r *http.Request) (*models.Task, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, errors.New("Не указан идентификатор")
	}

	task, err := s.repo.GetDataByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Service) GetTaskByID(id string) (*models.Task, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("invalid id")
	}

	return s.repo.GetDataByID(id)
}

func (s *Service) UpdateTask(r *http.Request) error {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return err
	}

	if task.ID == "" {
		return errors.New("Не указан идентификатор")
	}

	if task.Title == "" {
		return errors.New("Не указан заголовок задачи")
	}

	if task.Date == "" {
		task.Date = time.Now().Format(layout)
	}

	if err := repeatCheck(task.Repeat); err != nil {
		return err
	}

	taskDate, err := time.Parse(layout, task.Date)
	if err != nil {
		return err
	}

	if time.Now().After(taskDate) {
		task.Date = time.Now().Format(layout)
	}

	return s.repo.UpdateData(&task)
}

func (s *Service) NextDate(now, date, repeat string) (int, error) {
	var timeNow time.Time
	var err error

	if now == "" {
		timeNow = time.Now()
	} else {
		timeNow, err = time.Parse(layout, now)
		if err != nil {
			return 0, err
		}
	}

	newTime, err := nextDate(timeNow, date, repeat)
	if err != nil {
		return 0, err
	}

	return newTime, nil
}
