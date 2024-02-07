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
	_, err = createServiceTable(db)
	if err != nil {
		return nil, err
	}
	_, err = createDeploymentTable(db)
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

func createServiceTable(db *sql.DB) (sql.Result, error) {
	query := `
	CREATE TABLE service (
		id TEXT PRIMARY KEY,
		apiVersion TEXT,
		kind TEXT,
		metadataName TEXT,
		specType TEXT,
		specSelectorApp TEXT,
		specPortsProtocol TEXT,
		specPortsPort TEXT,
		specPortsTargetport TEXT,
		specPortsNodeport TEXT,
		specSelectorType TEXT,
		specClusterIP TEXT,
		metadataNamespace TEXT,
		specExternalname TEXT
	)
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func createDeploymentTable(db *sql.DB) (sql.Result, error) {
	query := `
	CREATE TABLE deployment (
		id SERIAL PRIMARY KEY,
		apiVersion TEXT,
		kind TEXT,
		metadata_name TEXT,
		metadata_labels_app TEXT,
		spec_selector_matchLabels_app TEXT,
		spec_template_metadata_labels_app TEXT,
		spec_template_spec_hostname TEXT,
		spec_template_spec_subdomain TEXT,
		spec_template_spec_containers_image TEXT,
		spec_template_spec_containers_imagePullPolicy TEXT,
		spec_template_spec_containers_name TEXT,
		spec_template_spec_containers_ports_containerPort TEXT
	)
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func createUserTable(db *sql.DB) (sql.Result, error) {
	query := `
	CREATE TABLE user (
		id TEXT PRIMARY KEY,
		password TEXT,
		email TEXT,
		schoolNum TEXT,
		detailRole TEXT,
		isLocked BOOLEAN,
	)
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}
