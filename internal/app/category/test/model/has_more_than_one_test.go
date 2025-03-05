package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"gotest.tools/v3/assert"
)

func TestHasMoreThanOne(t *testing.T) {
	t.Run("when has more than one", func(t *testing.T) {
		categories := category.Categories{
			{Name: "Food"},
			{Name: "Beverages"},
			{Name: "Salary"},
		}

		assert.Assert(t, categories.HasMoreThanOne())
	})

	t.Run("when has only one", func(t *testing.T) {
		categories := category.Categories{
			{Name: "Food"},
		}

		assert.Assert(t, !categories.HasMoreThanOne())
	})

	t.Run("when empty", func(t *testing.T) {
		categories := category.Categories{}

		assert.Assert(t, !categories.HasMoreThanOne())
	})
}
