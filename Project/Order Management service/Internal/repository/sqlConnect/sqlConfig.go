package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	host := os.Getenv("HOST")

	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, dbPort, dbName)

	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}

	log.Println("Connection established with", dbName, "successfully")

	return db, nil
}
