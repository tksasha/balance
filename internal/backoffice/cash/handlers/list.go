package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type ListHandler struct {
	*handler.Handler

	cashService cash.Service
	component   *component.Component
}

func NewListHandler(
	cashService cash.Service,
) *ListHandler {
	return &ListHandler{
		Handler:     handler.New(),
		cashService: cashService,
		component:   component.New(),
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.List(cashes).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.cashService.List(r.Context())
}
