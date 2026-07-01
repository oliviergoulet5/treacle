package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := getEnv("PORT", "8080")

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

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
