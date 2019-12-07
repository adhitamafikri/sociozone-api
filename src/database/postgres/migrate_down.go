package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://src/database/postgres/migrations",
		"postgres://mac@localhost:5432/scz_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
}
