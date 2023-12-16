package service

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type ServiceDto struct {
	ApiVersion            string
	Kind                  string
	Metadata_name         string
	Spec_ports_port       string
	Spec_ports_protocol   string
	Spec_ports_targetPort string
	Spec_selector_app     string
}

type ServiceRaw struct {
	Id                    string
	ApiVersion            string
	Kind                  string
	Metadata_name         string
	Spec_ports_port       string
	Spec_ports_protocol   string
	Spec_ports_targetPort string
	Spec_selector_app     string
}

type ServiceRepository struct {
	DB *sql.DB
}

var Repository ServiceRepository

func (r *ServiceRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *ServiceRepository) InsertService(n ServiceDto) (sql.Result, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO service
    (id, apiVersion, kind, metadata_name, spec_ports_port, spec_ports_protocol, spec_ports_targetPort, spec_selector_app)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
  `
	result, err := r.DB.Exec(query, id.String(), n.ApiVersion, n.Kind, n.Metadata_name, n.Spec_ports_port, n.Spec_ports_protocol, n.Spec_ports_targetPort, n.Spec_selector_app)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ServiceRepository) GetAllService() (*[]ServiceRaw, error) {
	var raws []ServiceRaw

	query := `SELECT * FROM service`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var raw ServiceRaw
		rows.Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.Metadata_name, &raw.Spec_ports_port, &raw.Spec_ports_protocol, &raw.Spec_ports_targetPort, &raw.Spec_selector_app)

		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *ServiceRepository) GetOneService(id string) (*ServiceRaw, error) {
	var raw ServiceRaw

	query := `SELECT * FROM service WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.Metadata_name, &raw.Spec_ports_port, &raw.Spec_ports_protocol, &raw.Spec_ports_targetPort, &raw.Spec_selector_app)

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

func (r *ServiceRepository) DeleteOneService(id string) (sql.Result, error) {
	query := `DELETE FROM service WHERE id = ?`
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

func (r *ServiceRepository) UpdateOneService(id string, n ServiceDto) (sql.Result, error) {
	query := `
    UPDATE service
    SET
        apiVersion = IFNULL(?, apiVersion),
    	kind = IFNULL(?, kind),
        metadata_name = IFNULL(?, metadata_name),
        spec_ports_port = IFNULL(?, spec_ports_port),
        spec_ports_protocol = IFNULL(?, spec_ports_protocol),
        spec_ports_targetPort = IFNULL(?, spec_ports_targetPort),
        spec_selector_app = IFNULL(?, spec_selector_app),
        
    WHERE
        id = ?
	`
	var apiVersion, kind, metadata_name, spec_ports_port, spec_ports_protocol, spec_ports_targetPort, spec_selector_app *string

	if n.ApiVersion != "" {
		apiVersion = &n.ApiVersion
	}

	if n.Kind != "" {
		kind = &n.Kind
	}

	if n.Metadata_name != "" {
		metadata_name = &n.Metadata_name
	}

	if n.Spec_ports_port != "" {
		spec_ports_port = &n.Spec_ports_port
	}

	if n.Spec_ports_protocol != "" {
		spec_ports_protocol = &n.Spec_ports_protocol
	}

	if n.Spec_ports_targetPort != "" {
		spec_ports_targetPort = &n.Spec_ports_targetPort
	}

	if n.Spec_selector_app != "" {
		spec_selector_app = &n.Spec_selector_app
	}

	result, err := r.DB.Exec(query, apiVersion, kind, metadata_name, spec_ports_port, spec_ports_protocol, spec_ports_targetPort, spec_selector_app, id)

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
