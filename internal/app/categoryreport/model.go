package categoryreport

import (
	"maps"
	"slices"
)

type Entity struct {
	CategoryName  string
	CategorySlug  string
	Sum           float64
	Supercategory int
}

type Entities []*Entity

func (e Entities) HasMoreThanOne() bool {
	return len(e) > 1
}

func (e Entities) Sum() float64 {
	var sum float64

	for _, entity := range e {
		sum += entity.Sum
	}

	return sum
}

type MappedEntities map[int]Entities

func (e MappedEntities) Keys() []int {
	keys := slices.Collect(maps.Keys(e))

	slices.Sort(keys)

	return keys
}
