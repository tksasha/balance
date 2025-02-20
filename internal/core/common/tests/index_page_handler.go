package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/helpers"
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

	baseHandler := common.NewBaseHandler()

	currentDateProvider := common.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := common.NewBaseComponent(helpers)

	monthsComonents := components.NewMonthsComponent(baseComponent)

	indexPageComponent := components.NewIndexPageComponent(baseComponent, monthsComonents)

	return handler.New(baseHandler, indexPageService, categoryService, indexPageComponent)
}
