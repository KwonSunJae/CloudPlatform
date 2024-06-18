package service

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type ServiceDto struct {
	// ClusterIP
	ApiVersion          string
	Kind                string
	MetadataName        string
	SpecType            string
	SpecSelectorApp     string
	SpecPortsProtocol   string // 배열로 나중에 바꿔야함
	SpecPortsPort       string
	SpecPortsTargetport string

	// NodePort
	SpecPortsNodeport string

	// LoadBalancer
	SpecSelectorType string
	SpecClusterIP    string

	// ExternalName
	SpecExternalname string

	UUID   string
	Status string
}

type ServiceRaw struct {
	// ClusterIP
	Id                  string
	ApiVersion          string
	Kind                string
	MetadataName        string
	SpecType            string
	SpecSelectorApp     string
	SpecPortsProtocol   string // 배열로 나중에 바꿔야함
	SpecPortsPort       string
	SpecPortsTargetport string

	// NodePort
	SpecPortsNodeport string

	// LoadBalancer
	SpecSelectorType string
	SpecClusterIP    string

	// ExternalName
	SpecExternalname string

	UUID   string
	Status string
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
    (id, apiVersion, kind, metadataName, specType, specSelectorApp, specPortsProtocol, specPortsPort, specPortsTargetport, specPortsNodeport, specSelectorType, specClusterIP, specExternalname, uuid, status)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?)
  `
	result, err := r.DB.Exec(query, id.String(), n.ApiVersion, n.Kind, n.MetadataName, n.SpecType, n.SpecSelectorApp, n.SpecPortsProtocol, n.SpecPortsPort, n.SpecPortsTargetport, n.SpecPortsNodeport, n.SpecSelectorType, n.SpecClusterIP, n.SpecExternalname, n.UUID, n.Status)

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
		rows.Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.MetadataName, &raw.SpecType, &raw.SpecSelectorApp, &raw.SpecPortsProtocol, &raw.SpecPortsPort, &raw.SpecPortsTargetport, &raw.SpecPortsNodeport, &raw.SpecSelectorType, &raw.SpecClusterIP, &raw.SpecExternalname, &raw.UUID, &raw.Status)

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
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.MetadataName, &raw.SpecType, &raw.SpecSelectorApp, &raw.SpecPortsProtocol, &raw.SpecPortsPort, &raw.SpecPortsTargetport, &raw.SpecPortsNodeport, &raw.SpecSelectorType, &raw.SpecClusterIP, &raw.SpecExternalname, &raw.UUID, &raw.Status)

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
        metadataName = IFNULL(?, metadataName),
        specType = IFNULL(?, specType),
        specSelectorApp = IFNULL(?, specSelectorApp),
        specPortsProtocol = IFNULL(?, specPortsProtocol),
        specPortsPort = IFNULL(?, specPortsPort),
		specPortsTargetport = IFNULL(?, specPortsTargetport),
		specPortsNodeport = IFNULL(?, specPortsNodeport),
		specSelectorType = IFNULL(?, specSelectorType),
		specClusterIP = IFNULL(?, specClusterIP),
		specExternalname  = IFNULL(?, specExternalname),
		uuid = IFNULL(?, uuid),
		status = IFNULL(?, status)
        
    WHERE
        id = ?
	`
	var apiVersion, kind, metadataName, specType, specSelectorApp, specPortsProtocol, specPortsPort, specPortsTargetport, specPortsNodeport, specSelectorType, specClusterIP, specExternalname, uuid, status *string

	if n.ApiVersion != "" {
		apiVersion = &n.ApiVersion
	}

	if n.Kind != "" {
		kind = &n.Kind
	}

	if n.MetadataName != "" {
		metadataName = &n.MetadataName
	}

	if n.SpecType != "" {
		specType = &n.SpecType
	}

	if n.SpecSelectorApp != "" {
		specSelectorApp = &n.SpecSelectorApp
	}

	if n.SpecPortsProtocol != "" {
		specPortsProtocol = &n.SpecPortsProtocol
	}

	if n.SpecPortsPort != "" {
		specPortsPort = &n.SpecPortsPort
	}

	if n.SpecPortsTargetport != "" {
		specPortsTargetport = &n.SpecPortsTargetport
	}

	if n.SpecPortsNodeport != "" {
		specPortsNodeport = &n.SpecPortsNodeport
	}

	if n.SpecSelectorType != "" {
		specSelectorType = &n.SpecSelectorType
	}

	if n.SpecClusterIP != "" {
		specClusterIP = &n.SpecClusterIP
	}

	if n.SpecExternalname != "" {
		specExternalname = &n.SpecExternalname
	}

	if n.UUID != "" {
		uuid = &n.UUID
	}

	if n.Status != "" {
		status = &n.Status
	}

	result, err := r.DB.Exec(query, apiVersion, kind, metadataName, specType, specSelectorApp, specPortsProtocol, specPortsPort, specPortsTargetport, specPortsNodeport, specSelectorType, specClusterIP, specExternalname, uuid, status, id)

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
