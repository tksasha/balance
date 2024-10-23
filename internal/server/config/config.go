package config

import (
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Address           string
	ReadHeaderTimeout time.Duration
}

func New() *Config {
	address := os.Getenv("ADDRESS")
	if address == "" {
		panic("ADDRESS must be set")
	}

	return &Config{
		Address:           address,
		ReadHeaderTimeout: 1 * time.Second,
	}
}
