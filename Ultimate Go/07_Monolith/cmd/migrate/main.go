package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shawshank44/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustLoad()

	if len(os.Args) < 2 {
		log.Fatal("usage : make migrate <up | down>")
	}

	m, err := migrate.New("file://migrations", cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("migrate.new : %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command : %s", os.Args[1])
	}

	fmt.Println("Running migrations")
}
