package user_test

import (
	"errors"
	"fmt"
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
		fmt.Println(raws)
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

	t.Run("TestGetOneUserByUUIDSuccess", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := user.NewUserService(&user.USConfig{Repository: mockUserRepository})

		mockUserRepository.On("GetOneUserByUUID", "test").Return(&model.UserRaw{
			Id:          "test",
			Name:        "test",
			UserID:      "test",
			EncryptedPW: "test",
			Role:        "test",
			Spot:        "test",
			Priority:    "test",
		}, nil)

		raw, err := userService.GetOneUserByUUID("test")
		fmt.Println(raw)
		assert.NotNil(t, raw)
		assert.Nil(t, err)
	})

	t.Run("TestGetOneUserByUUIDError", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		mockUserRepository.On("GetOneUserByUUID", "test").Return(nil, errors.New("error"))
		userService := user.UserService{Repository: mockUserRepository}
		raw, err := userService.GetOneUserByUUID("test")
		assert.Nil(t, raw)
		assert.NotNil(t, err)
	})

}

func TestCreateUser(t *testing.T) {
	t.Log("TestCreateUser")
	//logic 수정예정
}

func TestApproveUser(t *testing.T) {
	t.Log("TestApproveUser")
	//logic 수정예정
}

func TestDeleteOneUser(t *testing.T) {
	t.Log("TestDeleteOneUser")
	//logic 수정예정
}

func TestUpdateUser(t *testing.T) {
	t.Log("TestUpdateUser")
	//logic 수정예정
}

func TestIsUserIDValidate(t *testing.T) {
	t.Log("TestIsUserIDExit")
	t.Run("TestUserIDValidatetSuccess", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := user.NewUserService(&user.USConfig{Repository: mockUserRepository})

		mockUserRepository.On("IsUserIDExit", "test").Return(true, nil)

		isExit, err := userService.UserIDValidate("test")
		assert.Equal(t, true, isExit)
		assert.Nil(t, err)
	})

	t.Run("TestUserIDValidatetError", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		mockUserRepository.On("IsUserIDExit", "test").Return(false, errors.New("error"))
		userService := user.UserService{Repository: mockUserRepository}
		isExit, err := userService.UserIDValidate("test")
		assert.Equal(t, false, isExit)
		assert.NotNil(t, err)
	})

}

func TestUserLogin(t *testing.T) {
	t.Log("TestUserLogin")
	t.Run("TestUserLoginSuccess", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := user.NewUserService(&user.USConfig{Repository: mockUserRepository})

		mockUserRepository.On("GetEncryptedPW", "test").Return("test", nil)
		isLogin, err := userService.UserLogin("test", "test")
		assert.Equal(t, true, isLogin)
		assert.Nil(t, err)
	})
}
