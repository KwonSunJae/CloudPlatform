package deployment

import (
	"encoding/json"
	"errors"
	"net/http"
	"os/exec"
	"soms/service/container/deployment"

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

func DeploymentController(router *mux.Router) error {
	err := deployment.Service.InitService()

	if err != nil {
		return err
	}

	// GET 특정 id의 Deployment 데이터 반환
	router.HandleFunc("/deployment/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		raw, err := deployment.Service.GetOneDeployment(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 Deployment가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, raw, http.StatusOK, nil)

	}).Methods("GET")

	router.HandleFunc("/deploymenttest", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("terraform", "apply")
		cmd.Dir = "/home/ubuntu/test/"

		output, err := cmd.Output()

		if err != nil {
			Response(w, output, http.StatusOK, nil)
		} else {
			Response(w, err, http.StatusOK, nil)
		}

	}).Methods("GET")

	// GET 전체 Deployment 데이터 반환
	router.HandleFunc("/deployment", func(w http.ResponseWriter, r *http.Request) {
		raws, err := deployment.Service.GetAllDeployment()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, raws, http.StatusOK, nil)

	}).Methods("GET")

	// POST 새로운 Deployment 등록
	router.HandleFunc("/deployment", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			ApiVersion                                        string
			Kind                                              string
			Metadata_name                                     string
			Metadata_labels_app                               string
			Spec_selector_matchLabels_app                     string
			Spec_template_metadata_labels_app                 string
			Spec_template_spec_hostname                       string
			Spec_template_spec_subdomain                      string
			Spec_template_spec_containers_image               string
			Spec_template_spec_containers_imagePullPolicy     string
			Spec_template_spec_containers_name                string
			Spec_template_spec_containers_ports_containerPort string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}
		if body.Metadata_name == "" || body.Metadata_labels_app == "" || body.Spec_selector_matchLabels_app == "" || body.Spec_template_metadata_labels_app == "" ||
			body.Spec_template_spec_hostname == "" || body.Spec_template_spec_subdomain == "" || body.Spec_template_spec_containers_image == "" ||
			body.Spec_template_spec_containers_imagePullPolicy == "" || body.Spec_template_spec_containers_name == "" || body.Spec_template_spec_containers_ports_containerPort == "" {
			Response(w, nil, http.StatusBadRequest, errors.New("파라미터가 누락되었습니다."))
			return
		}

		err = deployment.Service.CreateDeployment(body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("POST")

	// PATCH 특정 id의 Deployment 데이터 수정
	router.HandleFunc("/deployment/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var body struct {
			ApiVersion                                        string
			Kind                                              string
			Metadata_name                                     string
			Metadata_labels_app                               string
			Spec_selector_matchLabels_app                     string
			Spec_template_metadata_labels_app                 string
			Spec_template_spec_hostname                       string
			Spec_template_spec_subdomain                      string
			Spec_template_spec_containers_image               string
			Spec_template_spec_containers_imagePullPolicy     string
			Spec_template_spec_containers_name                string
			Spec_template_spec_containers_ports_containerPort string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		// ApiVersion와 Kind를 추가하여 DeploymentDto 생성

		err = deployment.Service.UpdateDeployment(id, body)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 Deployment가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("PATCH")

	// DELETE 특정 id의 Deployment 데이터 삭제
	router.HandleFunc("/deplooyment/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err = deployment.Service.DeleteDeployment(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당되는 Deployment가 존재하지 않습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("DELETE")

	return nil
}
