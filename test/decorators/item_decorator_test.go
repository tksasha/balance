package decorators_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestDate(t *testing.T) {
	t.Run("when date is zero it should return an empty string", func(t *testing.T) {
		decorator := decorators.NewItemDecorator(
			&models.Item{
				Date: time.Time{},
			},
		)

		assert.Equal(t, decorator.Date, "")
	})

	t.Run("when date is not zero it should return formatted date", func(t *testing.T) {
		decorator := decorators.NewItemDecorator(
			&models.Item{
				Date: time.Date(2024, 10, 9, 10, 8, 53, 0, time.UTC),
			},
		)

		assert.Equal(t, decorator.Date, "2024-10-09")
	})
}

func TestFormula(t *testing.T) {
	decorator := decorators.NewItemDecorator(
		&models.Item{
			Formula: "4269.42+6942.69",
		},
	)

	assert.Equal(t, decorator.Formula, "4269.42+6942.69")
}

func TestSum(t *testing.T) {
	decorator := decorators.NewItemDecorator(
		&models.Item{
			Sum: 123456.78,
		},
	)

	assert.Equal(t, decorator.Sum, "123\u00A0456,78")
}

func TestCategoryName(t *testing.T) {
	decorator := decorators.NewItemDecorator(
		&models.Item{
			CategoryName: "Food",
		},
	)

	assert.Equal(t, decorator.CategoryName, "Food")
}

func TestDescription(t *testing.T) {
	decorator := decorators.NewItemDecorator(
		&models.Item{
			Description: "Just Simple Description",
		},
	)

	assert.Equal(t, decorator.Description, "Just Simple Description")
}
