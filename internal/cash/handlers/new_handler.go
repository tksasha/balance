package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/cash/components"
	"github.com/tksasha/balance/internal/common/handlers"
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
