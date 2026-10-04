package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Shawshank44/olx-api/internal/config"
	"github.com/Shawshank44/olx-api/internal/db"
	"github.com/Shawshank44/olx-api/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("main.db.connect : %v", err)
	}
	defer db.Close()

	lh := handlers.NewListingHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/delete/{id}", lh.Delete)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Println("Server running successfully.")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed : %v", err)
	}
}
