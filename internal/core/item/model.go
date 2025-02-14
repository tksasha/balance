package item

import (
	"time"

	"github.com/tksasha/balance/pkg/currencies"
)

type Item struct {
	ID           int
	Date         time.Time
	Formula      string
	Sum          float64
	CategoryID   int
	CategoryName string
	Description  string
	Currency     currencies.Currency
}

type Items []*Item
