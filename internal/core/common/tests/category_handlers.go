package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/category/handlers"
	"github.com/tksasha/balance/internal/core/common"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
)

func NewListCategoriesHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.ListHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := common.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewListHandler(baseHandler, categoryService, categoryComponent)
}

func NewCreateCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.CreateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := common.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewCreateHandler(baseHandler, categoryService, categoryComponent)
}

func NewEditCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.EditHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := common.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	categoryComponent := components.NewCategoryComponent(baseComponent)

	return handlers.NewEditHandler(baseHandler, categoryService, categoryComponent)
}

func NewUpdateCategoryHandler(
	t *testing.T,
	categoryService category.Service,
) *handlers.UpdateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := common.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

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
