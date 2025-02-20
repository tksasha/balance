package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type ListHandler struct {
	*handlers.BaseHandler

	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewListHandler(
	baseHandler *handlers.BaseHandler,
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *ListHandler {
	return &ListHandler{
		BaseHandler:   baseHandler,
		cashService:   cashService,
		cashComponent: cashComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.cashComponent.List(cashes).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.cashService.List(r.Context())
}
