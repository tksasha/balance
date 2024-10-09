package decorators_test

import (
	"testing"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNewCategoriesDecorator(t *testing.T) {
	income1 := decorators.NewCategoryDecorator(&models.Category{Income: true})
	income2 := decorators.NewCategoryDecorator(&models.Category{Income: true})

	expense1 := decorators.NewCategoryDecorator(&models.Category{Income: false})
	expense2 := decorators.NewCategoryDecorator(&models.Category{Income: false})

	categories := []*models.Category{
		income1.Category,
		income2.Category,
		expense1.Category,
		expense2.Category,
	}

	decorated := decorators.NewCategoriesDecorator(categories)

	assert.Assert(t, is.Contains(decorated.Income, income1))
	assert.Assert(t, is.Contains(decorated.Income, income2))

	assert.Assert(t, is.Contains(decorated.Expense, expense1))
	assert.Assert(t, is.Contains(decorated.Expense, expense2))
}
