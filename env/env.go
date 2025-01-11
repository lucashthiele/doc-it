package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return err
}

func (ek EnvKey) GetValue() string {
	return os.Getenv(string(ek))
}

const (
	GitHubClientID     EnvKey = "GH_CLIENT_ID"
	GitHubClientSecret EnvKey = "GH_CLIENT_SECRET"
)
