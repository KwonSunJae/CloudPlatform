package vm

import (
	"encoding/json"
	"errors"
	"net/http"
	reqchecker "soms/controller/checker"
	"soms/controller/checker/authority"
	response "soms/controller/response"
	"soms/service/vm"

	"github.com/gorilla/mux"
)

type VMUseCase interface {
	getVmById(w http.ResponseWriter, r *http.Request)
	getAllVm(w http.ResponseWriter, r *http.Request)
	getVmStatus(w http.ResponseWriter, r *http.Request)
	enrollVm(w http.ResponseWriter, r *http.Request)
	updateVm(w http.ResponseWriter, r *http.Request)
	deleteVm(w http.ResponseWriter, r *http.Request)
	approveVMCreation(w http.ResponseWriter, r *http.Request)
	powerOnVm(w http.ResponseWriter, r *http.Request)
	powerOffVm(w http.ResponseWriter, r *http.Request)
	softrebootVm(w http.ResponseWriter, r *http.Request)
	hardrebootVm(w http.ResponseWriter, r *http.Request)
	snapshotVm(w http.ResponseWriter, r *http.Request)
	getNetworkList(w http.ResponseWriter, r *http.Request)
	createNetwork(w http.ResponseWriter, r *http.Request)
	getFlavorList(w http.ResponseWriter, r *http.Request)
	createKeypair(w http.ResponseWriter, r *http.Request)
	getKeypairList(w http.ResponseWriter, r *http.Request)
	getSecurityGroupList(w http.ResponseWriter, r *http.Request)
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

	router.HandleFunc("/vm", enrollVm).Methods("POST")

	router.HandleFunc("/vm/{id}", updateVm).Methods("PATCH")

	router.HandleFunc("/vm/{id}", deleteVm).Methods("DELETE")

	router.HandleFunc("/action/approve/{id}", approveVMCreation).Methods("POST")

	router.HandleFunc("/action/poweron/{id}", powerOnVm).Methods("POST")

	router.HandleFunc("/action/poweroff/{id}", powerOffVm).Methods("POST")

	router.HandleFunc("/action/softreboot/{id}", softrebootVm).Methods("POST")

	router.HandleFunc("/action/hardreboot/{id}", hardrebootVm).Methods("POST")

	router.HandleFunc("/action/snapshot/{id}", snapshotVm).Methods("POST")

	router.HandleFunc("/resource/network", getNetworkList).Methods("GET")

	router.HandleFunc("/resource/network", createNetwork).Methods("POST")

	router.HandleFunc("/resource/flavor", getFlavorList).Methods("GET")

	router.HandleFunc("/resource/keypair", createKeypair).Methods("POST")

	router.HandleFunc("/resource/keypair", getKeypairList).Methods("GET")
	//router.HandleFunc("/vm/securitygroup", createSecurityGroup).Methods("POST")
	router.HandleFunc("/resource/securitygroup", getSecurityGroupList).Methods("GET")

	router.HandleFunc("/status/vm/{id}", getVmStatus).Methods("GET")

	router.HandleFunc("/action/vnc/{id}", getVmVnc).Methods("GET")

	return nil
}

// @Summary VM 정보 조회
// @Description VM의 정보를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /vm/{id} [get]
func getVmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//uuid := r.Header.Get("X-UUID")
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
// @Param X-UUID header string true "UUID"
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
// @Param X-UUID header string true "UUID"
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
}

// @Summary VM 등록
// @Description VM을 등록합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   body     body    CreateVmBody     true  "VM 정보"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /vm [post]
func enrollVm(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	if !authority.AuthorityFilterWithRole([]string{"Admin", "Master", "Student", "Researcher"}, uuid) {
		response.Response(w, nil, http.StatusUnauthorized, errors.New("권한이 없습니다"))
		return
	}
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
		UUID                  string
		Status                string
	}{
		Name:                  body.Name,
		FlavorID:              body.FlavorID,
		ExternalIP:            body.ExternalIP,
		InternalIP:            body.InternalIP,
		SelectedOS:            body.SelectedOS,
		UnionmountImage:       body.UnionmountImage,
		Keypair:               body.Keypair,
		SelectedSecuritygroup: body.SelectedSecuritygroup,
		UUID:                  uuid,
		Status:                "Pending",
	}
	err = vm.Service.EnrollVm(vmDto)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, "OK", http.StatusOK, nil)

}

