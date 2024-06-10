package main

import (
	"fmt"
	"os"
	openstack_api "soms/util/apis/openstack"

	"github.com/joho/godotenv"
)

func main() {

	errw := godotenv.Load()
	if errw != nil {
		panic("env file error")
	}
	// openstack_api.CreateUser("test2", "testpw", "test@email.com")
	res, err := openstack_api.GetUserToken("test", "testpw")
	fmt.Println("UserToken is " + res)
	if err != nil {
		panic(err)
	}
	openstack_api.CreateNetwork("test", "testpw", "test-network", os.Getenv("OPENSTACK_NETWORK_ENDPOINT"))
	openstack_api.ListNetworks("test", "testpw", os.Getenv("OPENSTACK_NETWORK_ENDPOINT"))
}
