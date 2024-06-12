package vm

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type VmDto struct {
	Name                  string
	FlavorID              string
	ExternalIP            string
	InternalIP            string
	SelectedOS            string
	UnionmountImage       string
	Keypair               string
	SelectedSecuritygroup string
	UUID                  string
	Status                string
}

type VmRaw struct {
	Id                    string
	Name                  string
	FlavorID              string
	ExternalIP            string
	InternalIP            string
	SelectedOS            string
	UnionmountImage       string
	Keypair               string
	SelectedSecuritygroup string
	UUID                  string
	Status                string
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
    INSERT INTO vm
    (id, name, flavorID, externalIP, internalIP,selectedOS, unionmountImage, keypair,selectedSecuritygroup ,uuid,status)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)
  `
	result, err := r.DB.Exec(query, id.String(), n.Name, n.FlavorID, n.ExternalIP, n.InternalIP, n.SelectedOS, n.UnionmountImage, n.Keypair, n.SelectedSecuritygroup, n.UUID, n.Status)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VmRepository) GetAllVm() (*[]VmRaw, error) {
	var raws []VmRaw

	query := `SELECT * FROM vm`
	rows, err := r.DB.Query(query)

	for rows.Next() {
		var raw VmRaw
		rows.Scan(&raw.Id, &raw.Name, &raw.FlavorID, &raw.ExternalIP, &raw.InternalIP, &raw.SelectedOS, &raw.UnionmountImage, &raw.Keypair, &raw.SelectedSecuritygroup, &raw.UUID, &raw.Status)
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

	query := `SELECT * FROM vm WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.Name, &raw.FlavorID, &raw.ExternalIP, &raw.InternalIP, &raw.SelectedOS, &raw.UnionmountImage, &raw.Keypair, &raw.SelectedSecuritygroup, &raw.UUID, &raw.Status)

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
	query := `DELETE FROM vm WHERE id = ?`
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
	query := `
    UPDATE vm
    SET
        name = IFNULL(?, name),
        flavorID = IFNULL(?, flavorID),
        externalIP = IFNULL(?, externalIP),
        internalIP = IFNULL(?, internalIP),
        selectedOS = IFNULL(?, selectedOS),
        unionmountImage = IFNULL(?, unionmountImage),
        keypair = IFNULL(?, keypair),
        selectedSecuritygroup = IFNULL(?, selectedSecuritygroup),
        uuid = IFNULL(?, uuid),
		status = IFNULL(?, status)
    WHERE
        id = ?
	`
	var name, flavorID, externalIP, internalIP, selectedOS, unionmountImage, keypair, selectedSecuritygroup, uuid, status *string

	if n.Name != "" {
		name = &n.Name
	}

	if n.FlavorID != "" {
		flavorID = &n.FlavorID
	}

	if n.ExternalIP != "" {
		externalIP = &n.ExternalIP
	}

	if n.InternalIP != "" {
		internalIP = &n.InternalIP
	}

	if n.SelectedOS != "" {
		selectedOS = &n.SelectedOS
	}

	if n.UnionmountImage != "" {
		unionmountImage = &n.UnionmountImage
	}

	if n.Keypair != "" {
		keypair = &n.Keypair
	}

	if n.SelectedSecuritygroup != "" {
		selectedSecuritygroup = &n.SelectedSecuritygroup
	}

	if n.UUID != "" {
		uuid = &n.UUID
	}

	if n.Status != "" {
		status = &n.Status
	}

	result, err := r.DB.Exec(query, name, flavorID, externalIP, internalIP, selectedOS, unionmountImage, keypair, selectedSecuritygroup, uuid, status, &id)

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
