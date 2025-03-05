package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"gotest.tools/v3/assert"
)

func TestExpense(t *testing.T) {
	categories := category.Categories{
		{ID: 1, Income: true},
		{ID: 2, Income: false},
		{ID: 3, Income: false},
	}

	for _, category := range categories.Expense() {
		assert.Assert(t, !category.Income)
	}
}
