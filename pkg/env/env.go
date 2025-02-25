package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

func Load() error {
	return godotenv.Load()
}

func (ek EnvKey) GetValue() (string, error) {
	envKey := string(ek)
	env := os.Getenv(envKey)
	if env == "" {
		return "", fmt.Errorf("%s env key not found", envKey)
	}
	return env, nil
}

const (
	GitHubClientID     EnvKey = "GH_CLIENT_ID"
	GitHubClientSecret EnvKey = "GH_CLIENT_SECRET"
)
