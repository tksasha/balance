package decorators_test

import (
	"testing"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestID(t *testing.T) {
	category := &models.Category{ID: 9931}

	decorator := decorators.NewCategoryDecorator(category)

	assert.Equal(t, decorator.ID, "9931")
}

func TestName(t *testing.T) {
	category := &models.Category{
		Name: "Category Number One",
	}

	decorator := decorators.NewCategoryDecorator(category)

	assert.Equal(t, decorator.Name, "Category Number One")
}
