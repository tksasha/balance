package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/category/handlers"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
)

func NewListCategoriesHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.ListHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewListHandler(categoryService, categoryComponent)
}

func NewCreateCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.CreateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewCreateHandler(categoryService, categoryComponent)
}

func NewEditCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.EditHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewEditHandler(categoryService, categoryComponent)
}
