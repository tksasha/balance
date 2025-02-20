package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/index"
	"github.com/tksasha/balance/internal/core/index/components"
)

type IndexHandler struct {
	service            index.Service
	categoryService    category.Service
	indexPageComponent *components.IndexPageComponent
}

func NewIndexHandler(
	service index.Service,
	categoryService category.Service,
	indexPageComponent *components.IndexPageComponent,
) *IndexHandler {
	return &IndexHandler{
		service:            service,
		categoryService:    categoryService,
		indexPageComponent: indexPageComponent,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = h.indexPageComponent.Index(r, categories).Render(w)

	handlers.SetError(w, err)
}

func (h *IndexHandler) handle(r *http.Request) (category.Categories, error) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return nil, err
	}

	return categories, nil
}
