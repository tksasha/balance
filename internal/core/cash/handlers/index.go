package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	components "github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type IndexHandler struct {
	service cash.Service
}

func NewIndexHandler(service cash.Service) *IndexHandler {
	return &IndexHandler{
		service: service,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Index(cashes).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *IndexHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.service.List(r.Context())
}
