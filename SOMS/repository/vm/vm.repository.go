package vm

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type VmDto struct {
	Title   string
	Author  string
	Content string
}

type VmRaw struct {
	Id      string
	Author  string
	Title   string
	Content string
}

type VmRepository struct {
	DB *sql.DB
}

var Repository VmRepository

func (r *VmRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *VmRepository) InsertVm(n VmDto) (sql.Result, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO Vm
    (id, title, author, content)
    VALUES (?, ?, ?, ?)
  `
	result, err := r.DB.Exec(query, id.String(), n.Title, n.Author, n.Content)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VmRepository) GetAllVm() (*[]VmRaw, error) {
	var raws []VmRaw

	query := `SELECT * FROM Vm`
	rows, err := r.DB.Query(query)

	for rows.Next() {
		var raw VmRaw
		rows.Scan(&raw.Id, &raw.Title, &raw.Author, &raw.Content)
		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *VmRepository) GetOneVm(id string) (*VmRaw, error) {
	var raw VmRaw

	query := `SELECT * FROM Vm WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.Title, &raw.Author, &raw.Content)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("NOT FOUND")
		} else {
			return nil, err
		}
	} else {
		return &raw, nil
	}
}

func (r *VmRepository) DeleteOneVm(id string) (sql.Result, error) {
	query := `DELETE FROM Vm WHERE id = ?`
	result, err := r.DB.Exec(query, id)

	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("NOT FOUND")
	}

	return result, nil
}

func (r *VmRepository) UpdateOneVm(id string, n VmDto) (sql.Result, error) {
	query := `UPDATE Vm SET title = IFNULL(?, title), author = IFNULL(?, author), content = IFNULL(?, content) WHERE id = ?`
	var title, author, content *string

	if n.Title != "" {
		title = &n.Title
	}

	if n.Author != "" {
		author = &n.Author
	}

	if n.Content != "" {
		content = &n.Content
	}

	result, err := r.DB.Exec(query, title, author, content, id)

	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("NOT FOUND")
	}

	return result, nil
}
