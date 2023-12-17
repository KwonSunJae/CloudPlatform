package vm

import (
	"fmt"
	"os"
	"os/exec"

	"soms/repository"
	"soms/repository/vm"
)

type VmService struct {
	Repository *vm.VmRepository
}

var Service VmService

func (s *VmService) InitService() error {
	db, err := repository.OpenWithMemory()

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

func (s *VmService) CreateVm(n vm.VmDto) error {
	_, err := s.Repository.InsertVm(n)
	if err != nil {
		return err
	}

	// Generate Terraform configuration
	terraformConfig := generateTerraformConfig(n)
	fileName := fmt.Sprintf("terraform/test/%s.tf", n.Name)
	// Write to vm.tf
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(terraformConfig)
	if err != nil {
		return err
	}
	fmt.Print(fileName)

	// Run `terraform apply -auto-approve` using an appropriate command execution method
	// ...
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "terraform/test/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, out)
	}

	return nil
}
func (s *VmService) GetStatusVM() (string, error) {
	// 고정된 파일 경로
	filePath := "terraform/test/terraform.tfstate"

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
	_, err2 := s.Repository.DeleteOneVm(id)
	if err2 != nil {
		return err
	}

	// Generate the filename based on the VM's name
	fileName := fmt.Sprintf("terraform/test/%s.tf", vmData.Name)

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
	return nil
}

func generateTerraformConfig(vmDto vm.VmDto) string {
	return fmt.Sprintf(`resource "openstack_compute_instance_v2" "%s" {
      name      = "%s"
      region    = "RegionOne"
      flavor_id = "%s"
      key_pair  = "%s"
      network {
        uuid = "2e26d161-5886-4e76-a9af-ad60d41761c5"
        name = "provider"
      }
      security_groups = ["default"]
      image_id = "%s"
    }`, vmDto.Name, vmDto.Name, vmDto.FlavorID, vmDto.Keypair, vmDto.SelectedOS)
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
