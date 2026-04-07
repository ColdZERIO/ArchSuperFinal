package repository

import (
	"Go/internal/models"
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddData(task *models.Task) (int, error) {
	query := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (:date, :title, :comment, :repeat);
	`

	res, err := r.db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (r *Repository) GetData() ([]models.Task, error) {
	query := `
	SELECT *
	FROM scheduler
	ORDER BY id DESC
	LIMIT 50;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Comment, &task.Date)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) UpdateDataByID(task models.Task) error {
	query := `
	UPDATE scheduler
	SET date = :date, title = :title, comment = :comment, repeat = :repeat
	WHERE id = :id;
	`

	_, err := r.db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)

	return err
}

func (r *Repository) GetDataByID(id int) (models.Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = :id;
	`

	row := r.db.QueryRow(query, sql.Named("id", id))

	var task models.Task

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
