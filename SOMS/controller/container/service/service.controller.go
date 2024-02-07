package service

import (
	"encoding/json"
	"errors"
	"net/http"
	checker "soms/controller/checker/container/service"
	"soms/service/container/service"

	"github.com/gorilla/mux"
)

type CommonResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

func Response(w http.ResponseWriter, data interface{}, status int, err error) {
	var res CommonResponse

	if status == http.StatusOK {
		res.Data = data
		res.Status = status
	} else {
		res.Status = status
		res.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func ServiceController(router *mux.Router) error {
	err := service.Service.InitService()

	if err != nil {
		return err
	}

	// GET 특정 id의 Service 데이터 반환
	router.HandleFunc("/service/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		raw, err := service.Service.GetOneService(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 Service가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, raw, http.StatusOK, nil)

	}).Methods("GET")

	// GET 전체 Service 데이터 반환
	router.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		raws, err := service.Service.GetAllService()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, raws, http.StatusOK, nil)

	}).Methods("GET")

	// GET Service status 반환
	router.HandleFunc("/servicestat", func(w http.ResponseWriter, r *http.Request) {
		rsp, err := service.Service.GetServiceStatus()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, rsp, http.StatusOK, nil)

	}).Methods("GET")

	// POST 새로운 Service 등록
	router.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
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
			MetadataNamespace string
			SpecExternalname  string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		// 서비스타입별 바디에 파라미터 누락 표시
		checkerErr := checker.ServiceTypeChecker(body)
		if checkerErr != nil {
			Response(w, nil, http.StatusBadRequest, checkerErr) // checker에서 반환된 에러를 전달
			return
		}

		err = service.Service.CreateService(body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("POST")

	// PATCH 특정 id의 Service 데이터 수정
	router.HandleFunc("/service/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var body struct {
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
			MetadataNamespace string
			SpecExternalname  string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		Response(w, nil, http.StatusBadRequest, checker.ServiceTypeChecker(body)) // 서비스타입별 바디에 파라미터 누락 표시

		err = service.Service.UpdateService(id, body)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 Service가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("PATCH")

	// DELETE 특정 id의 Service 데이터 삭제
	router.HandleFunc("/service/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err = service.Service.DeleteService(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당되는 Service가 존재하지 않습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("DELETE")

	return nil
}
