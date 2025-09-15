package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int       `json:"code"`
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

func WriteJSONError(w http.ResponseWriter, statusCode int, errorType ErrorType, message ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var msg string
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	} else {
		msg = string(errorType)
	}

	response := ErrorResponse{
		Code:    statusCode,
		Type:    errorType,
		Message: msg,
	}

	json.NewEncoder(w).Encode(response)
}
