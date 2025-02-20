package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/component"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type ListHandler struct {
	service       cash.Service
	cashComponent *component.Component
}

func NewListHandler(
	service cash.Service,
	cashComponent *component.Component,
) *ListHandler {
	return &ListHandler{
		service:       service,
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
	return h.service.List(r.Context())
}
