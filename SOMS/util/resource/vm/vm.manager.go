package resource

import (
	"fmt"
	"os"
	"os/exec"
)

type VmManager interface {
	Init(string) VmManager
	User(string) VmManager
	Flavor(string) VmManager
	Security_groups(string) VmManager
	Keypair(string) VmManager
	Image(string) VmManager
	PrivateNetwork(string) VmManager
	ExternalIP(bool) VmManager
	Build() error
}
type vmManager struct {
	fileName       string
	userID         string
	flavorID       string
	security_group string
	imageID        string
	keypairs       string
	privateNetwork string
	externalIP     bool
}

func New() VmManager {
	return &vmManager{}
}
func (vb *vmManager) Init(fn string) VmManager {
	vb.fileName = fn
	return vb
}

func (vb *vmManager) User(u string) VmManager {
	vb.userID = u
	return vb
}

func (vb *vmManager) Flavor(f string) VmManager {
	vb.flavorID = f
	return vb
}
func (vb *vmManager) Keypair(k string) VmManager {
	vb.keypairs = k
	return vb
}
func (vb *vmManager) PrivateNetwork(s string) VmManager {
	vb.privateNetwork = s
	return vb
}
func (vb *vmManager) ExternalIP(b bool) VmManager {
	vb.externalIP = b
	return vb
}
func (vb *vmManager) Security_groups(s string) VmManager {
	vb.security_group = s
	return vb
}
func (vb *vmManager) Image(i string) VmManager {
	vb.imageID = i
	return vb
}
func (vb *vmManager) Build() error {
	terraformConfig := generateTerraformConfig(*vb)
	fileName := fmt.Sprintf("terraform/%s/%s.tf", vb.userID, vb.fileName)
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
	cmd.Dir = "terraform/" + vb.userID + "/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("terraform apply failed: %v, output: %s", err, out)
	}
	return nil
}

func generateTerraformConfig(vb vmManager) string {
	if vb.externalIP {
		return fmt.Sprintf(`resource "openstack_compute_instance_v2" "%s" {
	  name      = "%s"
	  region    = "RegionOne"
	  flavor_id = "%s"
	  key_pair  = "%s"
	  network {
		uuid = "2e26d161-5886-4e76-a9af-ad60d41761c5"
		name = "provider"
	  }

	  network {
		uuid = "%s"
	  }
	  security_groups = ["%s"]
	  image_id = "%s"
	  floating_ip = "true"
	}`, vb.fileName, vb.fileName, vb.flavorID, vb.keypairs, vb.privateNetwork, vb.security_group, vb.imageID)

	}
	return fmt.Sprintf(`resource "openstack_compute_instance_v2" "%s" {
      name      = "%s"
      region    = "RegionOne"
      flavor_id = "%s"
      key_pair  = "%s"
      network {
        uuid = "%s"
      }
      security_groups = ["%s"]
      image_id = "%s"
    }`, vb.fileName, vb.fileName, vb.flavorID, vb.keypairs, vb.privateNetwork, vb.security_group, vb.imageID)
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
