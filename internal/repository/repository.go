package repository

import (
	"Go/internal/models"
	"database/sql"
	"errors"
	"log"
)

const limitRows int = 50

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
	SELECT id, date, title, comment, repeat
	FROM scheduler
	ORDER BY date ASC
	LIMIT :limitRows;
	`

	rows, err := r.db.Query(query, sql.Named("limitRows", limitRows))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) UpdateData(task *models.Task) error {
	query := `
	UPDATE scheduler
	SET date = :date, title = :title, comment = :comment, repeat = :repeat
	WHERE id = :id;
	`

	res, err := r.db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID),
	)

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Задача не найдена")
	}

	return nil
}

func (r *Repository) GetDataByID(id string) (*models.Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = :id;
	`

	row := r.db.QueryRow(query, sql.Named("id", id))

	task := &models.Task{}

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Задача не найдена")
		}
		return nil, err
	}

	return task, nil
}

func (r *Repository) DeleteTask(id string) error {
	query := `
	DELETE FROM scheduler
	WHERE id = :id
	`

	res, err := r.db.Exec(query, sql.Named("id", id))
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("задача не найдена")
	}

	return nil
}
