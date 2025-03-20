package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
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

	h.ok(w, cash, params.New(r.URL.Query()))
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.cashService.Edit(r.Context(), r.PathValue("id"))
}

func (h *EditHandler) ok(w http.ResponseWriter, cash *cash.Cash, params params.Params) {
	w.Header().Add("Hx-Trigger-After-Swap", "balance.cash.edit")

	if err := h.component.Edit(params, cash, nil).Render(w); err != nil {
		h.SetError(w, err)
	}
}
