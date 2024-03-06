package user

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type UserDto struct {
	Name        string
	UserID      string
	EncryptedPW string
	Role        string
	Spot        string
	Priority    string
}

type UserRaw struct {
	Id          string
	Name        string
	UserID      string
	EncryptedPW string
	Role        string
	Spot        string
	Priority    string
}

type UserRepository struct {
	DB *sql.DB
}

var Repository UserRepository

func (r *UserRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *UserRepository) InsertUser(n UserDto) (sql.Result, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO user
    (id, name, userID, encryptedPW, role, spot, priority)
    VALUES (?, ?, ?, ?, ?, ?, ?)
  `
	result, err := r.DB.Exec(query, id.String(), n.Name, n.UserID, n.EncryptedPW, n.Role, n.Spot, n.Priority)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) GetAllUser() (*[]UserRaw, error) {
	var raws []UserRaw

	query := `SELECT * FROM user`
	rows, err := r.DB.Query(query)

	for rows.Next() {
		var raw UserRaw
		rows.Scan(&raw.Id, &raw.Name, &raw.UserID, &raw.EncryptedPW, &raw.Role, &raw.Spot, &raw.Priority)
		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *UserRepository) GetOneUser(userID string) (*UserRaw, error) {
	var raw UserRaw

	query := `SELECT * FROM user WHERE userID = ?`
	err := r.DB.QueryRow(query, userID).Scan(&raw.Id, &raw.Name, &raw.UserID, &raw.EncryptedPW, &raw.Role, &raw.Spot, &raw.Priority)
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
func (r *UserRepository) IsUserIDExit(userID string) (bool, error) {
	var raw UserRaw

	query := `SELECT * FROM user WHERE userID = ?`
	err := r.DB.QueryRow(query, userID).Scan(&raw.Id, &raw.Name, &raw.UserID, &raw.EncryptedPW, &raw.Role, &raw.Spot, &raw.Priority)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return true, errors.New("NOT FOUND")
		} else {
			return false, err
		}
	} else {
		return false, nil
	}
}

func (r *UserRepository) DeleteOneUser(id string) (sql.Result, error) {
	query := `DELETE FROM user WHERE id = ?`
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

func (r *UserRepository) UpdateOneUser(id string, n UserDto) (sql.Result, error) {
	query := `
    UPDATE user
    SET
        name = IFNULL(?, name),
        encryptedPW = IFNULL(?, encryptedPW),
        role = IFNULL(?, role),
        spot = IFNULL(?, spot),
        priority = IFNULL(?, priority),
    WHERE
        userID = ?
	`
	var name, userID, encryptedPW, role, spot, priority *string

	if n.Name != "" {
		name = &n.Name
	}

	if n.UserID != "" {
		userID = &n.UserID
	}

	if n.EncryptedPW != "" {
		encryptedPW = &n.EncryptedPW
	}

	if n.Role != "" {
		role = &n.Role
	}

	if n.Spot != "" {
		spot = &n.Spot
	}

	if n.Priority != "" {
		priority = &n.Priority
	}

	result, err := r.DB.Exec(query, name, encryptedPW, role, spot, priority, userID)

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
func (r *UserRepository) GetEncryptedPW(userID string) (string, error) {
	var encryptedPW string
	//print(userID + "login check\n")
	query := `SELECT encryptedPW FROM user WHERE userID = ?`
	err := r.DB.QueryRow(query, userID).Scan(&encryptedPW)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", errors.New("NOT FOUND")
		} else {
			return "", err
		}
	} else {
		return encryptedPW, nil
	}
}
