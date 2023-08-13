package main

import (
	"net/http"
	"soms/controller/vm"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	err := vm.VmController(r)

	if err != nil {
		panic("서버 실행에 실패했습니다.")
	}

	http.ListenAndServe(":3000", r)
}
