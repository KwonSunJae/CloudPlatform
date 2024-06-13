package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	checker "soms/controller/checker/container/service"
	response "soms/controller/response"

	"soms/service/container/service"

	"github.com/gorilla/mux"
)

type ServiceUseCase interface {
	getOneService(id string) (interface{}, error)
	getAllService() (interface{}, error)
	getServiceStatus() (interface{}, error)
	createService(serviceDto interface{}) error
	updateService(id string, serviceDto interface{}) error
	deleteService(id string) error
}

func ServiceController(router *mux.Router) error {
	err := service.Service.InitService()

	if err != nil {
		return err
	}

	router.HandleFunc("/service/{id}", getOneService).Methods("GET")

	router.HandleFunc("/service", getAllService).Methods("GET")

	router.HandleFunc("/servicestat", getServiceStatus).Methods("GET")

	router.HandleFunc("/service", createService).Methods("POST")

	router.HandleFunc("/service/{id}", updateService).Methods("PATCH")

	router.HandleFunc("/service/{id}", deleteService).Methods("DELETE")

	router.HandleFunc("/approve/service/{id}", approveService).Methods("POST")

	return nil
}

// @Summary service 정보 조회
// @Description service의 정보를 조회합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "service uuid"
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /service/{id} [get]
func getOneService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	raw, err := service.Service.GetOneService(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 Service가 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, raw, http.StatusOK, nil)

}

// @Summary service 정보 전체 조회
// @Description service의 정보를 전체 조회합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /service [get]
func getAllService(w http.ResponseWriter, r *http.Request) {
	raws, err := service.Service.GetAllService()

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, raws, http.StatusOK, nil)

}

// @Summary service 상태 조회
// @Description service의 상태를 조회합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /servicestat [get]
func getServiceStatus(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	rsp, err := service.Service.GetServiceStatus(uuid)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)

}

type ServiceRequestBody struct {
	// ClusterIP
	ApiVersion          string
	Kind                string
	MetadataName        string
	SpecType            string
	SpecSelectorApp     string
	SpecPortsProtocol   string
	SpecPortsPort       string
	SpecPortsTargetport string

	// NodePort
	SpecPortsNodeport string

	// LoadBalancer
	SpecSelectorType string
	SpecClusterIP    string

	// ExternalName
	SpecExternalname string
}

// @Summary service 생성
// @Description service를 생성합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param   body     body    ServiceRequestBody    true  "service 정보"
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /service [post]
func createService(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	var body ServiceRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	// 서비스타입별 바디에 파라미터 누락 확인
	checkerErr := checker.ServiceTypeChecker(checker.RequestBody{
		ApiVersion:          body.ApiVersion,
		Kind:                body.Kind,
		MetadataName:        body.MetadataName,
		SpecType:            body.SpecType,
		SpecSelectorApp:     body.SpecSelectorApp,
		SpecPortsProtocol:   body.SpecPortsProtocol,
		SpecPortsPort:       body.SpecPortsPort,
		SpecPortsTargetport: body.SpecPortsTargetport,
		SpecPortsNodeport:   body.SpecPortsNodeport,
		SpecSelectorType:    body.SpecSelectorType,
		SpecClusterIP:       body.SpecClusterIP,
		SpecExternalname:    body.SpecExternalname,
	})
	if checkerErr != nil {
		response.Response(w, nil, http.StatusBadRequest, checkerErr) // checker에서 반환된 에러를 전달
		return
	}
	dto := struct {
		ApiVersion          string
		Kind                string
		MetadataName        string
		SpecType            string
		SpecSelectorApp     string
		SpecPortsProtocol   string
		SpecPortsPort       string
		SpecPortsTargetport string
		SpecPortsNodeport   string
		SpecSelectorType    string
		SpecClusterIP       string
		SpecExternalname    string
		UUID                string
		Status              string
	}{
		ApiVersion:          body.ApiVersion,
		Kind:                body.Kind,
		MetadataName:        body.MetadataName,
		SpecType:            body.SpecType,
		SpecSelectorApp:     body.SpecSelectorApp,
		SpecPortsProtocol:   body.SpecPortsProtocol,
		SpecPortsPort:       body.SpecPortsPort,
		SpecPortsTargetport: body.SpecPortsTargetport,
		SpecPortsNodeport:   body.SpecPortsNodeport,
		SpecSelectorType:    body.SpecSelectorType,
		SpecClusterIP:       body.SpecClusterIP,
		SpecExternalname:    body.SpecExternalname,
		UUID:                uuid,
		Status:              "Pending",
	}
	err = service.Service.CreateService(dto)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary service 수정
// @Description service를 수정합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "service uuid"
// @Param   body     body    ServiceRequestBody    true  "service 정보"
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /service/{id} [patch]
func updateService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	var body ServiceRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}
	dto := struct {
		// ClusterIP
		ApiVersion          string
		Kind                string
		MetadataName        string
		SpecType            string
		SpecSelectorApp     string
		SpecPortsProtocol   string
		SpecPortsPort       string
		SpecPortsTargetport string

		// NodePort
		SpecPortsNodeport string

		// LoadBalancer
		SpecSelectorType string
		SpecClusterIP    string

		// ExternalName
		SpecExternalname string

		UUID   string
		Status string
	}{
		ApiVersion:          body.ApiVersion,
		Kind:                body.Kind,
		MetadataName:        body.MetadataName,
		SpecType:            body.SpecType,
		SpecSelectorApp:     body.SpecSelectorApp,
		SpecPortsProtocol:   body.SpecPortsProtocol,
		SpecPortsPort:       body.SpecPortsPort,
		SpecPortsTargetport: body.SpecPortsTargetport,
		SpecPortsNodeport:   body.SpecPortsNodeport,
		SpecSelectorType:    body.SpecSelectorType,
		SpecClusterIP:       body.SpecClusterIP,
		SpecExternalname:    body.SpecExternalname,
		UUID:                uuid,
		Status:              "Pending",
	}

	err = service.Service.UpdateService(id, dto)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 Service가 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary service 정보 삭제
// @Description service의 정보를 삭제합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "service uuid"
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /service/{id} [delete]
func deleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := service.Service.DeleteService(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당되는 Service가 존재하지 않습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary service 승인
// @Description service를 승인합니다.
// @Tags service
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "service uuid"
// @Param X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /approve/service/{id} [post]
func approveService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := service.Service.ApproveService(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당되는 Service가 존재하지 않습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)
}
