package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/components"
	"github.com/tksasha/balance/internal/app/cash/handlers"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component"
)

func NewCreateCashHandler(t *testing.T, cashService cash.Service) *handlers.CreateHandler {
	t.Helper()

	cashComponent := components.NewCashComponent(component.New())

	return handlers.NewCreateHandler(common.NewBaseHandler(), cashService, cashComponent)
}

func NewEditCashHandler(t *testing.T, cashService cash.Service) *handlers.EditHandler {
	t.Helper()

	cashComponent := components.NewCashComponent(component.New())

	return handlers.NewEditHandler(common.NewBaseHandler(), cashService, cashComponent)
}

func NewListCashesHandler(t *testing.T, cashService cash.Service) *handlers.ListHandler {
	t.Helper()

	cashComponent := components.NewCashComponent(component.New())

	return handlers.NewListHandler(common.NewBaseHandler(), cashService, cashComponent)
}

func NewNewCasheHandler(t *testing.T) *handlers.NewHandler {
	t.Helper()

	cashComponent := components.NewCashComponent(component.New())

	return handlers.NewNewHandler(common.NewBaseHandler(), cashComponent)
}

func NewUpdateCashHandler(t *testing.T, cashService cash.Service) *handlers.UpdateHandler {
	t.Helper()

	cashComponent := components.NewCashComponent(component.New())

	return handlers.NewUpdateHandler(common.NewBaseHandler(), cashService, cashComponent)
}

func NewDeleteCashHandler(t *testing.T, cashService cash.Service) *handlers.DeleteHandler {
	t.Helper()

	return handlers.NewDeleteHandler(common.NewBaseHandler(), cashService)
}
