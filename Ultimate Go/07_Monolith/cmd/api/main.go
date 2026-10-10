package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Shawshank44/olx-api/internal/config"
	"github.com/Shawshank44/olx-api/internal/db"
	"github.com/Shawshank44/olx-api/internal/handlers"
	"github.com/Shawshank44/olx-api/internal/middlewares"
)

func main() {

	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("main.db.connect : %v", err)
	}
	defer db.Close()

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true, // not recommened for high traffic
		Level:     slog.LevelInfo,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	lh := handlers.NewListingHandler(db, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/delete/{id}", lh.Delete)

	handler := middlewares.RequestID(mux)
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
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
