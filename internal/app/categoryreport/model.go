package categoryreport

import (
	"maps"
	"slices"

	"github.com/shopspring/decimal"
)

type Entity struct {
	CategoryName  string
	CategorySlug  string
	Sum           decimal.Decimal
	Supercategory int
}

type Entities []*Entity

func (e Entities) HasMoreThanOne() bool {
	return len(e) > 1
}

func (e Entities) Sum() decimal.Decimal {
	var sum decimal.Decimal

	for _, entity := range e {
		sum = sum.Add(entity.Sum)
	}

	return sum
}

type MappedEntities map[int]Entities

func (e MappedEntities) Keys() []int {
	keys := slices.Collect(maps.Keys(e))

	slices.Sort(keys)

	return keys
}
