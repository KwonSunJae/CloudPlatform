package deployment

import (
	"encoding/json"
	"errors"
	"net/http"
	reqchecker "soms/controller/checker"
	response "soms/controller/response"
	"soms/service/container/deployment"

	"github.com/gorilla/mux"
)

type DeploymentUseCase interface {
	getOneDeployment(id string) (interface{}, error)
	getAllDeployment() (interface{}, error)
	getDeploymentsStatus() (interface{}, error)
	createDeployment(deploymentDto interface{}) error
	updateDeployment(id string, deploymentDto interface{}) error
	deleteDeployment(id string) error
}

func DeploymentController(router *mux.Router) error {
	err := deployment.Service.InitService()

	if err != nil {
		return err
	}

	router.HandleFunc("/deployment/{id}", getOneDeployment).Methods("GET")

	router.HandleFunc("/deployment", getAllDeployment).Methods("GET")

	router.HandleFunc("/deploymentstat", getDeploymentsStatus).Methods("GET")

	router.HandleFunc("/deployment", createDeployment).Methods("POST")

	router.HandleFunc("/deployment/{id}", updateDeployment).Methods("PATCH")

	router.HandleFunc("/deployment/{id}", deleteDeployment).Methods("DELETE")

	router.HandleFunc("/approve/deployment/{id}", approveDeployment).Methods("POST")

	return nil
}

// @Summary deployment 정보 조회
// @Description deployment의 정보를 조회합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "deployment uuid"
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/{id} [get]
func getOneDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	raw, err := deployment.Service.GetOneDeployment(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 Deployment가 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, raw, http.StatusOK, nil)

}

// @Summary deployment 정보 전체 조회
// @Description deployment의 정보를 전체 조회합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /deployment [get]
func getAllDeployment(w http.ResponseWriter, r *http.Request) {

	raws, err := deployment.Service.GetAllDeployment()

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, raws, http.StatusOK, nil)

}

// @Summary deployment 상태 조회
// @Description deployment의 상태를 조회합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /deploymentstat [get]
func getDeploymentsStatus(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")

	rsp, err := deployment.Service.GetDeploymentsStatus(uuid)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)

}

type createDeploymentBody struct {
	ApiVersion                                   string
	Kind                                         string
	MetadataName                                 string
	MetadataLabelsApp                            string
	SpecReplicas                                 string
	SpecSelectorMatchlabelsApp                   string
	SpecTemplateMetadataLabelsApp                string
	SpecTemplateSpecContainersName               string
	SpecTemplateSpecContainersImage              string
	SpecTemplateSpecContainersPortsContainerport string
}

// @Summary deployment 등록
// @Description deployment를 등록합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param   body     body    createDeploymentBody     true  "deployment 정보"
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /deployment [post]
func createDeployment(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	var body createDeploymentBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
	}

	prmErr := reqchecker.Check(body)
	if prmErr != nil {
		response.Response(w, nil, http.StatusBadRequest, prmErr)
		return
	}

	var deploymentDto = struct {
		ApiVersion                                   string
		Kind                                         string
		MetadataName                                 string
		MetadataLabelsApp                            string
		SpecReplicas                                 string
		SpecSelectorMatchlabelsApp                   string
		SpecTemplateMetadataLabelsApp                string
		SpecTemplateSpecContainersName               string
		SpecTemplateSpecContainersImage              string
		SpecTemplateSpecContainersPortsContainerport string
		UUID                                         string
		Status                                       string
	}{
		ApiVersion:                      body.ApiVersion,
		Kind:                            body.Kind,
		MetadataName:                    body.MetadataName,
		MetadataLabelsApp:               body.MetadataLabelsApp,
		SpecReplicas:                    body.SpecReplicas,
		SpecSelectorMatchlabelsApp:      body.SpecSelectorMatchlabelsApp,
		SpecTemplateMetadataLabelsApp:   body.SpecTemplateMetadataLabelsApp,
		SpecTemplateSpecContainersName:  body.SpecTemplateSpecContainersName,
		SpecTemplateSpecContainersImage: body.SpecTemplateSpecContainersImage,
		SpecTemplateSpecContainersPortsContainerport: body.SpecTemplateSpecContainersPortsContainerport,
		UUID:   uuid,
		Status: "Pending",
	}

	err = deployment.Service.EnrollDeployment(deploymentDto)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary deployment 수정
// @Description deployment를 수정합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "deployment uuid"
// @Param   body     body    createDeploymentBody     true  "deployment 정보"
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/{id} [patch]
func updateDeployment(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	var body createDeploymentBody
	vars := mux.Vars(r)
	id := vars["id"]

	var deploymentDto = struct {
		ApiVersion                                   string
		Kind                                         string
		MetadataName                                 string
		MetadataLabelsApp                            string
		SpecReplicas                                 string
		SpecSelectorMatchlabelsApp                   string
		SpecTemplateMetadataLabelsApp                string
		SpecTemplateSpecContainersName               string
		SpecTemplateSpecContainersImage              string
		SpecTemplateSpecContainersPortsContainerport string
		UUID                                         string
		Status                                       string
	}{
		ApiVersion:                      body.ApiVersion,
		Kind:                            body.Kind,
		MetadataName:                    body.MetadataName,
		MetadataLabelsApp:               body.MetadataLabelsApp,
		SpecReplicas:                    body.SpecReplicas,
		SpecSelectorMatchlabelsApp:      body.SpecSelectorMatchlabelsApp,
		SpecTemplateMetadataLabelsApp:   body.SpecTemplateMetadataLabelsApp,
		SpecTemplateSpecContainersName:  body.SpecTemplateSpecContainersName,
		SpecTemplateSpecContainersImage: body.SpecTemplateSpecContainersImage,
		SpecTemplateSpecContainersPortsContainerport: body.SpecTemplateSpecContainersPortsContainerport,
		UUID:   uuid,
		Status: "Pending",
	}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
	}

	err = deployment.Service.UpdateDeployment(id, deploymentDto)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 Deployment가 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary deployment 정보 삭제
// @Description deployment의 정보를 삭제합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "deployment uuid"
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /deployment/{id} [delete]
func deleteDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := deployment.Service.DeleteDeployment(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당되는 Deployment가 존재하지 않습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary deployment 승인
// @Description deployment를 승인합니다.
// @Tags deployment
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "deployment uuid"
// @Param  X-UUID header  string     true  "User UUID"
// @Success 200 {object} response.CommonResponse
// @Router /approve/deployment/{id} [post]
func approveDeployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := deployment.Service.ApproveDeployment(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 Deployment가 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)
}
