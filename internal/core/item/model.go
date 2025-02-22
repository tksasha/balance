package item

import (
	"database/sql"
	"time"

	"github.com/tksasha/balance/pkg/currencies"
)

type Item struct {
	ID           int
	Date         time.Time
	Formula      string
	Sum          float64
	CategoryID   int
	CategoryName sql.NullString
	Description  string
	Currency     currencies.Currency
}

type Items []*Item
