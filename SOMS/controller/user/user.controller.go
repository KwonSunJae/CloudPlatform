package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"soms/service/user"

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

func UserController(router *mux.Router) error {
	err := user.Service.InitService()

	if err != nil {
		return err
	}

	// GET 특정 id의 VM 데이터 반환
	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		raw, err := user.Service.GetOneUser(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 User가 없습니다"))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, raw, http.StatusOK, nil)

	}).Methods("GET")

	// GET 전체 VM 데이터 반환
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		raws, err := user.Service.GetAllUser()

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, raws, http.StatusOK, nil)

	}).Methods("GET")

	// POST 새로운 VM 등록
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Name       string
			UserID     string
			EncyptedPW string
			Role       string
			Spot       string
			Priority   string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		//checker 구현
		// if body.Name == "" || body.FlavorID == "" || body.ExternalIP == "" || body.InternalIP == "" ||
		// 	body.SelectedOS == "" || body.UnionmountImage == "" || body.Keypair == "" ||
		// 	body.SelectedSecuritygroup == "" || body.UserID == "" {
		// 	Response(w, nil, http.StatusBadRequest, errors.New("파라미터가 누락되었습니다."))
		// 	return
		// }

		err = user.Service.CreateUser(body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("POST")

	// PATCH 특정 id의 VM 데이터 수정
	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var body struct {
			Name       string
			UserID     string
			EncyptedPW string
			Role       string
			Spot       string
			Priority   string
		}

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			Response(w, nil, http.StatusInternalServerError, err)
		}

		err = user.Service.UpdateUser(id, body)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당 User가 없습니다"))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("PATCH")

	// DELETE 특정 id의 VM 데이터 삭제
	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err = user.Service.DeleteUser(id)

		if err != nil {
			switch err.Error() {
			case "NOT FOUND":
				Response(w, nil, http.StatusNotFound, errors.New("해당되는 User이 존재하지 않습니다"))
			default:
				Response(w, nil, http.StatusInternalServerError, err)
			}
			return
		}

		Response(w, "OK", http.StatusOK, nil)

	}).Methods("DELETE")

	return nil
}
