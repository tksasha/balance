package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
	"github.com/tksasha/balance/internal/core/item/handlers"
)

func NewCreateItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.CreateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewCreateHandler(itemService, categoryService, itemsComponent)
}

func NewEditItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.EditHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewEditHandler(itemService, categoryService, itemsComponent)
}

func NewUpdateItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.UpdateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewUpdateHandler(itemService, categoryService, itemsComponent)
}

func NewListItemsHandler(
	t *testing.T,
	itemService item.Service,
) *handlers.ListHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewListHandler(itemService, itemsComponent)
}
