package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"gotest.tools/v3/assert"
)

func TestIncome(t *testing.T) {
	categories := category.Categories{
		{ID: 1, Income: false},
		{ID: 2, Income: false},
		{ID: 3, Income: true},
	}

	for _, category := range categories.Income() {
		assert.Assert(t, category.Income)
	}
}
