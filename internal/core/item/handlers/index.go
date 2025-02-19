package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type IndexHandler struct {
	service item.Service
	helpers *helpers.Helpers
}

func NewIndexHandler(service item.Service, helpers *helpers.Helpers) *IndexHandler {
	return &IndexHandler{
		service: service,
		helpers: helpers,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = components.Index(h.helpers, items).Render(w)

	handlers.SetError(w, err)
}

func (h *IndexHandler) handle(r *http.Request) (item.Items, error) {
	request := item.IndexRequest{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	return h.service.Index(r.Context(), request)
}
