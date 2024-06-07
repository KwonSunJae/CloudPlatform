package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	reqchecker "soms/controller/checker"
	response "soms/controller/response"
	"soms/service/user"
	"soms/util/encrypt"

	"github.com/gorilla/mux"
)

func UserController(router *mux.Router) error {
	err := user.Service.InitService()

	if err != nil {
		return err
	}

	router.HandleFunc("/user/{userID}", getUserByUuid).Methods("GET")

	router.HandleFunc("/user", getAllUser).Methods("GET")

	router.HandleFunc("/user/validate/{userID}", userIDValidate).Methods("GET")

	router.HandleFunc("/user", userRegister).Methods("POST")

	router.HandleFunc("/user/login", userLogin).Methods("POST")

	router.HandleFunc("/user/{userID}", updateUser).Methods("PATCH")

	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	return nil
}

// @Summary 사용자 정보 조회
// @Description 사용자의 정보를 조회합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   userID     path    string     true  "User ID"
// @Success 200 {object} response.CommonResponse
// @Router /user/{userID} [get]
func getUserByUuid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	raw, err := user.Service.GetOneUser(userID)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 User가 없습니다"))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, raw, http.StatusOK, nil)

}

// @Summary 사용자 정보 전체 조회
// @Description 사용자의 정보를 전체 조회합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse
// @Router /user [get]
func getAllUser(w http.ResponseWriter, r *http.Request) {
	raws, err := user.Service.GetAllUser()

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, raws, http.StatusOK, nil)

}

// @Summary 사용자 ID 유효성 검사
// @Description 사용자 ID의 유효성을 검사합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   userID	path   string     true  "User ID"
// @Success 200 {object} response.CommonResponse
// @Router /user/validate/{userID} [get]
func userIDValidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	isExist, err := user.Service.UserIDValidate(userID)

	if !isExist {
		response.Response(w, isExist, http.StatusConflict, errors.New("해당 User가 이미 존재합니다"))
		return
	} else {
		response.Response(w, isExist, http.StatusOK, err)
		return
	}

}

type UserRequestBody struct {
	Name     string
	UserID   string
	PW       string
	Role     string
	Spot     string
	Priority string
}

// @Summary 사용자 가입
// @Description 사용자 가입을 진행합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   User	 body   UserRequestBody     true  "User Name"
// @Success 200 {object} response.CommonResponse
// @Router /user [post]
func userRegister(w http.ResponseWriter, r *http.Request) {
	var body UserRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusBadRequest, err)
		return
	}

	paramsErr := reqchecker.Check(body)
	if paramsErr != nil {
		response.Response(w, nil, http.StatusBadRequest, paramsErr)
		return
	}
	secretKey := os.Getenv("SECRET")
	if secretKey == "" {
		fmt.Println("SECRET key is not set.")
	}
	hasher := encrypt.NewPasswordHasher(secretKey)
	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	encryptedPW, HasherError := hasher.HashPassword(body.PW)
	if HasherError != nil {
		response.Response(w, nil, http.StatusInternalServerError, HasherError)
		return
	}

	dto := struct {
		Name        string
		UserID      string
		EncryptedPW string
		Role        string
		Spot        string
		Priority    string
	}{
		Name:        body.Name,
		UserID:      body.UserID,
		EncryptedPW: encryptedPW,
		Role:        body.Role,
		Spot:        body.Spot,
		Priority:    body.Priority,
	}

	id, err := user.Service.CreateUser(dto)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, id, http.StatusOK, nil)

}

type UserLoginRequestBody struct {
	UserID string
	PW     string
}

// @Summary 사용자 로그인
// @Description 사용자 로그인을 진행합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   UserLogin	 body    UserLoginRequestBody     true  "User Login Info"
// @Success 200 {object} response.CommonResponse
// @Router /user/login [post]
func userLogin(w http.ResponseWriter, r *http.Request) {
	var body UserLoginRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	paramsErr := reqchecker.Check(body)
	if paramsErr != nil {
		response.Response(w, nil, http.StatusBadRequest, paramsErr)
		return
	}
	if err != nil {
		response.Response(w, nil, http.StatusBadRequest, err)
	}

	rslt, err := user.Service.UserLogin(body.UserID, body.PW)

	if err != nil {
		response.Response(w, nil, http.StatusBadRequest, err)
		return
	}

	response.Response(w, rslt, http.StatusOK, nil)

}

// @Summary 사용자 정보 수정
// @Description 사용자의 정보를 수정합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   UserID	 path    string     true  "User ID"
// @Param   User    body    UserRequestBody     true  "User Name"
// @Success 200 {object} response.CommonResponse
// @Router /user/{userID} [patch]
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	var body UserRequestBody
	decodeErr := json.NewDecoder(r.Body).Decode(&body)

	if decodeErr != nil {
		response.Response(w, nil, http.StatusBadRequest, decodeErr)
		return
	}
	secretKey := os.Getenv("SECRET")
	if secretKey == "" {
		fmt.Println("SECRET key is not set.")
	}
	hasher := encrypt.NewPasswordHasher(secretKey)

	encryptedPW, HasherError := hasher.HashPassword(body.PW)
	if HasherError != nil {
		response.Response(w, nil, http.StatusInternalServerError, HasherError)
	}

	dto := struct {
		Name        string
		UserID      string
		EncryptedPW string
		Role        string
		Spot        string
		Priority    string
	}{
		Name:        body.Name,
		UserID:      userID,
		EncryptedPW: encryptedPW,
		Role:        body.Role,
		Spot:        body.Spot,
		Priority:    body.Priority,
	}

	err := user.Service.UpdateUser(userID, dto)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 User가 없습니다"))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary 사용자 정보 삭제
// @Description 사용자의 정보를 삭제합니다.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "User uuid"
// @Success 200 {object} response.CommonResponse
// @Router /user/{id} [delete]
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := user.Service.DeleteUser(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당되는 User이 존재하지 않습니다"))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}
