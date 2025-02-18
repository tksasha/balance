package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/index"
	"github.com/tksasha/balance/internal/core/index/components"
)

type Handler struct {
	helpers         *helpers.Helpers
	service         index.Service
	categoryService category.Service
}

func NewHandler(helpers *helpers.Helpers, service index.Service, categoryService category.Service) *Handler {
	return &Handler{
		helpers:         helpers,
		service:         service,
		categoryService: categoryService,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = components.Index(h.helpers, categories, r).Render(w)

	handlers.SetError(w, err)
}

func (h *Handler) handle(r *http.Request) (category.Categories, error) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return nil, err
	}

	return categories, nil
}
