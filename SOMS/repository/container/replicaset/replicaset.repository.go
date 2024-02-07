package replicaset

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type ReplicasetDto struct {
	ApiVersion                                   string
	Kind                                         string
	MetadataName                                 string
	SpecReplicas                                 string
	SpecSelectorMatchlabelsApp                   string
	SpecTemplateMetadataName                     string
	SpecTemplateMetadataLabelsApp                string
	SpecTemplateSpecContainersName               string
	SpecTemplateSpecContainersImage              string
	SpecTemplateSpecContainersPortsContainerport string
}

type ReplicasetRaw struct {
	Id                                           string
	ApiVersion                                   string
	Kind                                         string
	MetadataName                                 string
	SpecReplicas                                 string
	SpecSelectorMatchlabelsApp                   string
	SpecTemplateMetadataName                     string
	SpecTemplateMetadataLabelsApp                string
	SpecTemplateSpecContainersName               string
	SpecTemplateSpecContainersImage              string
	SpecTemplateSpecContainersPortsContainerport string
}

type ReplicasetRepository struct {
	DB *sql.DB
}

var Repository ReplicasetRepository

func (r *ReplicasetRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *ReplicasetRepository) InsertReplicaset(n ReplicasetDto) (sql.Result, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO replicaset
    (id, apiVersion, kind, metadataName, specReplicas, specSelectorMatchlabelsApp, specTemplateMetadataName, specTemplateMetadataLabelsApp, specTemplateSpecContainersName, specTemplateSpecContainersImage, specTemplateSpecContainersPortsContainerport)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `
	result, err := r.DB.Exec(query, id.String(), n.ApiVersion, n.Kind, n.MetadataName, n.SpecReplicas, n.SpecSelectorMatchlabelsApp, n.SpecTemplateMetadataName, n.SpecTemplateMetadataLabelsApp, n.SpecTemplateSpecContainersName, n.SpecTemplateSpecContainersImage, n.SpecTemplateSpecContainersPortsContainerport)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ReplicasetRepository) GetAllReplicaset() (*[]ReplicasetRaw, error) {
	var raws []ReplicasetRaw

	query := `SELECT * FROM replicaset`
	rows, err := r.DB.Query(query)

	for rows.Next() {
		var raw ReplicasetRaw
		rows.Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.MetadataName, &raw.SpecReplicas, &raw.SpecSelectorMatchlabelsApp, &raw.SpecTemplateMetadataName, &raw.SpecTemplateMetadataLabelsApp, &raw.SpecTemplateSpecContainersName, &raw.SpecTemplateSpecContainersImage, &raw.SpecTemplateSpecContainersPortsContainerport)

		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *ReplicasetRepository) GetOneReplicaset(id string) (*ReplicasetRaw, error) {
	var raw ReplicasetRaw

	query := `SELECT * FROM replicaset WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.MetadataName, &raw.SpecReplicas, &raw.SpecSelectorMatchlabelsApp, &raw.SpecTemplateMetadataName, &raw.SpecTemplateMetadataLabelsApp, &raw.SpecTemplateSpecContainersName, &raw.SpecTemplateSpecContainersImage, &raw.SpecTemplateSpecContainersPortsContainerport)

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

func (r *ReplicasetRepository) DeleteOneReplicaset(id string) (sql.Result, error) {
	query := `DELETE FROM replicaset WHERE id = ?`
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

func (r *ReplicasetRepository) UpdateOneReplicaset(id string, n ReplicasetDto) (sql.Result, error) {
	query := `
    UPDATE replicaset
    SET
        apiVersion = IFNULL(?, apiVersion),
    	kind = IFNULL(?, kind),
        metadataName = IFNULL(?, metadataName),
        specReplicas = IFNULL(?, specReplicas),
        specSelectorMatchlabelsApp = IFNULL(?, specSelectorMatchlabelsApp),
		specTemplateMetadataName = IFNULL(?, specTemplateMetadataName),
        specTemplateMetadataLabelsApp = IFNULL(?, specTemplateMetadataLabelsApp),
        specTemplateSpecContainersName = IFNULL(?, specTemplateSpecContainersName),
        specTemplateSpecContainersImage = IFNULL(?, specTemplateSpecContainersImage),
        specTemplateSpecContainersPortsContainerport = IFNULL(?, specTemplateSpecContainersPortsContainerport)
        
    WHERE
        id = ?
	`
	var apiVersion, kind, metadataName, specReplicas, specSelectorMatchlabelsApp, specTemplateMetadataName, specTemplateMetadataLabelsApp, specTemplateSpecContainersName, specTemplateSpecContainersImage, specTemplateSpecContainersPortsContainerport *string

	if n.ApiVersion != "" {
		apiVersion = &n.ApiVersion
	}

	if n.Kind != "" {
		kind = &n.Kind
	}

	if n.MetadataName != "" {
		metadataName = &n.MetadataName
	}

	if n.SpecReplicas != "" {
		specReplicas = &n.SpecReplicas
	}

	if n.SpecSelectorMatchlabelsApp != "" {
		specSelectorMatchlabelsApp = &n.SpecSelectorMatchlabelsApp
	}

	if n.SpecTemplateMetadataName != "" {
		specTemplateMetadataName = &n.SpecTemplateMetadataName
	}

	if n.SpecTemplateMetadataLabelsApp != "" {
		specTemplateMetadataLabelsApp = &n.SpecTemplateMetadataLabelsApp
	}

	if n.SpecTemplateSpecContainersName != "" {
		specTemplateSpecContainersName = &n.SpecTemplateSpecContainersName
	}

	if n.SpecTemplateSpecContainersImage != "" {
		specTemplateSpecContainersImage = &n.SpecTemplateSpecContainersImage
	}

	if n.SpecTemplateSpecContainersPortsContainerport != "" {
		specTemplateSpecContainersPortsContainerport = &n.SpecTemplateSpecContainersPortsContainerport
	}

	result, err := r.DB.Exec(query, apiVersion, kind, metadataName, specReplicas, specSelectorMatchlabelsApp, specTemplateMetadataName, specTemplateMetadataLabelsApp, specTemplateSpecContainersName, specTemplateSpecContainersImage, specTemplateSpecContainersPortsContainerport, id)

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
