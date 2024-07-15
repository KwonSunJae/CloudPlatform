package model

import (
	"database/sql"
)

type UserRepository interface {
	GetAllUser() (*[]UserRaw, error)
	GetOneUserByUUID(uuid string) (*UserRaw, error)
	GetOneUserByUserID(userID string) (*UserRaw, error)
	IsUserIDExit(userID string) (bool, error)
	DeleteOneUser(id string) (bool, error)
	InsertUser(n UserDto) (string, error)
	UpdateOneUser(id string, n UserDto) (bool, error)
	GetRoleByUUID(uuid string) (*UserRaw, error)
	AssignDB(db *sql.DB)
	GetEncryptedPW(userID string) (string, error)
}

type UserService interface {
	InitService() error
	GetAllUser() (*[]UserRaw, error)
	GetOneUserByUUID(uuid string) (*UserRaw, error)
	CreateUser(n UserDto) (string, error)
	ApproveUser(id string, role string, priority string) error
	IsUserIDExit(userID string) (bool, error)
	GetRoleByUUID(uuid string) (string, error)
}
