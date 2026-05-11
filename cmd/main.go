package main

import (
	"net/http"

	"github.com/omatheusq/rinha-de-backend-2026/internal/fraud"
)

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ready", readyHandler)
	mux.HandleFunc("/fraud-score", fraud.FraudScoreHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}