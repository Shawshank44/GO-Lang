package main

import (
	"blog_rest_api/internal/config"
	"blog_rest_api/internal/db"
	router "blog_rest_api/internal/handlers/Router"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	routers := router.MainRouter()

	fmt.Printf("Server successfully created on http://localhost%s", cfg.API_PORT)
	err = http.ListenAndServe(cfg.API_PORT, routers)
	if err != nil {
		log.Fatal(err.Error())
	}
}
