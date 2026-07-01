package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oliviergoulet5/treacle/internal/httpclient"
	"github.com/oliviergoulet5/treacle/internal/models"
)

// Handler serves the Treacle HTTP API.
type Handler struct{}

// Health responds with the service health status.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// Request executes an outbound HTTP request and returns the response.
func (h *Handler) Request(w http.ResponseWriter, r *http.Request) {
	var req models.ExecuteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := httpclient.Execute(req)
	if err != nil {
		log.Printf("Failed %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
