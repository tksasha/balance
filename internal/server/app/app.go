package app

import (
	"database/sql"

	"github.com/tksasha/balance/internal/models"
)

type App struct {
	DB              *sql.DB
	Currencies      models.Currencies
	DefaultCurrency *models.Currency
}

func New(db *sql.DB, currencies models.Currencies, defaultCurrency *models.Currency) *App {
	return &App{
		DB:              db,
		Currencies:      currencies,
		DefaultCurrency: defaultCurrency,
	}
}
