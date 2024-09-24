package decorators_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestDate(t *testing.T) {
	sbj := decorators.NewItemDecorator(
		&models.Item{
			Date: time.Date(2024, 9, 30, 0, 0, 0, 0, time.UTC),
		},
	).Date()

	exp := "2024-09-30"

	assert.Equal(t, sbj, exp)
}

func TestSum(t *testing.T) {
	sbj := decorators.NewItemDecorator(
		&models.Item{
			Sum: 123_456.78,
		},
	).Sum()

	exp := "123\u00A0456,78"

	assert.Equal(t, sbj, exp)
}

func TestCategoryName(t *testing.T) {
	sbj := decorators.NewItemDecorator(
		&models.Item{
			CategoryName: "drinks",
		},
	).CategoryName()

	exp := "drinks"

	assert.Equal(t, sbj, exp)
}

func TestDecription(t *testing.T) {
	sbj := decorators.NewItemDecorator(
		&models.Item{
			Description: "foods and beverages",
		},
	).Description()

	exp := "foods and beverages"

	assert.Equal(t, sbj, exp)
}
