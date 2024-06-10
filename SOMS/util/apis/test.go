package main

import (
	"fmt"
	openstack_api "soms/util/apis/openstack"

	"github.com/joho/godotenv"
)

func main() {

	errw := godotenv.Load()
	if errw != nil {
		panic("env file error")
	}
	openstack_api.CreateUser("test2", "testpw", "test@email.com")
	res, err := openstack_api.GetUserToken("test2", "testpw")
	fmt.Println("UserToken is " + res)
	if err != nil {
		panic(err)
	}
	openstack_api.CreateNetwork()
	openstack_api.ListNetworks()
}
