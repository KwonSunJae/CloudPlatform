package deployment

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type DeploymentDto struct {
	ApiVersion                             				string
	Kind                                   				string
	Metadata_name                          				string
	Metadata_labels_app									string
	Spec_selector_matchLabels_app						string
	Spec_template_metadata_labels_app					string
	Spec_template_spec_hostname							string
	Spec_template_spec_subdomain						string
	Spec_template_spec_containers_image					string
	Spec_template_spec_containers_imagePullPolicy  	 	string
	Spec_template_spec_containers_name					string
	Spec_template_spec_containers_ports_containerPort	string

}

type DeploymentRaw struct {
	Id													string
	ApiVersion                             				string
	Kind                                   				string
	Metadata_name                          				string
	Metadata_labels_app									string
	Spec_selector_matchLabels_app						string
	Spec_template_metadata_labels_app					string
	Spec_template_spec_hostname							string
	Spec_template_spec_subdomain						string
	Spec_template_spec_containers_image					string
	Spec_template_spec_containers_imagePullPolicy  	 	string
	Spec_template_spec_containers_name					string
	Spec_template_spec_containers_ports_containerPort	string
}

type DeploymentRepository struct {
    DB *sql.DB
}

var Repository DeploymentRepository

func (r *DeploymentRepository) AssignDB(db *sql.DB) {
    r.DB = db
}

func (r *DeploymentRepository) createDeployment(n DeploymentDto) (sql.Result, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO Deployment
    (id, apiVersion, kind, metadata_name, metadata_labels_app, spec_selector_matchLabels_app, spec_template_metadata_labels_app, spec_template_spec_hostname, spec_template_spec_subdomain, spec_template_spec_containers_image, spec_template_spec_containers_imagePullPolicy, spec_template_spec_containers_name, spec_template_spec_containers_ports_containerPort)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `
	result, err := r.DB.Exec(query, id.String(), n.ApiVersion, n.Kind, n.Metadata_name, n.Metadata_labels_app, n.Spec_selector_matchLabels_app, n.Spec_template_metadata_labels_app, n.Spec_template_spec_hostname, n.Spec_template_spec_subdomain, n.Spec_template_spec_containers_image, n.Spec_template_spec_containers_imagePullPolicy, n.Spec_template_spec_containers_name, n.Spec_template_spec_containers_ports_containerPort)

	if err != nil {
		return nil, err
	}

	return result, nil
}


func (r *DeploymentRepository) GetAllDeployment() (*[]DeploymentRepository, error) {
	var raws []DeploymentRaw

	query := `SELECT * FROM Deployment`
	rows, err := r.DB.query(query)

	for rows.Next() {
		var raw DeploymentRaw
		rows.Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.Metadata_name, &raw.Metadata_labels_app, &raw.Spec_selector_matchLabels_app, &raw.Spec_template_metadata_labels_app, &raw.Spec_template_spec_hostname, &raw.Spec_template_spec_subdomain, &raw.Spec_template_spec_containers_image, &raw.Spec_template_spec_containers_imagePullPolicy, &raw.Spec_template_spec_containers_name, &raw.Spec_template_spec_containers_ports_containerPort)

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

	query := `SELECT * FROM Deployment WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&raw.Id, &raw.ApiVersion, &raw.Kind, &raw.Metadata_name, &raw.Metadata_labels_app, &raw.Spec_selector_matchLabels_app, &raw.Spec_template_metadata_labels_app, &raw.Spec_template_spec_hostname, &raw.Spec_template_spec_subdomain, &raw.Spec_template_spec_containers_image, &raw.Spec_template_spec_containers_imagePullPolicy, &raw.Spec_template_spec_containers_name, &raw.Spec_template_spec_containers_ports_containerPort)

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
    UPDATE Deployment
    SET
        apiVersion = IFNULL(?, apiVersion),
    	kind = IFNULL(?, kind),
        metadata_name = IFNULL(?, metadata_name),
        metatdata_labels_app = IFNULL(?, metadata_labels_app),
        spec_selector_matchLabels_app = IFNULL(?, spec_selector_matchLabels_app),
        spec_template_metadata_labels_app = IFNULL(?, spec_template_metadata_labels_app),
        spec_template_spec_hostname = IFNULL(?, spec_template_spec_hostname),
        spec_template_spec_subdomain = IFNULL(?, spec_template_spec_subdomain),
        spec_template_spec_containers_image = IFNULL(?, spec_template_spec_containers_image),
        spec_template_spec_containers_imagePullPolicy = IFNULL(?, spec_template_spec_containers_imagePullPolicy),
        spec_template_spec_containers_name = IFNULL(?, spec_template_spec_containers_name),
        spec_template_spec_containers_ports_containerPort = IFNULL(?, spec_template_spec_containers_ports_containerPort)

    WHERE
        id = ?
	`
	var apiVersion, kind, metadata_name, metadata_labels_app, spec_selector_matchLabels_app, spec_template_metadata_labels_app, spec_template_spec_hostname, spec_template_spec_subdomain, spec_template_spec_containers_image, spec_template_spec_containers_imagePullPolicy, spec_template_spec_containers_name, spec_template_spec_containers_ports_containerPort *string

	if n.ApiVersion != "" {
		apiVersion = &n.ApiVersion
	}

	if n.Kind != "" {
		kind = &n.Kind
	}

	if n.Metadata_name != "" {
		metadata_name = &n.Metadata_name
	}

	if n.Metadata_labels_app != "" {
		metadata_labels_app = &n.Metadata_labels_app
	}

	if n.Spec_selector_matchLabels_app != "" {
		spec_selector_matchLabels_app = &n.Spec_selector_matchLabels_app
	}

	if n.Spec_template_metadata_labels_app != "" {
		spec_template_metadata_labels_app = &n.Spec_template_metadata_labels_app
	}

	if n.Spec_template_spec_hostname != "" {
		spec_template_spec_hostname = &n.Spec_template_spec_hostname
	}

	if n.Spec_template_spec_subdomain != "" {
		spec_template_spec_subdomain = &n.Spec_template_spec_subdomain
	}

	if n.Spec_template_spec_containers_image != "" {
		spec_template_spec_containers_image = &n.Spec_template_spec_containers_image
	}

	if n.Spec_template_spec_containers_imagePullPolicy != "" {
		spec_template_spec_containers_imagePullPolicy = &n.Spec_template_spec_containers_imagePullPolicy
	}

	if n.Spec_template_spec_containers_name != "" {
		spec_template_spec_containers_name = &n.Spec_template_spec_containers_name
	}

	if n.Spec_template_spec_containers_ports_containerPort != "" {
		spec_template_spec_containers_ports_containerPort = &n.Spec_template_spec_containers_ports_containerPort
	}

	result, err := r.DB.Exec(query, apiVersion, kind, metadata_name, metadata_labels_app, spec_selector_matchLabels_app, spec_template_metadata_labels_app, spec_template_spec_hostname, spec_template_spec_subdomain, spec_template_spec_containers_image, spec_template_spec_containers_imagePullPolicy, spec_template_spec_containers_name, spec_template_spec_containers_ports_containerPort, id)

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
