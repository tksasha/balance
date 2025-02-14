package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type NewHandler struct{}

func NewNewHandler() *NewHandler {
	return &NewHandler{}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	if err := components.CashNew(cash).Render(w); err != nil {
		handlers.E(w, err)
	}
}
