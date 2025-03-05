package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"gotest.tools/v3/assert"
)

func TestSum(t *testing.T) {
	entities := categoryreport.Entities{
		{Sum: 11.11},
		{Sum: 22.22},
		{Sum: 33.33},
	}

	assert.Equal(t, entities.Sum(), 66.66)
}
