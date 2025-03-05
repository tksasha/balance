package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"gotest.tools/v3/assert"
)

func TestHasMoreThanOne(t *testing.T) {
	t.Run("when has more than one", func(t *testing.T) {
		entities := categoryreport.Entities{
			{CategoryName: "Food"},
			{CategoryName: "Beverages"},
			{CategoryName: "Salary"},
		}

		assert.Assert(t, entities.HasMoreThanOne())
	})

	t.Run("when has only one", func(t *testing.T) {
		entities := categoryreport.Entities{
			{CategoryName: "Food"},
		}

		assert.Assert(t, !entities.HasMoreThanOne())
	})

	t.Run("when empty", func(t *testing.T) {
		entities := categoryreport.Entities{}

		assert.Assert(t, !entities.HasMoreThanOne())
	})
}
