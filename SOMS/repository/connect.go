package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenWithMemory() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	_, err = createNewsTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createNewsTable(db *sql.DB) (sql.Result, error) {
	query := `
  CREATE TABLE news (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    content TEXT NOT NULL
  )
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}
