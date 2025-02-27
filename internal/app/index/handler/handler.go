package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/index"
	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type Handler struct {
	*handler.Handler

	indexService    index.Service
	categoryService index.CategoryService
	component       *component.Component
}

func New(
	indexService index.Service,
	categoryService index.CategoryService,
) *Handler {
	return &Handler{
		Handler:         handler.New(),
		indexService:    indexService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.Index(categories, r.URL.Query()).Render(w)

	h.SetError(w, err)
}

func (h *Handler) handle(r *http.Request) (category.Categories, error) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return nil, err
	}

	return categories, nil
}
