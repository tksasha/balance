package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/category/handlers"
	"github.com/tksasha/balance/internal/core/common"
)

func NewListCategoriesHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.ListHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewListHandler(baseHandler, categoryService, categoryComponent)
}

func NewCreateCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.CreateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewCreateHandler(baseHandler, categoryService, categoryComponent)
}

func NewEditCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.EditHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewEditHandler(baseHandler, categoryService, categoryComponent)
}

func NewUpdateCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.UpdateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewUpdateHandler(baseHandler, categoryService, categoryComponent)
}

func NewCategoryDeleteHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.DeleteHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	return handlers.NewDeleteHandler(baseHandler, categoryService)
}
