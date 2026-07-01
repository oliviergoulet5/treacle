package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/oliviergoulet5/treacle/internal/api"
)

func main() {
	port := getEnv("PORT", "8080")

	mux := http.NewServeMux()

	h := &api.Handler{}
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("POST /request", h.Request)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("treacle running on :" + port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
