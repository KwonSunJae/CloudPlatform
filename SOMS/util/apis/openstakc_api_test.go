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
func isSkip() bool {
	return false
}

func TestKeystoneSuite(t *testing.T) {
	if isSkip() {
		t.SkipNow()
	}
	suite.Run(t, new(KeystoneTestSuite))
}
