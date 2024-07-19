package apis_test

import (
	"fmt"
	apis "soms/util/apis/openstack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type KeystoneTestSuite struct {
	suite.Suite
}

type VMTestSuite struct {
	suite.Suite
}

func (kts *KeystoneTestSuite) SetupSuite() {
	apis.Init()
	apis.DeleteUser("testS")
	apis.DeleteUser("testDuplicate")
	apis.CreateUser("testDuplicate", "test", "test")
}

func (vts *VMTestSuite) SetupSuite() {
	apis.Init()
	apis.CreateUser("testVM", "test", "test")
}

func (vts *VMTestSuite) TestCreateNetowrk() {
	res, err := apis.CreateNetwork("testVM", "test", "test")
	fmt.Println(res)
	assert.NotNil(vts.T(), res)
	assert.Nil(vts.T(), err)

	apis.DeleteNetwork("testVM", "test", res)
}

func (vts *VMTestSuite) TestListNetworks() {
	res, err := apis.ListNetworks("testVM", "test")
	fmt.Println(res)
	assert.NotNil(vts.T(), res)
	assert.Nil(vts.T(), err)
}

func (vts *VMTestSuite) TestDeleteNetwork() {
	networkID, err := apis.CreateNetwork("testVM", "test", "test")
	assert.Nil(vts.T(), err)

	res, delerr := apis.DeleteNetwork("testVM", "test", networkID)

	assert.Equal(vts.T(), true, res)
	assert.Nil(vts.T(), delerr)
}

func (vts *VMTestSuite) TestListFlavors() {
	res, err := apis.ListFlavors("testVM", "test")
	fmt.Println(res)
	assert.NotNil(vts.T(), res)
	assert.Nil(vts.T(), err)
}

func (vts *VMTestSuite) TestListSecurityGroups() {
	res, err := apis.ListSecurityGroups("testVM", "test")
	fmt.Println(res)
	assert.NotNil(vts.T(), res)
	assert.Nil(vts.T(), err)

}

func (vts *VMTestSuite) TestListKeyPairs() {
	res, err := apis.ListKeyPairs("testVM", "test")
	fmt.Println(res)
	assert.NotNil(vts.T(), res)
	assert.Nil(vts.T(), err)
}

func (vts *VMTestSuite) TestListImages() {
	res, err := apis.ListImages("testVM", "test")
	fmt.Println(res)
	assert.NotNil(vts.T(), res)
	assert.Nil(vts.T(), err)
}

func (kts *KeystoneTestSuite) TestCreateDeleteUser() {
	t := kts.T()
	t.Log("TestCreateUser")
	res, err := apis.CreateUser("testS", "test", "test")
	assert.Equal(t, true, res)
	assert.Nil(t, err)

	t.Log("TestDelete")
	res, err = apis.DeleteUser("testS")
	assert.Equal(t, true, res)
	assert.Nil(t, err)

}

func (kts *KeystoneTestSuite) TestCreateDuplicateUser() {
	t := kts.T()
	t.Log("TestCreateDuplicateUser")
	res, err := apis.CreateUser("testDuplicate", "test", "test")
	assert.Equal(t, false, res)
	assert.NotNil(t, err)
}

func (kts *KeystoneTestSuite) TestGetToken() {
	t := kts.T()
	token, err := apis.GetUserToken("testDuplicate", "test")
	assert.NotNil(t, token)
	assert.Nil(t, err)

}

func (kts *KeystoneTestSuite) TestGetTokenFail() {
	t := kts.T()
	token, err := apis.GetUserToken("testS", "test")
	assert.Equal(t, "", token)
	assert.NotNil(t, err)

}
func (kts *KeystoneTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf(":::::::::::::::::::::::EndTest :: suiteName:%s :: testName:%s::::::::::::::::::::::::::::\n", suiteName, testName)
}
func (kts *KeystoneTestSuite) TearDownSuite() {
	apis.DeleteUser("testS")
	apis.DeleteUser("testDuplicate")
}
func (kts *VMTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf(":::::::::::::::::::::::EndTest :: suiteName:%s :: testName:%s::::::::::::::::::::::::::::\n", suiteName, testName)
}
func (vts *VMTestSuite) TearDownSuite() {
	apis.DeleteUser("testVM")
}
func isSkip() bool {
	return false
}

func TestKeystoneSuite(t *testing.T) {
	if isSkip() {
		t.SkipNow()
	}
	suite.Run(t, new(KeystoneTestSuite))
}

func TestVMSuite(t *testing.T) {
	if isSkip() {
		t.SkipNow()
	}
	suite.Run(t, new(VMTestSuite))
}
