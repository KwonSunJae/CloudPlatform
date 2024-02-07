package replicaset

import (
	"encoding/json"
	"errors"
	"net/http"
	"soms/service/container/replicaset"

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

func ReplicasetController(router *mux.Router) error {
	err := replicaset.Service.InitService()

	if err != nil {
		return err
	}

	// GET 특정 id의 Replicaset 데이터 반환
	router.HandleFunc("/replicaset/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		raw, err := replicaset.Service.GetOneReplicaset(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 Replicaset이 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, raw, http.StatusOK, nil)

	}).Methods("GET")

	// GET 전체 Replicaset 데이터 반환
	router.HandleFunc("/replicaset", func(w http.ResponseWriter, r *http.Request) {
		raws, err := replicaset.Service.GetAllReplicaset()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, raws, http.StatusOK, nil)

	}).Methods("GET")

	// GET Replicaset status 반환
	router.HandleFunc("/replicasetstat", func(w http.ResponseWriter, r *http.Request) {
		rsp, err := replicaset.Service.GetReplicasetsStatus()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, rsp, http.StatusOK, nil)

	}).Methods("GET")

	// POST 새로운 Replicaset 등록
	router.HandleFunc("/replicaset", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			ApiVersion                                   string
			Kind                                         string
			MetadataName                                 string
			SpecReplicas                                 string
			SpecSelectorMatchlabelsApp                   string
			SpecTemplateMetadataName                     string
			SpecTemplateMetadataLabelsApp                string
			SpecTemplateSpecContainersName               string
			SpecTemplateSpecContainersImage              string
			SpecTemplateSpecContainersPortsContainerport string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		if body.ApiVersion == "" || body.Kind == "" || body.MetadataName == "" || body.SpecReplicas == "" ||
			body.SpecSelectorMatchlabelsApp == "" || body.SpecTemplateMetadataName == "" || body.SpecTemplateMetadataLabelsApp == "" ||
			body.SpecTemplateSpecContainersName == "" || body.SpecTemplateSpecContainersImage == "" || body.SpecTemplateSpecContainersPortsContainerport == "" {
			Response(w, nil, http.StatusBadRequest, errors.New("파라미터가 누락되었습니다."))
			return
		}

		err = replicaset.Service.CreateReplicaset(body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("POST")

	// PATCH 특정 id의 Replicaset 데이터 수정
	router.HandleFunc("/replicaset/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var body struct {
			ApiVersion                                   string
			Kind                                         string
			MetadataName                                 string
			SpecReplicas                                 string
			SpecSelectorMatchlabelsApp                   string
			SpecTemplateMetadataName                     string
			SpecTemplateMetadataLabelsApp                string
			SpecTemplateSpecContainersName               string
			SpecTemplateSpecContainersImage              string
			SpecTemplateSpecContainersPortsContainerport string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		err = replicaset.Service.UpdateReplicaset(id, body)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 Replicaset이 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("PATCH")

	// DELETE 특정 id의 Replicaset 데이터 삭제
	router.HandleFunc("/replicaset/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err = replicaset.Service.DeleteReplicaset(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당되는 Replicaset이 존재하지 않습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("DELETE")

	return nil
}
