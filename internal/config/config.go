package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Address string
}

func New() *Config {
	address := os.Getenv("ADDRESS")
	if address == "" {
		panic("ADDRESS must be set")
	}

	return &Config{
		Address: address,
	}
}
