package model_test

import (
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"gotest.tools/v3/assert"
)

func TestKeys(t *testing.T) {
	entities := categoryreport.MappedEntities{
		1: {},
		2: {},
		0: {},
	}

	actual := entities.Keys()

	expected := []int{0, 1, 2}

	assert.Assert(t, slices.Equal(actual, expected))
}
