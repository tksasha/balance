package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/indexpage"
	"github.com/tksasha/balance/internal/core/indexpage/components"
	"github.com/tksasha/balance/internal/core/indexpage/handler"
)

func NewIndexPageHandler(
	t *testing.T,
	indexPageService indexpage.Service,
	categoryService category.Service,
) *handler.Handler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	monthsComonents := components.NewMonthsComponent(baseComponent)

	indexPageComponent := components.NewIndexPageComponent(baseComponent, monthsComonents)

	return handler.New(indexPageService, categoryService, indexPageComponent)
}