// @Summary VM 생성 승인
// @Description VM 생성을 승인합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /vm/approve/{id} [post]
func approveVMCreation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")
	err := vm.Service.ApproveVMCreation(id, uuid)

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
	uuid := r.Header.Get("X-UUID")

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
		UUID                  string
		Status                string
	}{
		Name:                  body.Name,
		FlavorID:              body.FlavorID,
		ExternalIP:            body.ExternalIP,
		InternalIP:            body.InternalIP,
		SelectedOS:            body.SelectedOS,
		UnionmountImage:       body.UnionmountImage,
		Keypair:               body.Keypair,
		SelectedSecuritygroup: body.SelectedSecuritygroup,
		UUID:                  uuid,
		Status:                "Pending",
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
// @Param   id     path    string     true  "VM id"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /vm/{id} [delete]
func deleteVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	uuid := r.Header.Get("X-UUID")
	// if !authority.AuthorityFilterWithRole([]string{"Admin", "Master"}, uuid) {
	// 	response.Response(w, nil, http.StatusUnauthorized, errors.New("권한이 없습니다"))
	// 	return
	// }

	err := vm.Service.DeleteVm(id, uuid)

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
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /action/poweron/{id} [post]
func powerOnVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	result, err := vm.Service.PowerOn(uuid, id)
	if err != nil {
		if err == errors.New("NOT FOUND") {
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
		} else {
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	if result {
		response.Response(w, "OK", http.StatusOK, nil)
	} else {
		response.Response(w, nil, http.StatusInternalServerError, errors.New("해당 VM을 시작할 수 없습니다."))
	}
	return
}

// @Summary VM 종료
// @Description VM을 종료합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /action/poweroff/{id} [post]
func powerOffVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	result, err := vm.Service.PowerOff(uuid, id)
	if err != nil {
		if err == errors.New("NOT FOUND") {
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
		} else {
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	if result {
		response.Response(w, "OK", http.StatusOK, nil)
	} else {
		response.Response(w, nil, http.StatusInternalServerError, errors.New("해당 VM을 종료할 수 없습니다."))
	}
	return
}

// @Summary VM 소프트 리부팅
// @Description VM을 소프트 리부팅합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /action/softreboot/{id} [post]
func softrebootVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	result, err := vm.Service.SoftReboot(uuid, id)
	if err != nil {
		if err == errors.New("NOT FOUND") {
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
		} else {
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	if result {
		response.Response(w, "OK", http.StatusOK, nil)
	} else {
		response.Response(w, nil, http.StatusInternalServerError, errors.New("해당 VM을 소프트 리부팅할 수 없습니다."))
	}
	return
}

// @Summary VM 하드 리부팅
// @Description VM을 하드 리부팅합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /action/hardreboot/{id} [post]
func hardrebootVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	result, err := vm.Service.HardReboot(uuid, id)
	if err != nil {
		if err == errors.New("NOT FOUND") {
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
		} else {
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	if result {
		response.Response(w, "OK", http.StatusOK, nil)
	} else {
		response.Response(w, nil, http.StatusInternalServerError, errors.New("해당 VM을 하드 리부팅할 수 없습니다."))
	}
	return
}

// @Summary VM 스냅샷 생성
// @Description VM의 스냅샷을 생성합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param snapshotName body string true "스냅샷 이름"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /action/snapshot/{id} [post]
func snapshotVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	var CreateSnapshotBody struct {
		SnapshotName string
	}

	err := json.NewDecoder(r.Body).Decode(&CreateSnapshotBody)

	result, err := vm.Service.CreateSnapshot(uuid, id, CreateSnapshotBody.SnapshotName)
	if err != nil {
		if err == errors.New("NOT FOUND") {
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
		} else {
			response.Response(w, nil, http.StatusInternalServerError, err)
		}
		return
	}

	if result {
		response.Response(w, "OK", http.StatusOK, nil)
	} else {
		response.Response(w, nil, http.StatusInternalServerError, errors.New("해당 VM의 스냅샷을 생성할 수 없습니다."))
	}
	return
}

// @Summary VM 네트워크 리스트 조회
// @Description VM의 네트워크 리스트를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /resource/network [get]
func getNetworkList(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	rsp, err := vm.Service.ListNetworks(uuid)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)

}

// @Summary VM 네트워크 생성
// @Description VM의 네트워크를 생성합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   networkName body string true "네트워크 이름"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /resource/network [post]
func createNetwork(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	var networkCreateBody struct {
		networkName string
	}

	err := json.NewDecoder(r.Body).Decode(&networkCreateBody)
	if err != nil {
		response.Response(w, nil, http.StatusBadRequest, err)
		return
	}

	result, err := vm.Service.CreateNetwork(uuid, networkCreateBody.networkName)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, result, http.StatusOK, nil)
}

// @Summary VM 플레이버 리스트 조회
// @Description VM의 플레이버 리스트를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /resource/flavor [get]
func getFlavorList(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	rsp, err := vm.Service.ListFlavors(uuid)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)
}

// @Summary VM 키페어 생성
// @Description VM의 키페어를 생성합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   keypairName body string true "키페어 이름"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /resource/keypair [post]
func createKeypair(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	var keypairCreateBody struct {
		keypairName string
	}

	err := json.NewDecoder(r.Body).Decode(&keypairCreateBody)
	if err != nil {
		response.Response(w, nil, http.StatusBadRequest, err)
		return
	}

	result, err := vm.Service.CreateKeypair(uuid, keypairCreateBody.keypairName)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, result, http.StatusOK, nil)
}

// @Summary VM 키페어 리스트 조회
// @Description VM의 키페어 리스트를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /resource/keypair [get]
func getKeypairList(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	rsp, err := vm.Service.ListKeypairs(uuid)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)

}

// @Summary VM 보안그룹 리스트 조회
// @Description VM의 보안그룹 리스트를 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /resource/securitygroup [get]
func getSecurityGroupList(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("X-UUID")
	rsp, err := vm.Service.ListSecurityGroups(uuid)

	if err != nil {
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)

}

// @Summary VM VNC URL 조회
// @Description VM의 VNC URL을 조회합니다.
// @Tags vm
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "VM uuid"
// @Param X-UUID header string true "UUID"
// @Success 200 {object} response.CommonResponse
// @Router /action/vnc/{id} [get]
func getVmVnc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uuid := r.Header.Get("X-UUID")

	rsp, err := vm.Service.GetVnc(uuid, id)

	if err != nil {
		if err == errors.New("NOT FOUND") {
			response.Response(w, nil, http.StatusNotFound, errors.New("해당 VM이 없습니다."))
			return
		}
		response.Response(w, nil, http.StatusInternalServerError, err)
		return
	}

	response.Response(w, rsp, http.StatusOK, nil)
}
