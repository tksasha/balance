package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type EditHandler struct {
	service       cash.Service
	cashComponent *components.CashComponent
}

func NewEditHandler(
	service cash.Service,
	cashComponent *components.CashComponent,
) *EditHandler {
	return &EditHandler{
		service:       service,
		cashComponent: cashComponent,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = h.cashComponent.Edit(cash).Render(w)

	handlers.SetError(w, err)
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
