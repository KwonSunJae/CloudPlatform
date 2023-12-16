package main

import (
	"net/http"
	"soms/controller/container/deployment"
	"soms/controller/container/service"
	"soms/controller/vm"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	err := vm.VmController(r)
	if err != nil {
		panic("vm 서버 실행에 실패했습니다.")
	}
	err2 := deployment.DeploymentController(r)
	if err2 != nil {
		panic("deployment 실행에 실패했습니다.")
	}
	err3 := service.ServiceController(r)

	if err3 != nil {
		panic("service 실행에 실패했습니다.")
	}

	http.ListenAndServe(":3000", r)
}
