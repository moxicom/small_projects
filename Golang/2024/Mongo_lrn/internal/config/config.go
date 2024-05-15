package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Uri string
}

func Get() Config {
	godotenv.Load()

	var cfg Config

	cfg.Uri = os.Getenv("MONGODB_URI")
	if cfg.Uri == "" {
		panic("MONGODB_URI is required")
	}

	return cfg
}
