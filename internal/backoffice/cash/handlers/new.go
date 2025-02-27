package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type NewHandler struct {
	*handler.Handler

	component *component.Component
}

func NewNewHandler() *NewHandler {
	return &NewHandler{
		Handler:   handler.New(),
		component: component.New(),
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	err := h.component.New(cash).Render(w)

	h.SetError(w, err)
}
