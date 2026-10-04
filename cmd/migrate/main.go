package main

import (
	"errors"
	"log"
	"os"

	"github.com/feerdim/boilerplate-golang/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("ERROR load env file : %v", err)
		}
	}

	dbx, _, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("ERROR database connection setup : %v", err)
	}

	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(os.Getenv("MIGRATION_PATH"), "postgresql", driver)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	direction := "up"
	if len(os.Args) > 1 {
		direction = os.Args[1]
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migration up failed: %v", err)
		}

		log.Println("migration up successfully applied")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migration down failed: %v", err)
		}

		log.Println("migration down successfully applied")
	default:
		log.Fatalf("unknown direction")
	}

	if err := dbx.Close(); err != nil {
		log.Printf("ERROR close database connection : %s", err.Error())
		return
	}
}
