package repository

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func Init(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("cant open sql file:", err)
		return nil, err
	}

	sqlFile, err := os.ReadFile("pkg/db/scheduler.sql")
	if err != nil {
		log.Fatal("cant read sql file:", err)
		return nil, err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
