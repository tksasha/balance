package cash

import (
	"database/sql"

	"github.com/shopspring/decimal"
	"github.com/tksasha/balance/internal/common/currency"
)

type Cash struct {
	ID            int
	Currency      currency.Currency
	Formula       string
	Sum           decimal.Decimal
	Name          string
	Supercategory int
	Favorite      bool
	DeletedAt     sql.NullTime
}

type Cashes []*Cash
