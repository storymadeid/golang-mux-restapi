package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, paylaod interface{}) {
	response, _ := json.Marshal(paylaod)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, statusCode int, err string) {
	ResponseJson(w, statusCode, map[string]string{"message": err})
}
