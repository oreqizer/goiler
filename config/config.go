package config

import (
	"errors"
	"flag"
	"github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	// Port is the port to run the app on.
	Port string
	// DB is a url of our database.
	DB string
	// Firebase is Firebase's file path
	Firebase string
}

const firebase = "firebase-private-key.json"

func init() {
	_ = godotenv.Load()

	if dsn := os.Getenv("SENTRY"); dsn != "" {
		if err := raven.SetDSN(dsn); err != nil {
			log.Fatal(err)
		}
	}

	if fb := os.Getenv("FIREBASE_KEY_FILE"); fb != "" {
		if err := ioutil.WriteFile(firebase, []byte(fb), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func New() (*Config, error) {
	config := &Config{
		Port:     "8081",
		Firebase: firebase,
	}

	if db := flag.String("db", os.Getenv("DATABASE_URL"), "Postgres DB URL"); db != nil {
		config.DB = *db
	} else {
		return nil, errors.New("the DATABASE_URL environment variable is required")
	}

	if port := flag.String("port", os.Getenv("PORT"), "Server port"); port != nil {
		config.Port = *port
	}

	return config, nil
}
