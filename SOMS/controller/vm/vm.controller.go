package vm

import (
	"encoding/json"
	"errors"
	"net/http"

	"soms/service/vm"

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

func VmController(router *mux.Router) error {
	err := vm.Service.InitService()

	if err != nil {
		return err
	}

	// GET 특정 id의 뉴스 데이터 반환
	router.HandleFunc("/Vm/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		raw, err := vm.Service.GetOneVm(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 뉴스가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, raw, http.StatusOK, nil)

	}).Methods("GET")

	// GET 전체 뉴스 데이터 반환
	router.HandleFunc("/vm", func(w http.ResponseWriter, r *http.Request) {
		raws, err := vm.Service.GetAllVm()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, raws, http.StatusOK, nil)

	}).Methods("GET")

	// POST 새로운 뉴스 등록
	router.HandleFunc("/vm", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Title   string
			Author  string
			Content string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		if body.Title == "" || body.Author == "" || body.Content == "" {
			Response(w, nil, http.StatusBadRequest, errors.New("파라미터가 누락되었습니다."))
			return
		}

		err = vm.Service.CreateVm(body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("POST")

	// PATCH 특정 id의 뉴스 데이터 수정
	router.HandleFunc("/vm/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var body struct {
			Title   string
			Author  string
			Content string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		err = vm.Service.UpdateVm(id, body)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 뉴스가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("PATCH")

	// DELETE 특정 id의 뉴스 데이터 삭제
	router.HandleFunc("/Vm/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err = vm.Service.DeleteVm(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 뉴스가 없습니다."))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("DELETE")

	return nil
}
