package deployment

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type DeploymentDto struct {
	ApiVersion                                   string
	Kind                                         string
	MetadataName                                 string
	MetadataLabelsApp                            string
	SpecReplicas                                 string
	SpecSelectorMatchlabelsApp                   string
	SpecTemplateMetadataLabelsApp                string
	SpecTemplateSpecContainersName               string
	SpecTemplateSpecContainersImage              string
	SpecTemplateSpecContainersPortsContainerport string
}

type DeploymentRaw struct {
	Id                                           string
	ApiVersion                                   string
	Kind                                         string
	MetadataName                                 string
	MetadataLabelsApp                            string
	SpecReplicas                                 string
	SpecSelectorMatchlabelsApp                   string
	SpecTemplateMetadataLabelsApp                string
	SpecTemplateSpecContainersName               string
	SpecTemplateSpecContainersImage              string
	SpecTemplateSpecContainersPortsContainerport string
}

type DeploymentRepository struct {
	DB *sql.DB
}

var Repository DeploymentRepository

func (r *DeploymentRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *DeploymentRepository) InsertDeployment(n DeploymentDto) (sql.Result, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO deployment
    (id, apiVersion, kind, metadataName, metadataLabelsApp, specReplicas, specSelectorMatchlabelsApp, specTemplateMetadataLabelsApp, specTemplateSpecContainersName, specTemplateSpecContainersImage, specTemplateSpecContainersPortsContainerport)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `
	result, err := r.DB.Exec(query, id.String(), n.ApiVersion, n.Kind, n.MetadataName, n.MetadataLabelsApp, n.SpecReplicas, n.SpecSelectorMatchlabelsApp, n.SpecTemplateMetadataLabelsApp, n.SpecTemplateSpecContainersName, n.SpecTemplateSpecContainersImage, n.SpecTemplateSpecContainersPortsContainerport)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *DeploymentRepository) GetAllDeployment() (*[]DeploymentRaw, error) {
	var raws []DeploymentRaw

	query := `SELECT * FROM deployment`
	rows, err := r.DB.Query(query)

	for rows.Next() {
		var raw DeploymentRaw
		rows.Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.MetadataName, &raw.MetadataLabelsApp, &raw.SpecReplicas, &raw.SpecSelectorMatchlabelsApp, &raw.SpecTemplateMetadataLabelsApp, &raw.SpecTemplateSpecContainersName, &raw.SpecTemplateSpecContainersImage, &raw.SpecTemplateSpecContainersPortsContainerport)

		raws = append(raws, raw)
	}

	if err != nil {
		return nil, err
	} else {
		return &raws, nil
	}
}

func (r *DeploymentRepository) GetOneDeployment(id string) (*DeploymentRaw, error) {
	var raw DeploymentRaw

	query := `SELECT * FROM deployment WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.MetadataName, &raw.MetadataLabelsApp, &raw.SpecReplicas, &raw.SpecSelectorMatchlabelsApp, &raw.SpecTemplateMetadataLabelsApp, &raw.SpecTemplateSpecContainersName, &raw.SpecTemplateSpecContainersImage, &raw.SpecTemplateSpecContainersPortsContainerport)

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

func (r *DeploymentRepository) DeleteOneDeployment(id string) (sql.Result, error) {
	query := `DELETE FROM deployment WHERE id = ?`
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

func (r *DeploymentRepository) UpdateOneDeployment(id string, n DeploymentDto) (sql.Result, error) {
	query := `
    UPDATE deployment
    SET
        apiVersion = IFNULL(?, apiVersion),
    	kind = IFNULL(?, kind),
        metadataName = IFNULL(?, metadataName),
		metadataLabelsApp = IFNULL(?, metadataLabelsApp),
        specReplicas = IFNULL(?, specReplicas),
        specSelectorMatchlabelsApp = IFNULL(?, specSelectorMatchlabelsApp),
        specTemplateMetadataLabelsApp = IFNULL(?, specTemplateMetadataLabelsApp),
        specTemplateSpecContainersName = IFNULL(?, specTemplateSpecContainersName),
        specTemplateSpecContainersImage = IFNULL(?, specTemplateSpecContainersImage),
        specTemplateSpecContainersPortsContainerport = IFNULL(?, specTemplateSpecContainersPortsContainerport)
        
    WHERE
        id = ?
	`
	var apiVersion, kind, metadataName, metadataLabelsApp, specReplicas, specSelectorMatchlabelsApp, specTemplateMetadataLabelsApp, specTemplateSpecContainersName, specTemplateSpecContainersImage, specTemplateSpecContainersPortsContainerport *string

	if n.ApiVersion != "" {
		apiVersion = &n.ApiVersion
	}

	if n.Kind != "" {
		kind = &n.Kind
	}

	if n.MetadataName != "" {
		metadataName = &n.MetadataName
	}

	if n.MetadataLabelsApp != "" {
		metadataLabelsApp = &n.MetadataLabelsApp
	}

	if n.SpecReplicas != "" {
		specReplicas = &n.SpecReplicas
	}

	if n.SpecSelectorMatchlabelsApp != "" {
		specSelectorMatchlabelsApp = &n.SpecSelectorMatchlabelsApp
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

	result, err := r.DB.Exec(query, apiVersion, kind, metadataName, metadataLabelsApp, specReplicas, specSelectorMatchlabelsApp, specTemplateMetadataLabelsApp, specTemplateSpecContainersName, specTemplateSpecContainersImage, specTemplateSpecContainersPortsContainerport, id)

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
