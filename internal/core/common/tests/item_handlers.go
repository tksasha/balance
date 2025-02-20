package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/index"
	indexcomponents "github.com/tksasha/balance/internal/core/index/components"
	indexhandlers "github.com/tksasha/balance/internal/core/index/handlers"
	"github.com/tksasha/balance/internal/core/item"
	itemcomponents "github.com/tksasha/balance/internal/core/item/components"
	itemhandlers "github.com/tksasha/balance/internal/core/item/handlers"
)

func NewIndexPageHandler(
	t *testing.T,
	indexService index.Service,
	categoryService category.Service,
) *indexhandlers.IndexHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	monthsComonents := indexcomponents.NewMonthsComponent(baseComponent)

	indexPageComponent := indexcomponents.NewIndexPageComponent(baseComponent, monthsComonents)

	return indexhandlers.NewIndexHandler(indexService, categoryService, indexPageComponent)
}

func NewCreateItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *itemhandlers.CreateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := itemcomponents.NewItemsComponent(baseComponent)

	return itemhandlers.NewCreateHandler(itemService, categoryService, itemsComponent)
}

func NewEditItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *itemhandlers.EditHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := itemcomponents.NewItemsComponent(baseComponent)

	return itemhandlers.NewEditHandler(itemService, categoryService, itemsComponent)
}

func NewUpdateItemHandler(
	t *testing.T,
	itemService item.Service,
	categoryService category.Service,
) *itemhandlers.UpdateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := itemcomponents.NewItemsComponent(baseComponent)

	return itemhandlers.NewUpdateHandler(itemService, categoryService, itemsComponent)
}

func NewIndexItemsHandler(
	t *testing.T,
	itemService item.Service,
) *itemhandlers.IndexHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	itemsComponent := itemcomponents.NewItemsComponent(baseComponent)

	return itemhandlers.NewIndexHandler(itemService, itemsComponent)
}
