package apiHandlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /api/health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	response := HealthResponse{
		Status: "ok",
	}
	
	json.NewEncoder(w).Encode(response)
}
