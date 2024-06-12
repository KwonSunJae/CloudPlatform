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
	// rslt, err := openstack_api.CreateUser("realtest", "testpw", "test@email.com")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(rslt)

	// openstack_api.CreateNetwork("test", "testpw", "test-network", os.Getenv("OPENSTACK_NETWORK_ENDPOINT"))
	// openstack_api.ListNetworks("test", "testpw", os.Getenv("OPENSTACK_NETWORK_ENDPOINT"))
	// openstack_api.CreateKeyPair("test", "testpw", os.Getenv("OPENSTACK_COMPUTE_ENDPOINT"), "test-keypair")
	resp, err := openstack_api.ListKeyPairs((*testUser).UserID, (*testUser).EncryptedPW)
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
	}
	resp2, err := openstack_api.ListFlavors(testUser.UserID, testUser.EncryptedPW)
	fmt.Println(resp2)

	resp3, err := openstack_api.ListSecurityGroups(testUser.UserID, testUser.EncryptedPW)
	fmt.Println(resp3)

}
