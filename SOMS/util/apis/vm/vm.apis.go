// Openstack Compute API Interface
// This file is generated by code generator. DO NOT EDIT!

package vm

type OpenstackComputeAPI interface {
	// Get a list of servers
	// GET /v2/{tenant_id}/servers
	GetServers() (string, error)
}

func NewOpenstackComputeAPI() OpenstackComputeAPI {
	return &openstackComputeAPI{}
}

type openstackComputeAPI struct {
}

func (o *openstackComputeAPI) GetServers() (string, error) {
