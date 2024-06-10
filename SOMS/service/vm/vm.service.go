package vm

import (
	"fmt"
	"os"
	"os/exec"

	"soms/repository"
	"soms/repository/vm"
	resource "soms/util/resource/vm"
)

type VmService struct {
	Repository *vm.VmRepository
}

var Service VmService

func (s *VmService) InitService() error {
	db, err := repository.OpenWithFile()

	if err != nil {
		return err
	}

	s.Repository = &vm.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *VmService) GetAllVm() (*[]vm.VmRaw, error) {
	raws, err := s.Repository.GetAllVm()

	return raws, err
}

func (s *VmService) GetOneVm(id string) (*vm.VmRaw, error) {
	raw, err := s.Repository.GetOneVm(id)

	return raw, err
}

func (s *VmService) CreateVmWithTerrform(id string) error {
	n, err := s.Repository.GetOneVm(id)
	if err != nil {
		return err
	}
	// Generate Terraform configuration
	vmManager := resource.New()
	TerraformBuildErr := vmManager.
		Init(n.Name).
		User(n.UUID).
		Flavor(n.FlavorID).
		Security_groups(n.SelectedSecuritygroup).
		Keypair(n.Keypair).
		Image(n.SelectedOS).
		Build()
	if TerraformBuildErr != nil {
		return TerraformBuildErr
	}

	return nil
}
func (s *VmService) EnrollVm(n vm.VmDto) error {
	_, DBSaveErr := s.Repository.InsertVm(n)
	if DBSaveErr != nil {
		return DBSaveErr
	}
	return nil
}
func (s *VmService) GetStatusVM(userID string) (string, error) {
	// 고정된 파일 경로
	filePath := "terraform/" + userID + "/terraform.tfstate"

	// 파일 읽기
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("파일을 읽는 중 오류 발생: %v", err)
	}

	return string(fileContent), nil
}

func (s *VmService) UpdateVm(id string, n vm.VmDto) error {
	_, err := s.Repository.UpdateOneVm(id, n)

	return err
}

func (s *VmService) DeleteVm(id string) error {
	vmData, err := s.Repository.GetOneVm(id)
	if err != nil {
		return err
	}

	// Generate the filename based on the VM's name
	fileName := fmt.Sprintf("terraform/test/%s.tf", vmData.Name)
	fmt.Print(fileName)
	// Delete the Terraform file
	if err := os.Remove(fileName); err != nil {
		return err
	}
	// Run `terraform apply -auto-approve`
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "terraform/test/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, out)
	}
	_, err2 := s.Repository.DeleteOneVm(id)
	if err2 != nil {
		return err
	}
	return nil
}

func readFileContents(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	size := fileInfo.Size()

	// Read the file content
	content := make([]byte, size)
	_, err = file.Read(content)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
