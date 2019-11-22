package handler

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, i interface{}, statusCode int) {
	jsonBytes, err := json.Marshal(i)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}

func WriteEmptyResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Length", "0")
	w.WriteHeader(statusCode)
}

func WriteJSONErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	var jsonError ErrorResponse
	jsonError.Message = message
	jsonBytes, _ := json.Marshal(jsonError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
