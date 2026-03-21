package pkg

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func Init(dbFile string) error {

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("cant open sql file:", err)
		return err
	}
	defer db.Close()

	sqlFile, err := os.ReadFile("pkg/db/scheduler.sql")
	if err != nil {
		log.Fatal("cant read sql file:", err)
		return err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
