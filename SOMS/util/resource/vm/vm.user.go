package resource

import (
	"fmt"
	"os"
)

func CreateUserTerrformVariableFile(userID string, userPW string) error {
	// create file
	fileContent := fmt.Sprintf(`variable "userName" {
		default = "%s"
	  }
	  
	  variable "tenantName" {
		default = "DMSLABCLOUD"
	  }
	  variable "pw" {
		default = "%s"
	  }
	`, userID, userPW)
	dir := "terraform/" + userID

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	file, err := os.Create(dir + "/variables.tf")
	if err != nil {
		return err
	}

	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil

}

func CreateMainTerrformFile(userID string) error {
	// create file
	fileContent := fmt.Sprintf(`terraform {
		required_version = ">= 1.0.0"
		  required_providers {
			openstack = {
			  source  = "terraform-provider-openstack/openstack"
			  version = "~> 1.42.0"
			}
		  }
		}
		
		# Configure the OpenStack Provider
		provider "openstack" {
		  user_name   = "${var.userName}"
		  tenant_name   = "${var.tenantName}"
		  password    = "${var.pw}"
		  auth_url    = "%s/v3/"
		  region      = "RegionOne"
		  user_domain_name = "Default"
		  endpoint_type = "public"
		  endpoint_overrides = {
			"identity" = "%s/v3/"
			"network"  = "%s/"
			"compute"  = "%s/v2.1/"
			"image"    = "%s/"
			"placement" = "%s/"
		  }
		}
	`, os.Getenv("OPENSTACK_CTRL_URL"), os.Getenv("OPENSTACK_CTRL_URL"), os.Getenv("OPENSTACK_NETWORK_URL"), os.Getenv("OPENSTACK_COMPUTE_URL"), os.Getenv("OPENSTACK_IMAGE_URL"), os.Getenv("OPENSTACK_PLACEMENT_URL"))

	dir := "terraform/" + userID
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	file, err := os.Create(dir + "/main.tf")
	if err != nil {
		return err
	}

	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	return nil
}
