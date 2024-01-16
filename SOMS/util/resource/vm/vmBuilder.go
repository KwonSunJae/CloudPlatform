package resource

import (
	"fmt"
	"os"
	"os/exec"
)

type VmBuilder interface {
	Init(string) VmBuilder
	User(string) VmBuilder
	Flavor(string) VmBuilder
	Security_groups(string) VmBuilder
	Keypair(string) VmBuilder
	Image(string) VmBuilder
	Build() error
}
type vmBuilder struct {
	fileName       string
	userID         string
	flavorID       string
	security_group string
	imageID        string
	keypairs       string
}

func New() VmBuilder {
	return &vmBuilder{}
}
func (vb *vmBuilder) Init(fn string) VmBuilder {
	vb.fileName = fn
	return vb
}

func (vb *vmBuilder) User(u string) VmBuilder {
	vb.userID = u
	return vb
}

func (vb *vmBuilder) Flavor(f string) VmBuilder {
	vb.flavorID = f
	return vb
}
func (vb *vmBuilder) Keypair(k string) VmBuilder {
	vb.keypairs = k
	return vb
}
func (vb *vmBuilder) Security_groups(s string) VmBuilder {
	vb.security_group = s
	return vb
}
func (vb *vmBuilder) Image(i string) VmBuilder {
	vb.imageID = i
	return vb
}
func (vb *vmBuilder) Build() error {
	terraformConfig := generateTerraformConfig(*vb)
	fileName := fmt.Sprintf("terraform/test/%s.tf", vb.fileName)
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
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Dir = "terraform/test/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, out)
	}
	return nil
}

func generateTerraformConfig(vb vmBuilder) string {
	return fmt.Sprintf(`resource "openstack_compute_instance_v2" "%s" {
      name      = "%s"
      region    = "RegionOne"
      flavor_id = "%s"
      key_pair  = "%s"
      network {
        uuid = "2e26d161-5886-4e76-a9af-ad60d41761c5"
        name = "provider"
      }
      security_groups = ["%s"]
      image_id = "%s"
    }`, vb.fileName, vb.fileName, vb.flavorID, vb.keypairs, vb.security_group, vb.imageID)
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
