package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/indexpage"
	"github.com/tksasha/balance/internal/core/indexpage/components"
)

type Handler struct {
	indexPageService   indexpage.Service
	categoryService    category.Service
	indexPageComponent *components.IndexPageComponent
}

func New(
	indexPageService indexpage.Service,
	categoryService category.Service,
	indexPageComponent *components.IndexPageComponent,
) *Handler {
	return &Handler{
		indexPageService:   indexPageService,
		categoryService:    categoryService,
		indexPageComponent: indexPageComponent,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = h.indexPageComponent.Index(r, categories).Render(w)

	handlers.SetError(w, err)
}

func (h *Handler) handle(r *http.Request) (category.Categories, error) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return nil, err
	}

	return categories, nil
}
