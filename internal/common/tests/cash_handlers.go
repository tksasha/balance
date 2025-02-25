package tests

import (
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/components"
	"github.com/tksasha/balance/internal/app/cash/handlers"
)

func NewEditCashHandler(t *testing.T, cashService cash.Service) *handlers.EditHandler {
	t.Helper()

	cashComponent := components.NewCashComponent()

	return handlers.NewEditHandler(cashService, cashComponent)
}

func NewListCashesHandler(t *testing.T, cashService cash.Service) *handlers.ListHandler {
	t.Helper()

	cashComponent := components.NewCashComponent()

	return handlers.NewListHandler(cashService, cashComponent)
}

func NewNewCasheHandler(t *testing.T) *handlers.NewHandler {
	t.Helper()

	cashComponent := components.NewCashComponent()

	return handlers.NewNewHandler(cashComponent)
}

func NewUpdateCashHandler(t *testing.T, cashService cash.Service) *handlers.UpdateHandler {
	t.Helper()

	cashComponent := components.NewCashComponent()

	return handlers.NewUpdateHandler(cashService, cashComponent)
}

func NewDeleteCashHandler(t *testing.T, cashService cash.Service) *handlers.DeleteHandler {
	t.Helper()

	return handlers.NewDeleteHandler(cashService)
}
