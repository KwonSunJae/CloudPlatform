package response

import (
	"encoding/json"
	"net/http"
)

type CommonResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

func Response(w http.ResponseWriter, data interface{}, status int, err error) {

	var res = CommonResponse{
		Data:   data,
		Status: status,
		Error:  err,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if status >= 500 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(200)
	}
	json.NewEncoder(w).Encode(res)
}
