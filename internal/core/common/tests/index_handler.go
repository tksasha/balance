package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/component"
	"github.com/tksasha/balance/internal/core/index"
	"github.com/tksasha/balance/internal/core/index/components"
	"github.com/tksasha/balance/internal/core/index/handler"
)

func NewIndexPageHandler(
	t *testing.T,
	indexService index.Service,
	categoryService category.Service,
) *handler.Handler {
	t.Helper()

	component := component.New()

	monthsComonents := components.NewMonthsComponent(component)

	yearsComponent := components.NewYearsComponent(component)

	indexComponent := components.NewIndexComponent(component, monthsComonents, yearsComponent)

	return handler.New(common.NewBaseHandler(), indexService, categoryService, indexComponent)
}
