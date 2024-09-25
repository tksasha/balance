package model_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestGetDate(t *testing.T) {
	item := models.Item{
		Date: time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
	}

	sbj := item.GetDate()

	exp := "2024-09-30"

	assert.Equal(t, sbj, exp)
}

func TestGetSum(t *testing.T) {
	item := models.Item{
		Sum: 123_456.78,
	}

	sbj := item.GetSum()

	exp := "123\u00A0456,78"

	assert.Equal(t, sbj, exp)
}
