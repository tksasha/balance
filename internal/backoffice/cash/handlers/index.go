package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*handler.Handler

	cashService cash.Service
	component   *component.Component
}

func NewIndexHandler(
	cashService cash.Service,
) *IndexHandler {
	return &IndexHandler{
		Handler:     handler.New(),
		cashService: cashService,
		component:   component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.cashService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	w.Header().Add("Hx-Trigger-After-Swap", "backoffice.cashes.shown")

	err = h.component.List(cashes).Render(w)

	h.SetError(w, err)
}
