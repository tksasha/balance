package model_test

import (
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"gotest.tools/v3/assert"
)

func TestKeys(t *testing.T) {
	categories := category.GroupedCategories{
		1: {},
		2: {},
		0: {},
	}

	actual := categories.Keys()

	expected := []int{0, 1, 2}

	assert.Assert(t, slices.Equal(actual, expected))
}
