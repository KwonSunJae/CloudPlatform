package vm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	reqchecker "soms/controller/checker"
	response "soms/controller/response"
	"soms/service/vm"

	"github.com/gorilla/mux"
)

type VMUseCase interface {
	getVmById(w http.ResponseWriter, r *http.Request)
	getAllVm(w http.ResponseWriter, r *http.Request)
	getVmStatus(w http.ResponseWriter, r *http.Request)
	createVm(w http.ResponseWriter, r *http.Request)
	updateVm(w http.ResponseWriter, r *http.Request)
	deleteVm(w http.ResponseWriter, r *http.Request)
	startVm(w http.ResponseWriter, r *http.Request)
	stopVm(w http.ResponseWriter, r *http.Request)
	rebootVm(w http.ResponseWriter, r *http.Request)
	getVmConsole(w http.ResponseWriter, r *http.Request)
	getVmVnc(w http.ResponseWriter, r *http.Request)
}

func VmController(router *mux.Router) error {
	err := vm.Service.InitService()

	if err != nil {
		return err
	}

	router.HandleFunc("/vm/{id}", getVmById).Methods("GET")

	router.HandleFunc("/vm", getAllVm).Methods("GET")

	router.HandleFunc("/vmstat", getVmStatus).Methods("GET")

	router.HandleFunc("/vm", createVm).Methods("POST")

	router.HandleFunc("/vm/{id}", updateVm).Methods("PATCH")

	router.HandleFunc("/vm/{id}", deleteVm).Methods("DELETE")

	router.HandleFunc("/vm/start/{id}", startVm).Methods("POST")

	router.HandleFunc("/vm/stop/{id}", stopVm).Methods("POST")

	router.HandleFunc("/vm/reboot/{id}", rebootVm).Methods("POST")

	router.HandleFunc("/vm/status/{id}", getVmStatus).Methods("GET")

	router.HandleFunc("/vm/snapshot/{id}", getVmConsole).Methods("GET")

	router.HandleFunc("/vm/console/{id}", getVmConsole).Methods("GET")

	router.HandleFunc("/vm/vnc/{id}", getVmVnc).Methods("GET")

	return nil
}

// @Summary VM 정보 조회
// @Description VM의 정보를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/{id} [get]
func getVmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	raw, err := vm.Service.GetOneVm(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM가 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, raw, http.StatusOK, nil)

}

// @Summary VM 정보 전체 조회
// @Description VM의 정보를 전체 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse
// @Router /vm [get]
func getAllVm(w http.ResponseWriter, r *http.Request) {
	raws, err := vm.Service.GetAllVm()

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, raws, http.StatusOK, nil)

}

// @Summary VM 상태 조회
// @Description VM의 상태를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse
// @Router /vmstat [get]
func getVmStatus(w http.ResponseWriter, r *http.Request) {
	rsp, err := vm.Service.GetStatusVM("test")

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)

}

type CreateVmBody struct {
	Name                  string
	FlavorID              string
	ExternalIP            string
	InternalIP            string
	SelectedOS            string
	UnionmountImage       string
	Keypair               string
	SelectedSecuritygroup string
	UserID                string
}

// @Summary VM 등록
// @Description VM을 등록합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   body     body    CreateVmBody     true  "VM 정보"
// @Success 200 {object} response.CommonResponse
// @Router /vm [post]
func createVm(w http.ResponseWriter, r *http.Request) {
	var body CreateVmBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
	}

	prmErr := reqchecker.Check(body)
	if prmErr != nil {
		response.Response(w, nil, http.StatusBadRequest, prmErr)
		return
	}
	var vmDto = struct {
		Name                  string
		FlavorID              string
		ExternalIP            string
		InternalIP            string
		SelectedOS            string
		UnionmountImage       string
		Keypair               string
		SelectedSecuritygroup string
		UserID                string
	}{
		Name:                  body.Name,
		FlavorID:              body.FlavorID,
		ExternalIP:            body.ExternalIP,
		InternalIP:            body.InternalIP,
		SelectedOS:            body.SelectedOS,
		UnionmountImage:       body.UnionmountImage,
		Keypair:               body.Keypair,
		SelectedSecuritygroup: body.SelectedSecuritygroup,
		UserID:                body.UserID,
	}
	err = vm.Service.CreateVm(vmDto)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 정보 수정
// @Description VM의 정보를 수정합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param   body     body    CreateVmBody     true  "VM 정보"
// @Success 200 {object} response.CommonResponse
// @Router /vm/{id} [patch]
func updateVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var body CreateVmBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
	}

	var vmDto = struct {
		Name                  string
		FlavorID              string
		ExternalIP            string
		InternalIP            string
		SelectedOS            string
		UnionmountImage       string
		Keypair               string
		SelectedSecuritygroup string
		UserID                string
	}{
		Name:                  body.Name,
		FlavorID:              body.FlavorID,
		ExternalIP:            body.ExternalIP,
		InternalIP:            body.InternalIP,
		SelectedOS:            body.SelectedOS,
		UnionmountImage:       body.UnionmountImage,
		Keypair:               body.Keypair,
		SelectedSecuritygroup: body.SelectedSecuritygroup,
		UserID:                body.UserID,
	}
	err = vm.Service.UpdateVm(id, vmDto)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 정보 삭제
// @Description VM의 정보를 삭제합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/{id} [delete]
func deleteVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := vm.Service.DeleteVm(id)

	if err != nil {
		switch err.Error() {
		case "NOT FOUND":
			response.Response(w, nil, http.StatusNotFound, errors.New("해당되는 VM이 존재하지 않습니다."))
		default:
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 시작
// @Description VM을 시작합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/start/{id} [post]
func startVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("startVm id: ", id)
	//err := vm.Service.StartVm(id)

	// if err != nil {
	// 	switch err.Error() {
	// 	case "NOT FOUND":
	// 		response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
	// 	default:
	// 		response.Response(w, nil, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 중지
// @Description VM을 중지합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/stop/{id} [post]
func stopVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("stopVm id: ", id)
	// err := vm.Service.StopVm(id)

	// if err != nil {
	// 	switch err.Error() {
	// 	case "NOT FOUND":
	// 		response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
	// 	default:
	// 		response.Response(w, nil, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 재부팅
// @Description VM을 재부팅합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/reboot/{id} [post]
func rebootVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("rebootVm id: ", id)
	// err := vm.Service.RebootVm(id)

	// if err != nil {
	// 	switch err.Error() {
	// 	case "NOT FOUND":
	// 		response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
	// 	default:
	// 		response.Response(w, nil, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 콘솔 조회
// @Description VM의 콘솔을 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/console/{id} [get]
func getVmConsole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("getVmConsole id: ", id)
	// rsp, err := vm.Service.GetVmConsole(id)

	// if err != nil {
	// 	switch err.Error() {
	// 	case "NOT FOUND":
	// 		response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
	// 	default:
	// 		response.Response(w, nil, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM VNC 조회
// @Description VM의 VNC를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Success 200 {object} response.CommonResponse
// @Router /vm/vnc/{id} [get]
func getVmVnc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("getVmVnc id: ", id)
	// rsp, err := vm.Service.GetVmVnc(id)

	// if err != nil {
	// 	switch err.Error() {
	// 	case "NOT FOUND":
	// 		response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
	// 	default:
	// 		response.Response(w, nil, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	response.Response(w, "OK", http.StatusOK, nil)

}
