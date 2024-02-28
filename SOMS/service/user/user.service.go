package user

import (
	"errors"
	"fmt"
	"os"
	"soms/repository"
	"soms/repository/user"
	"soms/util/encrypt"
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

func (s *UserService) GetOneUser(userID string) (*user.UserRaw, error) {
	raw, err := s.Repository.GetOneUser(userID)

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

func (s *UserService) UserIDValidate(userID string) (bool, error) {
	isExist, err := s.Repository.IsUserIDExit(userID)
	if isExist {
		return true, nil
	} else {
		return false, err
	}

}
func (s *UserService) UpdateUser(userID string, n user.UserDto) error {
	_, err := s.Repository.UpdateOneUser(userID, n)

	return err
}

func (s *UserService) DeleteUser(id string) error {

	_, err2 := s.Repository.DeleteOneUser(id)
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *UserService) UserLogin(userID string, plainPW string) (bool, error) {

	encryptedPW, err := s.Repository.GetEncryptedPW(userID)
	if err != nil {
		return false, err
	}
	secretKey := os.Getenv("SECRET")
	if secretKey == "" {
		fmt.Println("SECRET key is not set.")
	}
	hasher := encrypt.NewPasswordHasher(secretKey)

	rslt := hasher.ComparePassword(plainPW, encryptedPW)
	if rslt != nil {
		return false, errors.New("pw is not correct")
	}
	return true, nil
}
