package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/component"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type NewHandler struct {
	cashComponent *component.Component
}

func NewNewHandler(cashComponent *component.Component) *NewHandler {
	return &NewHandler{
		cashComponent: cashComponent,
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	err := h.cashComponent.New(cash).Render(w)

	handlers.SetError(w, err)
}
