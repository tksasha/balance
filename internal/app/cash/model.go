package cash

import (
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
}

type Cashes []*Cash

func (c Cashes) HasMoreThanOne() bool {
	return len(c) > 1
}

func (c Cashes) Sum() decimal.Decimal {
	var sum decimal.Decimal

	for _, cash := range c {
		sum = sum.Add(cash.Sum)
	}

	return sum
}
