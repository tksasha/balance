package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/components"
	"github.com/tksasha/balance/internal/common/handler"
)

type EditHandler struct {
	*handler.Handler

	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewEditHandler(
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *EditHandler {
	return &EditHandler{
		Handler:       handler.New(),
		cashService:   cashService,
		cashComponent: cashComponent,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.cashComponent.Edit(cash).Render(w)

	h.SetError(w, err)
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.cashService.Edit(r.Context(), r.PathValue("id"))
}
