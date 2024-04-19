package integrity

import (
	"errors"
	"net/http"
	response "soms/controller/response"
	"soms/service/integrity"

	"github.com/gorilla/mux"
)

func IntegrityController(router *mux.Router) error {
	err := integrity.Service.InitService()

	if err != nil {
		return err
	}
	router.HandleFunc("/integrity/{request_id}", getIntegrityByUuid).Methods("GET")
	router.HandleFunc("/integrity", getAllIntegrity).Methods("GET")
	router.HandleFunc("/integrity/user/{user_id}", getIntegrityByUserID).Methods("GET")

	return nil
}

// @Summary 무결성 단일 정보 조회 By Request ID
// @Description 무결성 정보를 조회합니다.
// @Tags integrity
// @Accept  json
// @Produce  json
// @Param   request_id     path    string     true  "Request uuid"
// @Success 200 {object} response.CommonResponse
// @Router /integrity/{request_id} [get]
func getIntegrityByUuid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	request_id := vars["request_id"]

	raw, err := integrity.Service.GetOneIntegrity(request_id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 Request가 없습니다"))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, raw, http.StatusOK, nil)
}

// @Summary 무결성 정보 전체 조회
// @Description 무결성 정보를 전체 조회합니다.
// @Tags integrity
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse
// @Router /integrity [get]
func getAllIntegrity(w http.ResponseWriter, r *http.Request) {
	raws, err := integrity.Service.GetAllIntegrity()

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, raws, http.StatusOK, nil)
}

// @Summary 사용자별 무결성 정보 조회 By User ID
// @Description 사용자별 무결성 정보를 조회합니다.
// @Tags integrity
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true  "User uuid"
// @Success 200 {object} response.CommonResponse
// @Router /integrity/user/{user_id} [get]
func getIntegrityByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	raws, err := integrity.Service.GetIntegrityByUserID(user_id)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, raws, http.StatusOK, nil)
}
