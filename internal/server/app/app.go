package app

import (
	"database/sql"

	"github.com/tksasha/balance/internal/models"
)

type App struct {
	DB         *sql.DB
	Currencies models.Currencies
	Currency   models.Currency
}

func New(db *sql.DB) *App {
	return &App{
		DB: db,
		Currencies: models.Currencies{
			"UAH": {
				ID:   0,
				Code: "UAH",
			},
			"USD": {
				ID:   1,
				Code: "USD",
			},
			"EUR": {
				ID:   3, //nolint:mnd
				Code: "EUR",
			},
		},
		Currency: models.Currency{ID: 0, Code: "UAH"}, // default currency
	}
}
