package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	database "github.com/hosnibounechada/go-api/internal/db"
)

func main() {

	connectionStr := "host=localhost port=5432 user=postgres password=password dbname=gindb sslmode=disable"

	db, err := database.InitDatabase(connectionStr)
	if err != nil {
		log.Fatalf("Error: %v", err)
		log.Println("Database initialization failed")
		return
	}

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		log.Fatalf("Error: %v", err)
		log.Println("Migration driver creation failed")
		return
	}

	migrationDir := "../migrations"

	sourceDriver, err := source.Open(migrationDir)
	if err != nil {
		log.Fatalf("Error: %v", err)
		log.Println("Migration source creation failed")
		return
	}

	m, err := migrate.NewWithInstance("file", sourceDriver, "postgres", driver)
	if err != nil {
		log.Fatalf("Error: %v", err)
		log.Println("Migration creation failed")
		return
	}

	// You can specify the desired version or use "migrate.Up" to apply all available migrations.
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("Error: %v", err)
		log.Println("Failed to retrieve migration version")
		return
	}

	log.Printf("Current database version: %v\n", version)
	if dirty {
		log.Println("The database is in a dirty state")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error: %v", err)
		log.Println("Migration failed")
	}
}
