package user_test

import (
	"errors"
	mocks "soms/mocks/user"
	model "soms/model/user"
	"soms/service/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {

	t.Log("TestGetAllUser")
	t.Run("TestGetAllUserSuccess", func(t *testing.T) {

		mockUserRepository := new(mocks.UserRepository)
		userService := user.NewUserService(&user.USConfig{Repository: mockUserRepository})

		mockUserRepository.On("GetAllUser").Return(&[]model.UserRaw{
			{
				Id:          "test",
				Name:        "test",
				UserID:      "test",
				EncryptedPW: "test",
				Role:        "test",
				Spot:        "test",
				Priority:    "test",
			},
		}, nil)

		raws, err := userService.GetAllUser()
		assert.NotNil(t, raws)
		assert.Nil(t, err)
	})
	t.Run("TestGetAllUserError", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		mockUserRepository.On("GetAllUser").Return(nil, errors.New("error"))
		userService := user.UserService{Repository: mockUserRepository}
		raws, err := userService.GetAllUser()
		assert.Nil(t, raws)
		assert.NotNil(t, err)
	})
}

func TestGetOneUserByUUID(t *testing.T) {
	t.Log("TestGetOneUserByUUID")
}

func TestCreateUser(t *testing.T) {
	t.Log("TestCreateUser")
}

func TestApproveUser(t *testing.T) {
	t.Log("TestApproveUser")
}

func TestDeleteOneUser(t *testing.T) {
	t.Log("TestDeleteOneUser")
}

func TestUpdateUser(t *testing.T) {
	t.Log("TestUpdateUser")
}
