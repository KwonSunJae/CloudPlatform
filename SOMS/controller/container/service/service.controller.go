package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"os/exec"
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

	router.HandleFunc("/servicetest", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("terraform", "apply")
		cmd.Dir = "/home/ubuntu/test/"

		output, err := cmd.Output()

		if err != nil {
			Response(w, output, http.StatusOK, nil)
		} else {
			Response(w, err, http.StatusOK, nil)
		}

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
			ApiVersion            string
			Kind                  string
			Metadata_name         string
			Spec_ports_port       string
			Spec_ports_protocol   string
			Spec_ports_targetPort string
			Spec_selector_app     string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		if body.Metadata_name == "" || body.Spec_selector_app == "" || body.Spec_ports_protocol == "" ||
			body.Spec_ports_port == "" || body.Spec_ports_targetPort == "" || body.ApiVersion == "" || body.Kind == "" {
			Response(w, nil, http.StatusBadRequest, errors.New("파라미터가 누락되었습니다."))
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
			ApiVersion            string
			Kind                  string
			Metadata_name         string
			Spec_ports_port       string
			Spec_ports_protocol   string
			Spec_ports_targetPort string
			Spec_selector_app     string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

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
	router.HandleFunc("/deplooyment/{id}", func(w http.ResponseWriter, r *http.Request) {
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
