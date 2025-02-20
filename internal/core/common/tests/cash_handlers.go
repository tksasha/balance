package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/cash/handlers"
	"github.com/tksasha/balance/internal/core/common"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
)

func NewCreateCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.CreateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := components.NewCashComponent(baseComponent)

	return handlers.NewCreateHandler(baseHandler, cashService, cashComponent)
}

func NewEditCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.EditHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := components.NewCashComponent(baseComponent)

	return handlers.NewEditHandler(baseHandler, cashService, cashComponent)
}

func NewListCashesHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.ListHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := components.NewCashComponent(baseComponent)

	return handlers.NewListHandler(baseHandler, cashService, cashComponent)
}

func NewNewCasheHandler(
	t *testing.T,
) *handlers.NewHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := components.NewCashComponent(baseComponent)

	return handlers.NewNewHandler(baseHandler, cashComponent)
}

func NewUpdateCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.UpdateHandler {
	t.Helper()

	baseHandler := common.NewBaseHandler()

	currentDateProvider := providers.NewTimeProvider()

	helpers := helpers.New(currentDateProvider)

	baseComponent := commoncomponents.NewBaseComponent(helpers)

	cashComponent := components.NewCashComponent(baseComponent)

	return handlers.NewUpdateHandler(baseHandler, cashService, cashComponent)
}

func NewDeleteCashHandler(
	t *testing.T,
	cashService cash.Service,
) *handlers.DeleteHandler {
	t.Helper()

	return handlers.NewDeleteHandler(common.NewBaseHandler(), cashService)
}
