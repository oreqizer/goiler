package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/oreqizer/goiler/cmd/migrate/seed"
	"log"
	"os"
)

const files = "file://migrations"

/*
CLI to reset (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

migrate -path=migrations -database=<DB> drop
*/

func main() {
	_ = godotenv.Load()

	url := os.Getenv("DATABASE_URL")
	staging := os.Getenv("STAGING") != ""

	m, err := migrate.New(files, url)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if staging {
		reset(m, url)
	} else {
		upgrade(m)
	}
}

func upgrade(m *migrate.Migrate) {
	v, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			log.Println("No migration present")
		} else {
			log.Fatal(err)
		}
	}

	if dirty {
		log.Fatalf("Current DIRTY version: %d, please fix issues\n", v)
	}

	log.Printf("Current version: %d\n", v)

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No change")
			return
		}

		log.Fatal(err)
	}

	v, _, err = m.Version()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New version: %d\n", v)
}

func reset(m *migrate.Migrate, url string) {
	log.Println("Migrating down")
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrating up")
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding")
	if err := seed.Seed(url); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
