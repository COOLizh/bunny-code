// Package handlers for API
package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheck view
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("ok")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// AuthCheck needed to test authentication
func AuthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("secure endpoint")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
