package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/app/index"
	"github.com/tksasha/balance/internal/app/index/components"
	"github.com/tksasha/balance/internal/app/index/handler"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component"
)

func NewIndexPageHandler(
	t *testing.T,
	indexService index.Service,
	categoryService index.CategoryService,
) *handler.Handler {
	t.Helper()

	component := component.New()

	monthsComonents := components.NewMonthsComponent(component)

	yearsComponent := components.NewYearsComponent(component)

	indexComponent := components.NewIndexComponent(component, monthsComonents, yearsComponent)

	return handler.New(common.NewBaseHandler(), indexService, categoryService, indexComponent)
}
