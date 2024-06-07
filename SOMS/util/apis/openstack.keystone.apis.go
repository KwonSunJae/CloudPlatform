package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// .env 파일에서 가져온 환경변수를 사용하여 OpenStack 인증 URL, 프로젝트 ID, 도메인 ID, 관리자 사용자 이름, 관리자 비밀번호를 설정
var (
	authURL   = os.Getenv("OPENSTACK_CTRL_URL") + "/v3/auth/tokens" // OpenStack 인증 URL
	projectID = os.Getenv("OPENSTACK_PROJECT_ID")
	domainID  = os.Getenv("OPENSTACK_DOMAIN_ID")      // 기본 도메인 ID
	adminUser = os.Getenv("OPENSTACK_ADMIN_USERNAME") // 관리자 사용자 이름
	adminPass = os.Getenv("OPENSTACK_ADMIN_PW")       // 관리자 비밀번호
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

func getAdminAuthToken() (string, error) {
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
	authReq := AuthRequest{}
	authReq.Auth.Identity.Methods = []string{"password"}
	authReq.Auth.Identity.Password.User.Name = username
	authReq.Auth.Identity.Password.User.Domain.ID = domainID
	authReq.Auth.Identity.Password.User.Password = password
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
	new
	return nil
}

func CreateUser(username string, password string, userEmail string) (bool, error) {

	authURL = os.Getenv("OPENSTACK_CTLR_URL") + "/v3/auth/tokens" // OpenStack 인증 URL
	projectID = os.Getenv("OPENSTACK_PROJECT_ID")
	domainID = os.Getenv("OPENSTACK_DOMAIN_ID")       // 기본 도메인 ID
	adminUser = os.Getenv("OPENSTACK_ADMIN_USERNAME") // 관리자 사용자 이름
	adminPass = os.Getenv("OPENSTACK_ADMIN_PW")       // 관리자 비밀번호

	authToken, err := getAdminAuthToken()
	if err != nil {
		return false, errors.New("get Admin AuthTOKEN ERROR : " + err.Error())
	}

	err = createUser(authToken, username, password, userEmail)
	if err != nil { // 에러 발생 시
		return false, errors.New("creat user ERROR: " + err.Error())
	}

	return true, nil
}
