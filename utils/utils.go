package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}


func WriteError(w http.ResponseWriter, message string, status int) {
    response := map[string]any{
        "success": false,
        "message": message,
    }
    JSONResponse(w, response, status)
}
