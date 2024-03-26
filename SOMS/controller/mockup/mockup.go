package mockup

import (
	"encoding/json"
	"errors"
	"net/http"
	"soms/controller/response"

	"github.com/gorilla/mux"
)

func MockupController(router *mux.Router) error {
	router.HandleFunc("/mockup/200", Handle200OKAPI).Methods("GET")
	router.HandleFunc("/mockup/200", Handle200OKPost).Methods("POST")
	router.HandleFunc("/mockup/400", Handle40XErrorAPI).Methods("GET")
	router.HandleFunc("/mockup/400", Handle40XErrorPost).Methods("POST")
	router.HandleFunc("/mockup/500", Handle50XErrorAPI).Methods("GET")
	router.HandleFunc("/mockup/500", Handle50XErrorPost).Methods("POST")
	return nil
}

// Handle200OKAPI는 200 OK 응답만을 던지는 API 핸들러입니다.
// @Summary 200 OK 응답
// @Description 200 OK 응답을 반환합니다.
// @Tags mockup
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse
// @Router /mockup/200 [get]
func Handle200OKAPI(w http.ResponseWriter, r *http.Request) {
	response.Response(w, "200OK", http.StatusOK, nil)
}

// Handle40XErrorAPI는 40X 에러만을 던지는 API 핸들러입니다.
// @Summary 40X 에러 응답
// @Description 40X 에러 응답을 반환합니다.
// @Tags mockup
// @Accept  json
// @Produce  json
// @Success 400 {object} response.CommonResponse
// @Router /mockup/400 [get]
func Handle40XErrorAPI(w http.ResponseWriter, r *http.Request) {
	response.Response(w, "400 BADREQUEST", http.StatusBadRequest, errors.New("this is bad request"))
}

// Handle50XErrorAPI는 50X 에러만을 던지는 API 핸들러입니다.
// @Summary 50X 에러 응답
// @Description 50X 에러 응답을 반환합니다.
// @Tags mockup
// @Accept  json
// @Produce  json
// @Success 500 {object} response.CommonResponse
// @Router /mockup/500 [get]
func Handle50XErrorAPI(w http.ResponseWriter, r *http.Request) {
	response.Response(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError, errors.New("this is internal server error"))
}

// Handle200OKPost는 200 OK 응답만을 던지는 API 핸들러입니다.
// @Summary 200 OK 응답
// @Description 200 OK 응답을 반환합니다.
// @Tags mockup
// @Accept  json
// @Produce  json
// @Param   body    body    string     true  "body"
// @Success 200 {object} response.CommonResponse
// @Router /mockup/200 [post]
func Handle200OKPost(w http.ResponseWriter, r *http.Request) {
	var body string
	err := json.NewDecoder(r.Body).Decode(&body)

	response.Response(w, "200OK", http.StatusOK, err)
}

// Handle40XErrorPost는 40X 에러만을 던지는 API 핸들러입니다.
// @Summary 40X 에러 응답
// @Description 40X 에러 응답을 반환합니다.
// @Tags mockup
// @Accept  json
// @Produce  json
// @Param   body    body    string     true  "body"
// @Success 400 {object} response.CommonResponse
// @Router /mockup/400 [post]
func Handle40XErrorPost(w http.ResponseWriter, r *http.Request) {
	var body string
	err := json.NewDecoder(r.Body).Decode(&body)

	response.Response(w, "400 BADREQUEST", http.StatusBadRequest, err)
}

// Handle50XErrorPost는 50X 에러만을 던지는 API 핸들러입니다.
// @Summary 50X 에러 응답
// @Description 50X 에러 응답을 반환합니다.
// @Tags mockup
// @Accept  json
// @Produce  json
// @Param   body    body    string     true  "body"
// @Success 500 {object} response.CommonResponse
// @Router /mockup/500 [post]
func Handle50XErrorPost(w http.ResponseWriter, r *http.Request) {
	var body string
	err := json.NewDecoder(r.Body).Decode(&body)
	response.Response(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError, err)
}
