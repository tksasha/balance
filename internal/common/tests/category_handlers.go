package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/components"
	"github.com/tksasha/balance/internal/app/category/handlers"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component"
)

func NewListCategoriesHandler(t *testing.T, categoryService category.Service) *handlers.ListHandler {
	t.Helper()

	categoryComponent := components.NewCategoryComponent(component.New())

	return handlers.NewListHandler(common.NewBaseHandler(), categoryService, categoryComponent)
}

func NewEditCategoryHandler(t *testing.T, categoryService category.Service) *handlers.EditHandler {
	t.Helper()

	categoryComponent := components.NewCategoryComponent(component.New())

	return handlers.NewEditHandler(common.NewBaseHandler(), categoryService, categoryComponent)
}

func NewUpdateCategoryHandler(t *testing.T, categoryService category.Service) *handlers.UpdateHandler {
	t.Helper()

	categoryComponent := components.NewCategoryComponent(component.New())

	return handlers.NewUpdateHandler(common.NewBaseHandler(), categoryService, categoryComponent)
}
