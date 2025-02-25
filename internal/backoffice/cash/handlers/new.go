package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/components"
	"github.com/tksasha/balance/internal/common/handler"
)

type NewHandler struct {
	*handler.Handler

	cashComponent *components.CashComponent
}

func NewNewHandler(
	cashComponent *components.CashComponent,
) *NewHandler {
	return &NewHandler{
		Handler:       handler.New(),
		cashComponent: cashComponent,
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	err := h.cashComponent.New(cash).Render(w)

	h.SetError(w, err)
}
