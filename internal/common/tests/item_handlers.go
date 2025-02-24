package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component"
	indexcomponents "github.com/tksasha/balance/internal/core/index/components"
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

	itemsComponent := components.NewItemsComponent(component.New())

	return handlers.NewCreateHandler(common.NewBaseHandler(), itemService, categoryService, itemsComponent)
}

func NewEditItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.EditHandler {
	t.Helper()

	itemsComponent := components.NewItemsComponent(component.New())

	return handlers.NewEditHandler(common.NewBaseHandler(), itemService, categoryService, itemsComponent)
}

func NewUpdateItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *handlers.UpdateHandler {
	t.Helper()

	itemsComponent := components.NewItemsComponent(component.New())

	return handlers.NewUpdateHandler(common.NewBaseHandler(), itemService, categoryService, itemsComponent)
}

func NewListItemsHandler(
	t *testing.T,
	itemService item.Service,
) *handlers.ListHandler {
	t.Helper()

	component := component.New()

	itemsComponent := components.NewItemsComponent(component)

	monthsComponent := indexcomponents.NewMonthsComponent(component)

	yearsComponent := indexcomponents.NewYearsComponent(component)

	return handlers.NewListHandler(
		common.NewBaseHandler(),
		itemService,
		itemsComponent,
		monthsComponent,
		yearsComponent,
	)
}

func NewDeleteItemHandler(t *testing.T, itemService item.Service) *handlers.DeleteHandler {
	t.Helper()

	return handlers.NewDeleteHandler(common.NewBaseHandler(), itemService)
}
