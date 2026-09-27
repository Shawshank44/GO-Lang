package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Shawshank44/olx-api/internal/config"
)

func main() {

	cfg := config.MustLoad()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status" : "ALL OK"}`))
	})

	srv := http.Server{
		Addr:         cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Println("Server running successfully.")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed : %v", err)
	}
}
