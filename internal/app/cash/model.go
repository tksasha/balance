package cash

import (
	"github.com/tksasha/balance/internal/common/currency"
)

type Cash struct {
	ID            int
	Currency      currency.Currency
	Formula       string
	Sum           float64
	Name          string
	Supercategory int
}

type Cashes []*Cash
