package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type ListHandler struct {
	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewListHandler(
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *ListHandler {
	return &ListHandler{
		cashService:   cashService,
		cashComponent: cashComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = h.cashComponent.List(cashes).Render(w)

	handlers.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.cashService.List(r.Context())
}
