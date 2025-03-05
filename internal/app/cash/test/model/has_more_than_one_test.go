package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"gotest.tools/v3/assert"
)

func TestHasMoreThanOne(t *testing.T) {
	t.Run("when has more than one", func(t *testing.T) {
		cashes := cash.Cashes{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		}

		assert.Assert(t, cashes.HasMoreThanOne())
	})

	t.Run("when has only one", func(t *testing.T) {
		cashes := cash.Cashes{
			{ID: 1},
		}

		assert.Assert(t, !cashes.HasMoreThanOne())
	})

	t.Run("when empty", func(t *testing.T) {
		cashes := cash.Cashes{}

		assert.Assert(t, !cashes.HasMoreThanOne())
	})
}
