package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"gotest.tools/v3/assert"
)

func TestSum(t *testing.T) {
	cashes := cash.Cashes{
		{Sum: 11.11},
		{Sum: 22.22},
		{Sum: 33.33},
	}

	assert.Equal(t, cashes.Sum(), 66.66)
}
