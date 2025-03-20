package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
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
	params := params.New(r.URL.Query())

	currency := currency.GetByCode(params.Get("currency"))

	cash := &cash.Cash{
		Currency: currency,
	}

	err := h.component.New(params, cash).Render(w)

	h.SetError(w, err)
}
