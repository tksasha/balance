package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/component"
	"github.com/tksasha/balance/internal/core/cash/handlers"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
)

func NewCreateCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.CreateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := component.New(baseComponent)

	return handlers.NewCreateHandler(cashService, cashComponent)
}

func NewEditCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.EditHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := component.New(baseComponent)

	return handlers.NewEditHandler(cashService, cashComponent)
}

func NewListCashesHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.IndexHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := component.New(baseComponent)

	return handlers.NewIndexHandler(cashService, cashComponent)
}

func NewNewCasheHandler(
	t *testing.T,
) *handlers.NewHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := component.New(baseComponent)

	return handlers.NewNewHandler(cashComponent)
}

func NewUpdateCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.UpdateHandler {
	t.Helper()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := component.New(baseComponent)

	return handlers.NewUpdateHandler(cashService, cashComponent)
}
