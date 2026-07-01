package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := getEnv("PORT", "8080")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /request", requestHandler)

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Println("treacle running on :" + port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Get an environment variable. If the environment variable is not set, returns
// a fallback.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

type ExecuteRequest struct {
	Method string `json:"method"`
	URL string `json:"url"`
}

type ExecuteRequestResponse struct {
	StatusCode int `json:"statusCode"`
	Body string `json:"body"`
}

// Request handler for POST /request
func requestHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	var executeRequest ExecuteRequest
	err := json.NewDecoder(r.Body).Decode(&executeRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// TODO: Sanitize executeRequest.Method
	req, err := http.NewRequest(executeRequest.Method, executeRequest.URL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	executeRequestRes := ExecuteRequestResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(executeRequestRes); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
