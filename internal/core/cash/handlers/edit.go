package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common"
)

type EditHandler struct {
	*common.BaseHandler

	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewEditHandler(
	baseHandler *common.BaseHandler,
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *EditHandler {
	return &EditHandler{
		BaseHandler:   baseHandler,
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
