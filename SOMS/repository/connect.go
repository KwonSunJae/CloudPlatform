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
	_, err = createReplicasetTable(db)
	if err != nil {
		return nil, err
	}
	_, err = createUserTable(db)
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
	uuid TEXT NOT NULL,
	status TEXT NOT NULL,
    FOREIGN KEY(uuid) REFERENCES user (id)
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
		specExternalname TEXT,
		uuid TEXT NOT NULL,
		status TEXT NOT NULL,
		FOREIGN KEY(uuid) REFERENCES user (id)
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
		id TEXT PRIMARY KEY,
		apiVersion TEXT,
		kind TEXT,
		metadataName TEXT,
		metadataLabelsApp TEXT,
		specReplicas TEXT,
		specSelectorMatchlabelsApp TEXT,
		specTemplateMetadataLabelsApp TEXT,
		specTemplateSpecContainersName TEXT,
		specTemplateSpecContainersImage TEXT,
		specTemplateSpecContainersPortsContainerport TEXT,
		uuid TEXT NOT NULL,
		status TEXT NOT NULL,
		FOREIGN KEY(uuid) REFERENCES user (id)
	)
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func createReplicasetTable(db *sql.DB) (sql.Result, error) {
	query := `
	CREATE TABLE replicaset (
		id TEXT PRIMARY KEY,
		apiVersion TEXT,
		kind TEXT,
		metadataName TEXT,
		specReplicas TEXT,
		specSelectorMatchlabelsApp TEXT,
		specTemplateMetadataName TEXT,
		specTemplateMetadataLabelsApp TEXT,
		specTemplateSpecContainersName TEXT,
		specTemplateSpecContainersImage TEXT,
		specTemplateSpecContainersPortsContainerport TEXT
		uuid TEXT NOT NULL,
		status TEXT NOT NULL,
		FOREIGN KEY(uuid) REFERENCES user (id)
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
		name TEXT ,
		userID TEXT UNIQUE,
		encryptedPW TEXT,
		role TEXT,
		spot TEXT,
		priority TEXT
	)
  `

	result, err := db.Exec(query)

	if err != nil {
		return nil, err
	}

	return result, nil
}
