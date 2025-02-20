package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
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

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewCreateHandler(baseHandler, itemService, categoryService, itemsComponent)
}

func NewEditItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.EditHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewEditHandler(baseHandler, itemService, categoryService, itemsComponent)
}

func NewUpdateItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.UpdateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewUpdateHandler(baseHandler, itemService, categoryService, itemsComponent)
}

func NewListItemsHandler(
	t *testing.T,
	itemService item.Service,
) *handlers.ListHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	baseComponent := common.NewBaseComponent()

	itemsComponent := components.NewItemsComponent(baseComponent)

	return handlers.NewListHandler(baseHandler, itemService, itemsComponent)
}

func NewDeleteItemHandler(
	t *testing.T,
	itemService item.Service,
) *handlers.DeleteHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	return handlers.NewDeleteHandler(baseHandler, itemService)
}
