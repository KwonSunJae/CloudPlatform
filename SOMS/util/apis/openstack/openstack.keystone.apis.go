package openstack_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// .env 파일에서 가져온 환경변수를 사용하여 OpenStack 인증 URL, 프로젝트 ID, 도메인 ID, 관리자 사용자 이름, 관리자 비밀번호를 설정
var (
	authURL   = os.Getenv("OPENSTACK_CTRL_URL") + "/v3/auth/tokens" // OpenStack 인증 URL
	projectID = os.Getenv("OPENSTACK_PROJECT_ID")
	domainID  = os.Getenv("OPENSTACK_DOMAIN_ID")      // 기본 도메인 ID
	adminUser = os.Getenv("OPENSTACK_ADMIN_USERNAME") // 관리자 사용자 이름
	adminPass = os.Getenv("OPENSTACK_ADMIN_PW")       // 관리자 비밀번호
	initFlag  = false
)

type AuthRequest struct {
	Auth struct {
		Identity struct {
			Methods  []string `json:"methods"`
			Password struct {
				User struct {
					Name   string `json:"name"`
					Domain struct {
						ID string `json:"id"`
					} `json:"domain"`
					Password string `json:"password"`
				} `json:"user"`
			} `json:"password"`
		} `json:"identity"`
		Scope struct {
			Project struct {
				ID string `json:"id"`
			} `json:"project"`
		} `json:"scope"`
	} `json:"auth"`
}

type UserRequest struct {
	User struct {
		Name     string `json:"name"`
		DomainID string `json:"domain_id"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Enabled  bool   `json:"enabled"`
	} `json:"user"`
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("env file error")
	}
	authURL = os.Getenv("OPENSTACK_CTLR_URL") + "/v3/auth/tokens" // OpenStack 인증 URL
	projectID = os.Getenv("OPENSTACK_PROJECT_ID")
	domainID = os.Getenv("OPENSTACK_DOMAIN_ID")       // 기본 도메인 ID
	adminUser = os.Getenv("OPENSTACK_ADMIN_USERNAME") // 관리자 사용자 이름
	adminPass = os.Getenv("OPENSTACK_ADMIN_PW")       // 관리자 비밀번호
	initFlag = true
}

func getAdminAuthToken() (string, error) {
	if !initFlag {
		Init()
	}
	authReq := AuthRequest{}
	authReq.Auth.Identity.Methods = []string{"password"}
	authReq.Auth.Identity.Password.User.Name = adminUser
	authReq.Auth.Identity.Password.User.Domain.ID = domainID
	authReq.Auth.Identity.Password.User.Password = adminPass
	authReq.Auth.Scope.Project.ID = projectID

	requestBody, err := json.Marshal(authReq)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to get token, status code: %d", resp.StatusCode)
	}

	token := resp.Header.Get("X-Subject-Token")
	return token, nil
}

func GetUserToken(username string, password string) (string, error) {
	if !initFlag {
		Init()
	}
	authReq := AuthRequest{}
	authReq.Auth.Identity.Methods = []string{"password"}
	authReq.Auth.Identity.Password.User.Name = username
	authReq.Auth.Identity.Password.User.Domain.ID = domainID
	authReq.Auth.Identity.Password.User.Password = password
	authReq.Auth.Scope.Project.ID = os.Getenv("OPENSTACK_COMMON_PROJECT_ID")

	requestBody, err := json.Marshal(authReq)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to get token, status code: %d", resp.StatusCode)
	}

	token := resp.Header.Get("X-Subject-Token")
	return token, nil
}

func createUser(authToken, name, password, email string) error {
	userReq := UserRequest{}
	userReq.User.Name = name
	userReq.User.DomainID = domainID
	userReq.User.Password = password
	userReq.User.Email = email
	userReq.User.Enabled = true

	requestBody, err := json.Marshal(userReq)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("OPENSTACK_CTLR_URL")+"/v3/users", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to create user, status code: %d, response: %s", resp.StatusCode, string(body))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var userMap map[string]interface{}
	json.Unmarshal(body, &userMap)

	newUserID := userMap["user"].(map[string]interface{})["id"].(string)
	fmt.Println("new user ID: ", newUserID)

	err2 := AddUserToProject(newUserID)
	if err2 != nil {
		return err
	}
	return nil
}

func CreateUser(username string, password string, userEmail string) (bool, error) {
	Init()
	authToken, err := getAdminAuthToken()
	if err != nil {
		return false, errors.New("get Admin AuthTOKEN ERROR : " + err.Error())
	}

	err = createUser(authToken, username, password, userEmail)
	if err != nil { // 에러 발생 시
		return false, errors.New("create user ERROR: " + err.Error())
	}

	return true, nil
}

func addUserToProject(authToken, endpoint, projectID, userID, roleID string) error {
	url := fmt.Sprintf("%s/v3/projects/%s/users/%s/roles/%s", endpoint, projectID, userID, roleID)

	req, _ := http.NewRequest("PUT", url, nil)
	req.Header.Set("X-Auth-Token", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("failed to add user to project: " + err.Error())
	}
	defer resp.Body.Close()

	return nil
}

func AddUserToProject(userID string) error {
	authToken, err := getAdminAuthToken()
	if err != nil {
		return errors.New("get Admin AuthTOKEN ERROR : " + err.Error())
	}
	var roleID = os.Getenv("OPENSTACK_COMMON_ROLE_ID")
	var commonProjectID = os.Getenv("OPENSTACK_COMMON_PROJECT_ID")
	var endpoint = os.Getenv("OPENSTACK_CTLR_URL")

	addUserToProject(authToken, endpoint, commonProjectID, userID, roleID)
	return nil
}

// Delete User
func deleteUserFromProject(authToken, endpoint, userID, roleID, projectID string) (bool, error) {
	url := fmt.Sprintf("%s/v3/users/%s", endpoint, userID)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("X-Auth-Token", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("failed to delete user: " + err.Error())
	}

	res := http.StatusNoContent == resp.StatusCode
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	return res, nil
}
func getUserID(authToken, endpoint, username string) (string, error) {
	url := fmt.Sprintf("%s/v3/users?name=%s", endpoint, username)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to get user ID: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var userMap map[string]interface{}
	json.Unmarshal(body, &userMap)

	userList := userMap["users"].([]interface{})
	if len(userList) == 0 {
		return "", errors.New("user not found")
	}
	userID := userList[0].(map[string]interface{})["id"].(string)

	return userID, nil
}
func GetUserID(username string) (string, error) {
	authToken, err := getAdminAuthToken()
	if err != nil {
		return "", errors.New("get Admin AuthTOKEN ERROR : " + err.Error())
	}
	var endpoint = os.Getenv("OPENSTACK_CTLR_URL")

	return getUserID(authToken, endpoint, username)
}

func DeleteUser(userName string) (bool, error) {
	authToken, err := getAdminAuthToken()
	if err != nil {
		return false, errors.New("get Admin AuthTOKEN ERROR : " + err.Error())
	}
	var userID, err404 = GetUserID(userName)
	if err404 != nil {
		return false, errors.New("get User ID ERROR : " + err404.Error())
	}
	var endpoint = os.Getenv("OPENSTACK_CTLR_URL")
	var roleID = os.Getenv("OPENSTACK_COMMON_ROLE_ID")
	var commonProjectID = os.Getenv("OPENSTACK_COMMON_PROJECT_ID")

	return deleteUserFromProject(authToken, endpoint, userID, roleID, commonProjectID)
}
