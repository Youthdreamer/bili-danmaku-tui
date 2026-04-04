package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() string {
	godotenv.Load()
	return os.Getenv("BLIVE_COOKIE")
}
