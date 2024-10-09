package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tksasha/balance/internal/models"
)

type Config struct {
	Address    string
	Currencies models.Currencies
}

func New() *Config {
	address := os.Getenv("ADDRESS")
	if address == "" {
		panic("ADDRESS must be set")
	}

	return &Config{
		Address: address,
		Currencies: models.Currencies{
			models.Currency{
				ID:   0,
				Name: "uah",
			},
			models.Currency{
				ID:   1,
				Name: "usd",
			},
			models.Currency{
				ID:   3, //nolint:mnd
				Name: "eur",
			},
		},
	}
}
