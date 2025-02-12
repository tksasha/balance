package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/cash"
	components "github.com/tksasha/balance/internal/cash/components"
	"github.com/tksasha/balance/internal/handlers/utils"
)

type ListHandler struct {
	service cash.Service
}

func NewListHandler(service cash.Service) *ListHandler {
	return &ListHandler{
		service: service,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		utils.E(w, err)

		return
	}

	if err := components.CashList(cashes).Render(w); err != nil {
		utils.E(w, err)
	}
}

func (h *ListHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.service.List(r.Context())
}
