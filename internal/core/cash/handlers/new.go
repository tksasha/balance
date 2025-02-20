package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type NewHandler struct {
	*handlers.BaseHandler

	cashComponent *components.CashComponent
}

func NewNewHandler(
	baseHandler *handlers.BaseHandler,
	cashComponent *components.CashComponent,
) *NewHandler {
	return &NewHandler{
		BaseHandler:   baseHandler,
		cashComponent: cashComponent,
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	err := h.cashComponent.New(cash).Render(w)

	h.SetError(w, err)
}
