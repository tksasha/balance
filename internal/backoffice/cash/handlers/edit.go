package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type EditHandler struct {
	*handler.Handler

	cashService cash.Service
	component   *component.Component
}

func NewEditHandler(
	cashService cash.Service,
) *EditHandler {
	return &EditHandler{
		Handler:     handler.New(),
		cashService: cashService,
		component:   component.New(),
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	params := params.New().SetCurrency(cash.Currency)

	err = h.component.Edit(params, cash).Render(w)

	h.SetError(w, err)
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.cashService.Edit(r.Context(), r.PathValue("id"))
}
