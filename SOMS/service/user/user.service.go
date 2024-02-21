package user

import (
	"soms/repository"
	"soms/repository/user"
)

type UserService struct {
	Repository *user.UserRepository
}

var Service UserService

func (s *UserService) InitService() error {
	db, err := repository.OpenWithMemory()

	if err != nil {
		return err
	}

	s.Repository = &user.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *UserService) GetAllUser() (*[]user.UserRaw, error) {
	raws, err := s.Repository.GetAllUser()

	return raws, err
}

func (s *UserService) GetOneUser(id string) (*user.UserRaw, error) {
	raw, err := s.Repository.GetOneUser(id)

	return raw, err
}

func (s *UserService) CreateUser(n user.UserDto) error {

	// Generate Openstack Account

	// Generate Terraform Repositroy

	// Generate K8s Repository

	_, DBSaveErr := s.Repository.InsertUser(n)
	if DBSaveErr != nil {
		return DBSaveErr
	}
	return nil
}

func (s *UserService) UpdateUser(id string, n user.UserDto) error {
	_, err := s.Repository.UpdateOneUser(id, n)

	return err
}

func (s *UserService) DeleteUser(id string) error {

	_, err2 := s.Repository.DeleteOneUser(id)
	if err2 != nil {
		return err2
	}
	return nil
}
