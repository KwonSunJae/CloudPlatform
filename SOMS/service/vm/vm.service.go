package vm

import (
	"fmt"
	"os"
	"os/exec"

	"soms/repository"

	"soms/repository/vm"

	user "soms/repository/user"
	openstack_api "soms/util/apis/openstack"
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

func (s *VmService) ApproveVMCreation(id string, uuid string) error {
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return err
	}
	n, err := s.Repository.GetOneVm(id)

	if err != nil {
		return err
	}

	var floating_ip bool
	if n.ExternalIP == "true" {
		floating_ip = true
	} else {
		floating_ip = false
	}
	// Generate Terraform configuration & Execute `terraform apply -auto-approve`
	vmManager := resource.New()
	TerraformBuildErr := vmManager.
		Init(n.Name).
		User(targetUser.UserID).
		Flavor(n.FlavorID).
		Security_groups(n.SelectedSecuritygroup).
		PrivateNetwork(n.InternalIP).
		ExternalIP(floating_ip).
		Keypair(n.Keypair).
		Image(n.SelectedOS).
		Build()
	if TerraformBuildErr != nil {
		return TerraformBuildErr
	}

	// Update the VM status
	n.Status = "Approved"
	var an vm.VmDto
	an.Status = "Approved"
	_, err = s.Repository.UpdateOneVm(id, an)
	if err != nil {
		return err
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
func (s *VmService) GetStatusVM(userID string) (string, error) {
	// 고정된 파일 경로
	filePath := "terraform/" + userID + "/terraform.tfstate"

	// 파일 읽기
	fileContent, err := readFileContents(filePath)
	if err != nil {
		return "", fmt.Errorf("파일을 읽는 중 오류 발생: %v", err)
	}

	return string(fileContent), nil
}

func (s *VmService) UpdateVm(id string, n vm.VmDto) error {
	_, err := s.Repository.UpdateOneVm(id, n)

	return err
}

func (s *VmService) DeleteVm(id string, uuid string) error {
	// Get the User data
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return err
	}

	// Get the VM data
	vmData, err := s.Repository.GetOneVm(id)
	if err != nil {
		return err
	}

	// Generate the filename based on the VM's name
	fileName := fmt.Sprintf("terraform/%s/%s.tf", targetUser.UserID, vmData.Name)
	fmt.Print(fileName)
	// Delete the Terraform file
	if err := os.Remove(fileName); err != nil {
		return err
	}
	// Run `terraform apply -auto-approve`
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "terraform/" + targetUser.UserID + "/"
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

func (s *VmService) CreateNetwork(uuid string, networkName string) (string, error) {
	//Openstack API call to create network
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.CreateNetwork(targetUser.UserID, targetUser.EncryptedPW, networkName)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *VmService) ListNetworks(uuid string) (string, error) {
	//Openstack API call to list networks
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.ListNetworks(targetUser.UserID, targetUser.EncryptedPW)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *VmService) ListFlavors(uuid string) (string, error) {
	//Openstack API call to list flavors
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.ListFlavors(targetUser.UserID, targetUser.EncryptedPW)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *VmService) ListKeypairs(uuid string) (string, error) {
	//Openstack API call to list keypairs
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.ListKeyPairs(targetUser.UserID, targetUser.EncryptedPW)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *VmService) CreateKeypair(uuid string, keypairName string) (string, error) {
	//Openstack API call to create keypair
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.CreateKeyPair(targetUser.UserID, targetUser.EncryptedPW, keypairName)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *VmService) ListSecurityGroups(uuid string) (string, error) {
	//Openstack API call to list security groups

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.ListSecurityGroups(targetUser.UserID, targetUser.EncryptedPW)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *VmService) CreateSnapshot(uuid string, vmID string, snapshotName string) (bool, error) {
	//Openstack API call to create snapshot

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return false, err
	}

	result, err := openstack_api.CreateSnapshot(targetUser.UserID, targetUser.EncryptedPW, vmID, snapshotName)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *VmService) SoftReboot(uuid string, vmID string) (bool, error) {
	//Openstack API call to soft reboot

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return false, err
	}

	result, err := openstack_api.SoftReboot(targetUser.UserID, targetUser.EncryptedPW, vmID)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *VmService) HardReboot(uuid string, vmID string) (bool, error) {
	//Openstack API call to hard reboot

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return false, err
	}

	result, err := openstack_api.HardReboot(targetUser.UserID, targetUser.EncryptedPW, vmID)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (s *VmService) PowerOff(uuid string, vmID string) (bool, error) {
	//Openstack API call to power off

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return false, err
	}

	result, err := openstack_api.PowerOff(targetUser.UserID, targetUser.EncryptedPW, vmID)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *VmService) PowerOn(uuid string, vmID string) (bool, error) {
	//Openstack API call to power on

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return false, err
	}

	result, err := openstack_api.PowerOn(targetUser.UserID, targetUser.EncryptedPW, vmID)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *VmService) GetVnc(uuid string, vmID string) (string, error) {
	//Openstack API call to get VNC console

	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", err
	}

	result, err := openstack_api.GetVNCConsoleURL(targetUser.UserID, targetUser.EncryptedPW, vmID)
	if err != nil {
		return "", err
	}

	return result, nil
}
