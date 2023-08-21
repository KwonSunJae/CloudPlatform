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

	_, err = createVmTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createVmTable(db *sql.DB) (sql.Result, error) {
	query := `
  CREATE TABLE vm (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    flavorID TEXT NOT NULL,
    externalIP TEXT NOT NULL,
    internalIP TEXT NOT NULL,
    selectedOS TEXT NOT NULL,
    unionmountImage TEXT NOT NULL,
    keypair TEXT NOT NULL,
    selectedSecuritygroup TEXT NOT NULL,
    userID TEXT NOT NULL
)
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}
