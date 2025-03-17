package cash

import (
	"maps"
	"slices"

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

func (c Cashes) HasMoreThanOne() bool {
	return len(c) > 1
}

func (c Cashes) Sum() float64 {
	var sum float64

	for _, cash := range c {
		sum += cash.Sum
	}

	return sum
}

type GroupedCashes map[int]Cashes

func (c GroupedCashes) Keys() []int {
	keys := slices.Collect(maps.Keys(c))

	slices.Sort(keys)

	return keys
}
