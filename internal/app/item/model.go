package item

import (
	"time"

	"github.com/tksasha/balance/internal/common/currency"
)

type Item struct {
	ID           int
	Date         time.Time
	Formula      string
	Sum          float64
	CategoryID   int
	CategoryName string
	Description  string
	Currency     currency.Currency
}

type Items []*Item
