package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/index"
	"github.com/tksasha/balance/internal/core/index/components"
	"github.com/tksasha/balance/internal/core/index/handlers"
)

func NewIndexPageHandler(
	t *testing.T,
	indexService index.Service,
	categoryService category.Service,
) *handlers.IndexHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	monthsComonents := components.NewMonthsComponent(baseComponent)

	indexPageComponent := components.NewIndexPageComponent(baseComponent, monthsComonents)

	return handlers.NewIndexHandler(indexService, categoryService, indexPageComponent)
}
