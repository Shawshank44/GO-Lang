package db

import (
	"blog_rest_api/internal/config"
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connecting to the database.....")

	connection_URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB_USER, cfg.DB_PASSWORD, cfg.HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := sql.Open("mysql", connection_URL)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected and Established with", cfg.DB_NAME, "successfully")
	return db, nil
}
