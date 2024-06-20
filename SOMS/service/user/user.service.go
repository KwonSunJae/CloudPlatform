package user

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"soms/repository"
	"soms/repository/user"
	openstack_api "soms/util/apis/openstack"
	"soms/util/encrypt"
	resource "soms/util/resource/vm"
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
	id, DBSaveErr := s.Repository.InsertUser(n)
	if DBSaveErr != nil {
		return "", DBSaveErr
	}
	return id, nil
}
func (s *UserService) ApproveUser(id string, role string, priority string) error {
	if priority == "Denied" {
		s.Repository.DeleteOneUser(id)
		return nil
	}
	var approvedUser *user.UserRaw
	approvedUser, err := s.Repository.GetOneUserByUUID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Generate Openstack Account
	res, err := openstack_api.CreateUser(approvedUser.UserID, approvedUser.EncryptedPW, approvedUser.Name+"@"+approvedUser.Spot+"."+priority)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !res {

		return fmt.Errorf("openstack account creation failed")
	}
	// Generate Terraform Repositroy & Create main.tf, variables.tf
	dirName := approvedUser.UserID
	err = os.Mkdir("terraform/"+dirName, 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}
	creationVarTFErr := resource.CreateUserTerrformVariableFile(approvedUser.UserID, approvedUser.EncryptedPW)
	if creationVarTFErr != nil {
		fmt.Println(creationVarTFErr)
		return errors.New("Failed to create terraform variable file :" + err.Error())
	}
	creationMainTFErr := resource.CreateMainTerrformFile(approvedUser.UserID)
	if creationMainTFErr != nil {
		fmt.Println(creationMainTFErr)
		return errors.New("Failed to create terraform main file :" + err.Error())
	}
	initTFerr := resource.InitTerraform(approvedUser.UserID)
	if initTFerr != nil {
		fmt.Println(initTFerr)
		return errors.New("Failed to init terraform :" + err.Error())
	}
	// Generate K8s Repository
	err = os.Mkdir("k8s/"+dirName, 0755)
	if err != nil {
		return err
	}

	// kubectl create namespace -n $userID
	cmd := exec.Command("kubectl", "create", "namespace", approvedUser.UserID)
	_, err2 := cmd.CombinedOutput()
	if err2 != nil {
		fmt.Println(err2)
		return fmt.Errorf("Failed to create namespace: %v", err2)
	}

	// Update User Role and Priority
	var n user.UserDto
	n.Role = role
	n.Priority = priority

	rest, err := s.Repository.UpdateOneUser(id, n)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if rest == nil {
		return fmt.Errorf("NOT FOUND")
	}

	return nil
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
