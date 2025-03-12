package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common/handler"
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
	cash, err := h.cashService.Edit(r.Context(), r.PathValue("id"))
	if err != nil {
		h.SetError(w, err)

		return
	}

	w.Header().Add("Hx-Trigger-After-Swap", "balance.cash.edit")

	err = h.component.Edit(r.URL.Query(), cash, nil).Render(w)

	h.SetError(w, err)
}
