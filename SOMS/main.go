package main

import (
	"net/http"
	"soms/controller/container/deployment"
	"soms/controller/container/replicaset"
	service "soms/controller/container/service"
	"soms/controller/mockup"
	"soms/controller/user"
	"soms/controller/vm"

	_ "soms/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title SOMS API
// @version 1.0
// @description Cloud Platform API Server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func main() {
	r := mux.NewRouter()

	envLoad()

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
	err4 := replicaset.ReplicasetController(r)
	if err4 != nil {
		panic("replicaset 실행에 실패했습니다.")
	}
	err5 := user.UserController(r)
	if err5 != nil {
		panic("user 실행에 실패했습니다.")
	}
	err6 := mockup.MockupController(r)
	if err6 != nil {
		panic("mockup 실행에 실패했습니다.")
	}
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":3000", corsMiddleware(r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic("env file error")
	}
}
