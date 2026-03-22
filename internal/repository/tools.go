package repository

import (
	"Go/internal/models"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertSQL(task *models.Task) error {
	query := `INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (:date, :title, :comment, :repeat)`

	_, err := r.db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)

	return err
}
