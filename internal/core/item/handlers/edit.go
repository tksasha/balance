package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type EditHandler struct {
	service         item.Service
	categoryService category.Service
}

func NewEditHandler(service item.Service, categoryService category.Service) *EditHandler {
	return &EditHandler{
		service:         service,
		categoryService: categoryService,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(w, r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = components.Edit(item, categories).Render(w)

	handlers.SetError(w, err)
}

func (h *EditHandler) handle(_ http.ResponseWriter, r *http.Request) (*item.Item, error) {
	return h.service.Edit(r.Context(), r.PathValue("id"))
}
