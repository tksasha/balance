package cash

import (
	"database/sql"

	"github.com/tksasha/balance/pkg/currencies"
)

type Cash struct {
	ID            int
	Currency      currencies.Currency
	Formula       string
	Sum           float64
	Name          string
	Supercategory int
	Favorite      bool
	DeletedAt     sql.NullTime
}

type Cashes []*Cash
