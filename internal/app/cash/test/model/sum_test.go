package model_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/tksasha/balance/internal/app/cash"
	"gotest.tools/v3/assert"
)

func TestSum(t *testing.T) {
	cashes := cash.Cashes{
		{Sum: decimal.NewFromFloat(11.11)},
		{Sum: decimal.NewFromFloat(22.22)},
		{Sum: decimal.NewFromFloat(33.33)},
	}

	sum := cashes.Sum()

	assert.Assert(t, decimal.NewFromFloat(66.66).Equal(sum))
}
