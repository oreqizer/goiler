package config

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	// Port is the port to run the app on
	Port string
	// Cors specifies cross-origin resource sharing policy
	Cors string
	// DB is a url of our database
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
		Cors:     "*",
		Firebase: firebase,
	}

	if db := os.Getenv("DATABASE_URL"); db != "" {
		config.DB = db
	} else {
		return nil, errors.New("the DATABASE_URL environment variable is required")
	}

	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	}

	if cors := os.Getenv("CORS"); cors != "" {
		config.Cors = cors
	}

	return config, nil
}
