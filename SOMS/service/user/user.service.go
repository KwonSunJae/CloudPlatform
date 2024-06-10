package user

import (
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
	db, err := repository.OpenWithFile()

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

func (s *UserService) GetOneUserByUUID(uuid string) (*user.UserRaw, error) {
	raw, err := s.Repository.GetOneUserByUUID(uuid)

	return raw, err
}

func (s *UserService) CreateUser(n user.UserDto) (string, error) {

	// Generate Openstack Account

	// Generate Terraform Repositroy

	// Generate K8s Repository

	id, DBSaveErr := s.Repository.InsertUser(n)
	if DBSaveErr != nil {
		return "", DBSaveErr
	}
	return id, nil
}

func (s *UserService) UserIDValidate(userID string) (bool, error) {
	isExist, err := s.Repository.IsUserIDExit(userID)
	if isExist {
		return true, err
	} else {
		return false, err
	}

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

func (s *UserService) UserLogin(userID string, plainPW string) (string, error) {

	encryptedPW, err := s.Repository.GetEncryptedPW(userID)
	if err != nil {
		return "null", err
	}
	secretKey := os.Getenv("SECRET")
	if secretKey == "" {
		fmt.Println("SECRET key is not set.")
	}

	hasher := encrypt.NewPasswordHasher(secretKey)
	rslt := hasher.ComparePassword(plainPW+secretKey, encryptedPW)

	if rslt != nil {
		return "null", rslt
	}

	raw, err := s.Repository.GetOneUserByUserID(userID)

	if err != nil {
		return "error", err
	}
	return raw.Id, nil
}
