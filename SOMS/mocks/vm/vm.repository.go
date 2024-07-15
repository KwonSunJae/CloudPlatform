package mocks

import (
	"github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type VmRepository struct {
	mock.Mock
}

type VmRaw struct {
	Id                    string
	Name                  string
	FlavorID              string
	ExternalIP            string
	InternalIP            string
	SelectedOS            string
	UnionmountImage       string
	Keypair               string
	SelectedSecuritygroup string
	UUID                  string
	Status                string
}

// GetAllVm provides a mock function with given fields:
func (_m *VmRepository) GetAllVm() (*[]VmRaw, error) {
	ret := _m.Called()

	var r0 *[]VmRaw
	if rf, ok := ret.Get(0).(func() *[]VmRaw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]VmRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneVmByUUID provides a mock function with given fields: uuid
func (_m *VmRepository) GetOneVmByUUID(uuid string) (*VmRaw, error) {
	ret := _m.Called(uuid)

	var r0 *VmRaw
	if rf, ok := ret.Get(0).(func(string) *VmRaw); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VmRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneVm provides a mock function with given fields: id
func (_m *VmRepository) GetOneVm(id string) (*VmRaw, error) {
	ret := _m.Called(id)

	var r0 *VmRaw
	if rf, ok := ret.Get(0).(func(string) *VmRaw); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VmRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOneVm provides a mock function with given fields: id
func (_m *VmRepository) DeleteOneVm(id string) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVm provides a mock function with given fields: name, flavorID, selectedOS, unionmountImage, keypair, selectedSecuritygroup
func (_m *VmRepository) CreateVm(name string, flavorID string, selectedOS string, unionmountImage string, keypair string, selectedSecuritygroup string) (int64, error) {
	ret := _m.Called(name, flavorID, selectedOS, unionmountImage, keypair, selectedSecuritygroup)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string) int64); ok {
		r0 = rf(name, flavorID, selectedOS, unionmountImage, keypair, selectedSecuritygroup)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string, string) error); ok {
		r1 = rf(name, flavorID, selectedOS, unionmountImage, keypair, selectedSecuritygroup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOneVm provides a mock function with given fields: id, n
func (_m *VmRepository) UpdateOneVm(id string, n VmRaw) (int64, error) {
	ret := _m.Called(id, n)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, VmRaw) int64); ok {
		r0 = rf(id, n)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, VmRaw) error); ok {
		r1 = rf(id, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}