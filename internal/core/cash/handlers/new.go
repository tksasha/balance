package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type NewHandler struct {
	cashComponent *components.CashComponent
}

func NewNewHandler(cashComponent *components.CashComponent) *NewHandler {
	return &NewHandler{
		cashComponent: cashComponent,
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	err := h.cashComponent.New(cash).Render(w)

	handlers.SetError(w, err)
}
