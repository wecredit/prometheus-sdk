package utils

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := JsonResponse{
		Status:  http.StatusText(status),
		Message: message,
	}
	json.NewEncoder(w).Encode(resp)
}
