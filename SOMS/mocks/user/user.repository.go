package mocks

import (
	"github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

type UserRaw struct {
	Id          string
	Name        string
	UserID      string
	EncryptedPW string
	Role        string
	Spot        string
	Priority    string
}

// GetAllUser provides a mock function with given fields:
func (_m *UserRepository) GetAllUser() (*[]UserRaw, error) {
	ret := _m.Called()

	var r0 *[]UserRaw
	if rf, ok := ret.Get(0).(func() *[]UserRaw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]UserRaw)
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

// GetOneUserByUUID provides a mock function with given fields: uuid
func (_m *UserRepository) GetOneUserByUUID(uuid string) (*UserRaw, error) {
	ret := _m.Called(uuid)

	var r0 *UserRaw
	if rf, ok := ret.Get(0).(func(string) *UserRaw); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserRaw)
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

// GetOneUserByUserID provides a mock function with given fields: userID
func (_m *UserRepository) GetOneUserByUserID(userID string) (*UserRaw, error) {
	ret := _m.Called(userID)

	var r0 *UserRaw
	if rf, ok := ret.Get(0).(func(string) *UserRaw); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsUserIDExit provides a mock function with given fields: userID
func (_m *UserRepository) IsUserIDExit(userID string) (bool, error) {
	ret := _m.Called(userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: name, userID, encryptedPW, role, spot, priority
func (_m *UserRepository) CreateUser(name string, userID string, encryptedPW string, role string, spot string, priority string) (*UserRaw, error) {
	ret := _m.Called(name, userID, encryptedPW, role, spot, priority)

	var r0 *UserRaw
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) *UserRaw); ok {
		r0 = rf(name, userID, encryptedPW, role, spot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(name, userID, encryptedPW, role, spot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApproveUser provides a mock function with given fields: id
func (_m *UserRepository) ApproveUser(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOneUser provides a mock function with given fields: id
func (_m *UserRepository) DeleteOneUser(id string) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOneUser provides a mock function with given fields: id, n
func (_m *UserRepository) UpdateOneUser(id string, n UserRaw) (bool, error) {
	ret := _m.Called(id, n)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, UserRaw) bool); ok {
		r0 = rf(id, n)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, UserRaw) error); ok {
		r1 = rf(id, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
