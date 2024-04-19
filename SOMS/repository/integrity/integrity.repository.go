package integrity

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type IntegrityDto struct {
	RequestID string
	Type      string
	Action    string
	Result    string
	UserID    string
}

type IntegrityRaw struct {
	RequestID string
	Type      string
	Action    string
	Result    string
	UserID    string
}

type IntegrityRepository struct {
	DB *sql.DB
}

var Repository IntegrityRepository

func (r *IntegrityRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *IntegrityRepository) InsertIntegrity(n IntegrityDto) (string, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	query := `
	INSERT INTO integrity
	(request_id, type, action, result, user_id)
	VALUES (?, ?, ?, ?, ?)
  `
	rslt, err := r.DB.Exec(query, id.String(), n.Type, n.Action, n.Result, n.UserID)

	if err != nil {
		return "", err
	}
	if rslt == nil {
		return "", errors.New("NOT FOUND")
	}
	return id.String(), nil
}

func (r *IntegrityRepository) GetOneIntegrity(request_id string) (IntegrityRaw, error) {
	var raw IntegrityRaw

	query := `
	SELECT *
	FROM integrity
	WHERE request_id = ?
  `
	err := r.DB.QueryRow(query, request_id).Scan(&raw.RequestID, &raw.Type, &raw.Action, &raw.Result, &raw.UserID)

	if err != nil {
		return raw, err
	}

	return raw, nil
}

func (r *IntegrityRepository) GetIntegrityByUserID(user_id string) (*[]IntegrityRaw, error) {
	var raws []IntegrityRaw

	query := `
	SELECT *
	FROM integrity
	WHERE user_id = ?
  `
	rows, err := r.DB.Query(query, user_id)

	for rows.Next() {
		var raw IntegrityRaw
		rows.Scan(&raw.RequestID, &raw.Type, &raw.Action, &raw.Result, &raw.UserID)
		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *IntegrityRepository) GetAllIntegrity() (*[]IntegrityRaw, error) {
	var raws []IntegrityRaw

	query := `SELECT * FROM integrity`
	rows, err := r.DB.Query(query)

	for rows.Next() {
		var raw IntegrityRaw
		rows.Scan(&raw.RequestID, &raw.Type, &raw.Action, &raw.Result, &raw.UserID)
		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *IntegrityRepository) UpdateIntegrity(n IntegrityDto) error {
	query := `
	UPDATE integrity
	SET type = ?, action = ?, result = ?, user_id = ?
	WHERE request_id = ?
  `
	rslt, err := r.DB.Exec(query, n.Type, n.Action, n.Result, n.UserID, n.RequestID)

	if err != nil {
		return err
	}
	if rslt == nil {
		return errors.New("NOT FOUND")
	}
	return nil
}

func (r *IntegrityRepository) DeleteIntegrity(request_id string) error {
	query := `DELETE FROM integrity WHERE request_id = ?`
	rslt, err := r.DB.Exec(query, request_id)

	if err != nil {
		return err
	}
	if rslt == nil {
		return errors.New("NOT FOUND")
	}
	return nil
}
