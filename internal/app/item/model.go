package item

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/tksasha/balance/internal/common/currency"
)

type Item struct {
	ID           int
	Date         time.Time
	Formula      string
	Sum          decimal.Decimal
	CategoryID   int
	CategoryName string
	CategorySlug string
	Description  string
	Currency     currency.Currency
}

type Items []*Item
