package integration

import (
	"encoding/json"
	"net/http"
)

type ResponseAPI struct {
	Message string `json:"menssage"`
	Status  string `json:"status"`
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	response := ResponseAPI{
		Message: "Hello World!",
		Status:  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
